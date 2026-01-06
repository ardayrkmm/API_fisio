package latihanuser

import (
	"api_fisioterapi/internal/config"
	"api_fisioterapi/internal/controller/helpers"
	latihanmodel "api_fisioterapi/internal/models/latihan"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


func CreateLatihan(c *gin.Context) {
	nama := c.PostForm("nama_latihan")
	idKategori := c.PostForm("id_kategori")
	deskripsi := c.PostForm("deskripsi")

	// upload image
	file, err := c.FormFile("gambar")
	if err != nil {
		c.JSON(400, gin.H{"error": "gambar wajib diupload"})
		return
	}

	filename := "uploads/images/" + uuid.NewString() + "-" + file.Filename
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(500, gin.H{"error": "gagal upload gambar"})
		return
	}

	idLatihan := helpers.GenerateRandom4Digit()
	data := latihanmodel.Latihan{
		IDLatihan: idLatihan,
		NamaLatihan: nama,
		IDKategori:  idKategori,
		UrlGambar:   filename,
		Deskripsi:   deskripsi,
		CreatedAt:   time.Now(),
	}

	config.DB.Create(&data)

	c.JSON(201, gin.H{"msg": "latihan dibuat", "data": data})
}