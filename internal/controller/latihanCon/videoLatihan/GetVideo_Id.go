package videolatihan

import (
	"api_fisioterapi/internal/config"
	models "api_fisioterapi/internal/models/latihan"

	"github.com/gin-gonic/gin"
)

func GetVideoByID(c *gin.Context) {
	id := c.Param("id_list_video")

	var video models.ListVideoLatihan
	if err := config.DB.First(&video, "id_list_video = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"error": "video tidak ditemukan"})
		return
	}

	c.JSON(200, video)
}
