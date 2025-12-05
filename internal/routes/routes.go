package routes

import (
	authCon "api_fisioterapi/internal/controller/authCon"
	"api_fisioterapi/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// Public routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", authCon.Register)
		auth.POST("/login", authCon.Login)
	}

	// Protected routes (require JWT)
	protected := r.Group("/auth")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", authCon.GetProfile)
		protected.POST("/refresh-token", authCon.RefreshToken)

		// Email verification (NON-DB, JWT)
		protected.POST("/send-verification", authCon.SendVerificationToken)
		protected.POST("/verify-email", authCon.VerifyEmail)
	}

	// Health Check
	r.GET("/health", authCon.HealthCheck)
}