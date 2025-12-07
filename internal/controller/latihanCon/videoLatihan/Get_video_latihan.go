package videolatihan

import (
	"api_fisioterapi/internal/config"
	models "api_fisioterapi/internal/models/latihan"

	"github.com/gin-gonic/gin"
)

func GetVideoByLatihan(c *gin.Context) {
	idLatihan := c.Param("id_latihan")

	var videos []models.ListVideoLatihan
	if err := config.DB.Where("id_latihan = ?", idLatihan).Find(&videos).Error; err != nil {
		c.JSON(500, gin.H{"error": "gagal ambil data video"})
		return
	}

	c.JSON(200, gin.H{
		"id_latihan": idLatihan,
		"videos":     videos,
	})
}
