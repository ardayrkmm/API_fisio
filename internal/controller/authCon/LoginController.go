package authCon

import (
	"net/http"
	"strings"

	userModel "api_fisioterapi/internal/models/users"

	ds "api_fisioterapi/internal/config"
	middlewares "api_fisioterapi/internal/middleware"
	rs "api_fisioterapi/internal/models/users/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
    var req rs.LoginRequest

    // Validasi dan bind request body
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Validation failed",
            "message": err.Error(),
        })
        return
    }

    // Normalize email
    req.Email = strings.ToLower(strings.TrimSpace(req.Email))

    // Cari user berdasarkan email
    var user userModel.User
    if err := ds.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            // User tidak ditemukan
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Authentication failed",
                "message": "Invalid email or password",
            })
        } else {
            // Database error
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Database error",
                "message": "Failed to fetch user data",
            })
        }
        return
    }

    // Verifikasi password
    if err := user.CheckPassword(req.Password); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "Authentication failed",
            "message": "Invalid email or password",
        })
        return
    }

    // Generate JWT token
    token, err := middlewares.GenerateToken(user.IDUser, user.Email, user.Nama)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Token generation failed",
            "message": "Login successful but failed to generate authentication token",
        })
        return
    }

    // Success response
    c.JSON(http.StatusOK, rs.AuthResponse{
        Message: "Login successful",
        User:    user.ToPublicUser(),
        Token:   token,
    })
}