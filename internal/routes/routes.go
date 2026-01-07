package routes

import (
	ques "api_fisioterapi/internal/controller"
	artikelCon "api_fisioterapi/internal/controller/artikel"
	authCon "api_fisioterapi/internal/controller/authCon"
	latihanUserCon "api_fisioterapi/internal/controller/latihanCon/latihanUser"

	videoLatihanCon "api_fisioterapi/internal/controller/latihanCon/videoLatihan"
	"api_fisioterapi/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	//
	api := r.Group("/api")
	{
q := api.Group("/questions")
	{
	q.POST("", ques.CreateQuestion)
	q.GET("", ques.GetQuestions)
	q.PUT("/:id", ques.UpdateQuestion)
	q.DELETE("/:id", ques.DeleteQuestion)
}

//auth

auth := api.Group("/auth")
	{
		auth.POST("/register", authCon.Register)
		auth.POST("/login", authCon.Login)
		auth.POST("/send-verification", authCon.SendVerificationToken)
		auth.POST("/verify-email", authCon.VerifyEmail)
	}
//api
	category := api.Group("/categories")
	{
		category.POST("/", latihanUserCon.CreateCategory)
		category.GET("/", latihanUserCon.GetCategories)
		category.GET("/:id", latihanUserCon.GetCategoryByID)
		category.PUT("/:id", latihanUserCon.UpdateCategory)
		category.DELETE("/:id", latihanUserCon.DeleteCategory)
	}

		protected := api.Group("/auth")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", authCon.GetProfile)
		protected.POST("/refresh-token", authCon.RefreshToken)

		// Email verification (NON-DB, JWT)
		
	}
		latihanAdmin := api.Group("/latihanAdmin")
		{
				latihanAdmin.POST("/", latihanUserCon.CreateLatihan)
	latihanAdmin.GET("/usr", latihanUserCon.GetLatihan)

	latihanAdmin.PUT("/:id", latihanUserCon.UpdateLatihan)
	latihanAdmin.DELETE("/:id", latihanUserCon.DeleteLatihan)
		}
	latihanUser := api.Group("/latihanuser")
	latihanUser.Use(middleware.AuthMiddleware())
	{
		latihanUser.POST("/jadwal/:id_jadwal/selesai", latihanUserCon.SelesaikanLatihan)
	
	latihanUser.POST("/kondisi", latihanUserCon.CreateKondisiUser)


	// LIST VIDEO DALAM LATIHAN USER
	video := latihanUser.Group("/video")
	{
		video.POST("/", videoLatihanCon.CreateListVideo)
		video.GET("/", videoLatihanCon.GetVideoByLatihan)
		video.GET("/:id", videoLatihanCon.GetVideoByLatihan)
		video.PUT("/:id", videoLatihanCon.UpdateVideo)
		
	}
	}
		tubuh := api.Group("/bagian-tubuh")
	{
		tubuh.POST("/", ques.CreateBagianTubuh)
		tubuh.GET("/", ques.GetAllBagianTubuh)
		tubuh.GET("/:id", ques.GetBagianTubuhByID)
		tubuh.PUT("/:id", ques.UpdateBagianTubuh)
	
	}

	artikel := api.Group("/artikel")
	artikel.Use(middleware.AuthMiddleware())
	{
    artikel.POST("/", artikelCon.CreateArtikel)
    artikel.GET("/", artikelCon.GetAllArtikel)
    artikel.GET("/:id", artikelCon.GetArtikelByID)
    artikel.PUT("/:id", artikelCon.UpdateArtikel)
    artikel.DELETE("/:id", artikelCon.DeleteArtikel)
	}
	}
	
	
	

	// Protected routes (require JWT)


	
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