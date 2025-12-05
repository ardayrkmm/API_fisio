package main

import (
	"log"

	"api_fisioterapi/internal/config"
	"api_fisioterapi/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    config.InitDB()
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)
    log.Println("Server berjalan...")

	r.Run(":8080")
}
