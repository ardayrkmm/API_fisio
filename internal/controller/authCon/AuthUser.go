package authCon

import (
	"net/http"

	userModel "api_fisioterapi/internal/models/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	ds "api_fisioterapi/internal/config"
	middlewares "api_fisioterapi/internal/middleware"
)

// GetProfile handler untuk mendapatkan data user yang sedang login
func GetProfile(c *gin.Context) {
    // Ambil user ID dari context (di-set oleh auth middleware)
    userID, exists := middlewares.GetUserIDFromContext(c)
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "Unauthorized",
            "message": "User ID not found in context",
        })
        return
    }

    // Fetch user data dari database
    var user userModel.User
    if err := ds.DB.First(&user, userID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{
                "error": "User not found",
                "message": "User account no longer exists",
            })
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Database error",
                "message": "Failed to fetch user profile",
            })
        }
        return
    }

    // Success response
    c.JSON(http.StatusOK, gin.H{
        "message": "Profile fetched successfully",
        "user": user.ToPublicUser(),
    })
}

// RefreshToken handler untuk refresh JWT token
func RefreshToken(c *gin.Context) {
    // Ambil user info dari context
    userID, _ := middlewares.GetUserIDFromContext(c)
    userEmail, _ := c.Get("userEmail")
    userName, _ := c.Get("userName")

    // Generate token baru
    token, err := middlewares.GenerateToken(
        userID,
        userEmail.(string),
        userName.(string),
    )

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Token generation failed",
            "message": "Failed to refresh authentication token",
        })
        return
    }

    // Success response
    c.JSON(http.StatusOK, gin.H{
        "message": "Token refreshed successfully",
        "token": token,
    })
}

// HealthCheck handler untuk cek status API
func HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status": "healthy",
        "message": "API is running",
    })
}