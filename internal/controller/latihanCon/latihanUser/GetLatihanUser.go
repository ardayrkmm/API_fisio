package latihanuser

import (
	"api_fisioterapi/internal/config"
	models "api_fisioterapi/internal/models/latihan"
	"api_fisioterapi/internal/models/users"

	"github.com/gin-gonic/gin"
)

func GetLatihanUser(c *gin.Context) {
	idUser := c.GetString("id_user") // dari JWT

	// Step 1: ambil semua form kondisi user
	var kondisi []users.KondisiUser
	if err := config.DB.Where("id_user = ?", idUser).Find(&kondisi).Error; err != nil {
		c.JSON(500, gin.H{"error": "gagal ambil kondisi user"})
		return
	}

	if len(kondisi) == 0 {
		c.JSON(200, gin.H{"latihan": []string{}})
		return
	}

	// Ambil semua id_form
	var formIDs []string
	for _, k := range kondisi {
		formIDs = append(formIDs, k.IDForm)
	}

	// Step 2: ambil semua latihan berdasarkan banyak IDForm
	var latihan []models.Latihan
	if err := config.DB.
		Preload("ListVideo").
		Where("id_form IN ?", formIDs).
		Find(&latihan).Error; err != nil {
		c.JSON(500, gin.H{"error": "gagal ambil latihan user"})
		return
	}

	c.JSON(200, gin.H{
		"latihan": latihan,
	})
}
