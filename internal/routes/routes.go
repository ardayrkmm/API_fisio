package routes

import (
	artikelCon "api_fisioterapi/internal/controller/artikel"
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

	artikel := r.Group("/api/artikel")
	artikel.Use(middleware.AuthMiddleware())
	{
    artikel.POST("/", artikelCon.CreateArtikel)
    artikel.GET("/", artikelCon.GetAllArtikel)
    artikel.GET("/:id", artikelCon.GetArtikelByID)
    artikel.PUT("/:id", artikelCon.UpdateArtikel)
    artikel.DELETE("/:id", artikelCon.DeleteArtikel)
	}
	 r.NoRoute(func(c *gin.Context) {
        c.JSON(404, gin.H{
            "error": "Not Found",
            "message": "The requested endpoint does not exist",
        })
    })
	// Health Check
	r.GET("/health", authCon.HealthCheck)
}


func SetupMiddlewares(router *gin.Engine) {

    router.Use(gin.Recovery())

    router.Use(gin.Logger())

    
    router.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    })
}