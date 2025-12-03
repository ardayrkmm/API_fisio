package authCon

import (
	"net/http"
	"strings"
	"time"

	userModel "api_fisioterapi/internal/models/users"

	ds "api_fisioterapi/internal/config"
	middlewares "api_fisioterapi/internal/middleware"
	rs "api_fisioterapi/internal/models/users/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)



func Register(c *gin.Context) {
	var req rs.RegisterRequest

	// Validasi dan bind request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"message": err.Error(),
		})
		return
	}

	// Normalize email ke lowercase
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))
	req.Nama = strings.TrimSpace(req.Nama)
	

	// Cek apakah email sudah terdaftar
	var existingUser userModel.User
	result := ds.DB.Where("email = ?", req.Email).First(&existingUser)

	if result.Error == nil {
		// User dengan email ini sudah ada
		c.JSON(http.StatusConflict, gin.H{
			"error":   "Email already registered",
			"message": "An account with this email already exists",
		})
		return
	} else if result.Error != gorm.ErrRecordNotFound {
		// Error database lainnya
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Database error",
			"message": "Failed to check existing user",
		})
		return
	}

	// Buat user baru
	user := userModel.User{
		Nama:     req.Nama,
		Email:    req.Email,
		Password: req.Password, // Akan di-hash oleh BeforeCreate hook
		NoTelepon: req.NomerTelepon, // Akan di-hash oleh BeforeCreate hook

		    Role:             "user",                // default role agar tidak kosong
    VerifikasiStatus: 0,                     // default belum verifikasi
    CreatedAt:        time.Now(), 
	}

	// Simpan ke database
	if err := ds.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Registration failed",
			"message": "Failed to create user account",
		})
		return
	}

	// Generate JWT token untuk auto-login setelah register
	token, err := middlewares.GenerateToken(user.IDUser, user.Email, user.Nama)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Token generation failed",
			"message": "Account created but failed to generate authentication token",
		})
		return
	}

	// Success response
	c.JSON(http.StatusCreated, rs.AuthResponse{
		Message: "Registration successful",
		User:    user.ToPublicUser(),
		Token:   token,
	})
}