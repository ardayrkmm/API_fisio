package authCon

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	userModel "api_fisioterapi/internal/models/users"

	ds "api_fisioterapi/internal/config"
	middlewares "api_fisioterapi/internal/middleware"
	rs "api_fisioterapi/internal/models/users/auth"
	dss "api_fisioterapi/internal/services"

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

    // --- PERBAIKAN DI SINI ---
    // 1. Generate 4 angka acak (OTP)
    rand.Seed(time.Now().UnixNano())
    otpCode := fmt.Sprintf("%04d", rand.Intn(10000))

    vToken, otpCode, err := GenerateEmailVerificationToken(user.IDUser, user.Email)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to generate verification token"})
        return
    }

    // 2. Kirim ANGKA-nya ke email
    err = dss.SendVerificationEmail(user.Email, otpCode)
    if err != nil {
        fmt.Println("Error sending email:", err)
    }

    // 3. Generate token login (seperti biasa)
    loginToken, _ := middlewares.GenerateToken(user.IDUser, user.Email, user.Nama)

    // 4. Kirim Response
    c.JSON(http.StatusCreated, gin.H{
        "message": "Registrasi berhasil, cek email untuk kode OTP",
        "user":    user.ToPublicUser(),
        "token":   loginToken,      // Ini untuk login
        "verification_token": vToken, // INI YANG HARUS DIPAKAI DI ENDPOINT VERIFY
    })
}