package latihanuser

import (
	"api_fisioterapi/internal/config"
	"api_fisioterapi/internal/controller/helpers"
	latihanmodel "api_fisioterapi/internal/models/latihan"
	"strconv"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


func CreateLatihan(c *gin.Context) {
    namaLatihan := c.PostForm("nama_latihan")
    idFase := c.PostForm("id_fase")
    idBagian := c.PostForm("id_bagian")
    levelStr := c.PostForm("level")
    deskripsi := c.PostForm("deskripsi")

    // ===== VALIDASI WAJIB =====
    if namaLatihan == "" || idFase == "" || idBagian == "" || levelStr == "" {
        c.JSON(400, gin.H{"error": "field wajib belum lengkap"})
        return
    }

    level, err := strconv.Atoi(levelStr)
    if err != nil || level < 1 || level > 3 {
        c.JSON(400, gin.H{"error": "level harus angka 1–3"})
        return
    }

    // ===== UPLOAD GAMBAR =====
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

    // ===== SIMPAN DATA =====
    data := latihanmodel.Latihan{
        IDLatihan:   helpers.GenerateRandom4Digit(),
        NamaLatihan: namaLatihan,
        IDFase:      idFase,
        IDBagian:    idBagian,
        Level:       level,
        UrlGambar:   filename,
        Deskripsi:   deskripsi,
		
        CreatedAt:   time.Now(),
    }

    if err := config.DB.Create(&data).Error; err != nil {
        c.JSON(500, gin.H{"error": "gagal menyimpan data"})
        return
    }

    c.JSON(201, gin.H{
        "message": "latihan berhasil dibuat",
        "data":    data,
    })
}
