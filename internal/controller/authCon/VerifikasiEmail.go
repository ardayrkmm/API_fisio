package authCon

import (
	ds "api_fisioterapi/internal/config"
	middlewares "api_fisioterapi/internal/middleware"
	userModel "api_fisioterapi/internal/models/users"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateEmailVerificationToken(userID, email string) (string, error) {
	expiration := time.Now().Add(30 * time.Minute)

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"purpose": "email_verification",
		"exp":     expiration.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")

	return token.SignedString([]byte(secret))
}
func SendVerificationToken(c *gin.Context) {
	userID, _ := middlewares.GetUserIDFromContext(c)

	var user userModel.User
	ds.DB.First(&user, "id_user = ?", userID)

	token, err := GenerateEmailVerificationToken(user.IDUser, user.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	// Debug: kirim token (nanti kirim via email)
	c.JSON(200, gin.H{
		"message": "Email verification token generated",
		"token":   token,
	})
}

type VerifyRequest struct {
	Token string `json:"token" binding:"required"`
}

func VerifyEmail(c *gin.Context) {
	var req VerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Token required"})
		return
	}

	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(req.Token, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		c.JSON(400, gin.H{"error": "Invalid token"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	// Cek purpose untuk memastikan token hanya untuk verifikasi email
	if claims["purpose"] != "email_verification" {
		c.JSON(400, gin.H{"error": "Invalid purpose"})
		return
	}

	userID := claims["user_id"].(string)

	// Update status user → inilah satu-satunya operasi DB
	ds.DB.Model(&userModel.User{}).
		Where("id_user = ?", userID).
		Update("verifikasi_status", 1)

	c.JSON(200, gin.H{"message": "Email verified successfully"})
}
