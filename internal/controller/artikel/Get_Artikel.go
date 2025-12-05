package artikel

import (
	"api_fisioterapi/internal/config"
	artikelModel "api_fisioterapi/internal/models/artikel"

	"github.com/gin-gonic/gin"
)

func GetAllArtikel(c *gin.Context) {
	var list []artikelModel.Artikel
	config.DB.Preload("Galeri").Find(&list)

	c.JSON(200, gin.H{"data": list})
}
func GetArtikelByID(c *gin.Context) {
    id := c.Param("id")
    var artikel artikelModel.Artikel

    if err := config.DB.Preload("Galeri").
        First(&artikel, "id_artikel = ?", id).Error; err != nil {
        c.JSON(404, gin.H{"error": "Artikel tidak ditemukan"})
        return
    }

    c.JSON(200, gin.H{"data": artikel})
}
