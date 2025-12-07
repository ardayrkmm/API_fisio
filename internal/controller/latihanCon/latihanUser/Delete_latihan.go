package latihanuser

import (
	"api_fisioterapi/internal/config"
	models "api_fisioterapi/internal/models/latihan"

	"github.com/gin-gonic/gin"
)

func DeleteLatihan(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Latihan{}, "id_latihan = ?", id).Error; err != nil {
		c.JSON(400, gin.H{"error": "gagal hapus"})
		return
	}

	c.JSON(200, gin.H{"msg": "deleted"})
}
