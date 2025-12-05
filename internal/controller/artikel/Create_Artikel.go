package artikel

import (
	"api_fisioterapi/internal/config"
	artikelModel "api_fisioterapi/internal/models/artikel"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateArtikel(c *gin.Context) {
	var input struct {
		Judul     string `form:"judul"`
		Deskripsi string `form:"deskripsi"`
		IDTags    string `form:"id_tags"`
	}

	if err := c.Bind(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	idArtikel := uuid.New().String()

	artikel := artikelModel.Artikel{
		IDArtikel: idArtikel,
		Judul:     input.Judul,
		Deskripsi: input.Deskripsi,
		IDTags:    input.IDTags,
		CreatedAt: time.Now(),
	}

	// Simpan Artikel
	if err := config.DB.Create(&artikel).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Upload Multiple Files
	form, _ := c.MultipartForm()
	files := form.File["gambar"]

	for _, file := range files {
		fileName := uuid.New().String() + filepath.Ext(file.Filename)
		path := "uploads/artikel/" + fileName

		if err := c.SaveUploadedFile(file, path); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Simpan ke database
		gambar := artikelModel.GaleriGambar{
			IDGambar:  uuid.New().String(),
			UrlFile:   path,
			IDArtikel: idArtikel,
			CreatedAt: time.Now(),
		}

		config.DB.Create(&gambar)
	}

	c.JSON(200, gin.H{
		"message": "Artikel berhasil dibuat",
		"data":    artikel,
	})
}
