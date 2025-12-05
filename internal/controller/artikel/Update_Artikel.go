package artikel

import (
	"api_fisioterapi/internal/config"
	artikelModel "api_fisioterapi/internal/models/artikel"

	"github.com/gin-gonic/gin"
)

func UpdateArtikel(c *gin.Context) {
	id := c.Param("id")
	var artikel artikelModel.Artikel

	if err := config.DB.First(&artikel, "id_artikel = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Artikel tidak ditemukan"})
		return
	}

	var input struct {
		Judul     string `json:"judul"`
		Deskripsi string `json:"deskripsi"`
		IDTags    string `json:"id_tags"`
	}

	c.BindJSON(&input)

	config.DB.Model(&artikel).Updates(input)

	c.JSON(200, gin.H{"message": "Artikel diperbarui"})
}
