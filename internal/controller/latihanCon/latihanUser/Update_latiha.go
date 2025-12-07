package latihanuser

import (
	"api_fisioterapi/internal/config"
	latihanmodel "api_fisioterapi/internal/models/latihan"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UpdateLatihan(c *gin.Context) {
	id := c.Param("id")

	var latihan latihanmodel.Latihan
	if err := config.DB.First(&latihan, "id_latihan = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"error": "latihan tidak ditemukan"})
		return
	}

	nama := c.PostForm("nama_latihan")
	idKategori := c.PostForm("id_kategori")
	deskripsi := c.PostForm("deskripsi")

	if nama != "" {
		latihan.NamaLatihan = nama
	}
	if idKategori != "" {
		latihan.IDKategori = idKategori
	}
	if deskripsi != "" {
		latihan.Deskripsi = deskripsi
	}

	file, err := c.FormFile("gambar")
	if err == nil {
		filename := "uploads/images/" + uuid.NewString() + "-" + file.Filename
		c.SaveUploadedFile(file, filename)
		latihan.UrlGambar = filename
	}

	config.DB.Save(&latihan)
	c.JSON(200, gin.H{"msg": "updated", "data": latihan})
}
