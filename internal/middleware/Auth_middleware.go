package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID string, email, name string) (string, error) {
	// Token expiration time (24 jam)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Membuat claims (payload) untuk token
	claims := &Claims{
		UserID: userID,
		Email:  email,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "go-gin-auth-api",
			Subject:   userID,
		},
	}

	// Membuat token dengan claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token dengan secret key
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", fmt.Errorf("JWT_SECRET is not set")
	}

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// AuthMiddleware memverifikasi JWT token dari request header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			c.Abort()
			return
		}

		// Format: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Authorization format",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]
		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "JWT secret not configured",
			})
			c.Abort()
			return
		}

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return []byte(jwtSecret), nil
			},
		)

		// ✅ FIX JWT v5 ERROR HANDLING
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Token expired",
				})
			} else if errors.Is(err, jwt.ErrSignatureInvalid) {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Invalid token signature",
				})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
			}
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		// ✅ SET CONTEXT (WAJIB)
		c.Set("userID", claims.UserID)
		c.Set("userEmail", claims.Email)
		c.Set("userName", claims.Name)
		c.Set("claims", claims)

		c.Next()
	}
}

// GetUserIDFromContext mengambil user ID dari context
func GetUserIDFromContext(c *gin.Context) (string, bool) {
	userID, exists := c.Get("userID")
	if !exists {
		return "", false
	}

	// Type assertion untuk convert ke uint
	if id, ok := userID.(string); ok {
		return id, true
	}

	return "", false
}