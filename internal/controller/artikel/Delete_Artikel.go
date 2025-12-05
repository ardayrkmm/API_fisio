package artikel

import (
	"api_fisioterapi/internal/config"
	artikelModel "api_fisioterapi/internal/models/artikel"
	"os"

	"github.com/gin-gonic/gin"
)

func DeleteArtikel(c *gin.Context) {
	id := c.Param("id")

	var gambar []artikelModel.GaleriGambar
	config.DB.Where("id_artikel = ?", id).Find(&gambar)

	// Hapus file fisik
	for _, g := range gambar {
		os.Remove(g.UrlFile)
	}

	// Hapus gambar dari DB
	config.DB.Where("id_artikel = ?", id).Delete(&artikelModel.GaleriGambar{})

	// Hapus artikel
	config.DB.Delete(&artikelModel.Artikel{}, "id_artikel = ?", id)

	c.JSON(200, gin.H{"message": "Artikel dan gambar dihapus"})
}
