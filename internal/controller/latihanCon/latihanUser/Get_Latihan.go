package latihanuser

import (
	"api_fisioterapi/internal/config"
	models "api_fisioterapi/internal/models/latihan"

	"github.com/gin-gonic/gin"
)

func GetLatihan(c *gin.Context) {
	var data []models.Latihan
	config.DB.Find(&data)
	c.JSON(200, data)
}
