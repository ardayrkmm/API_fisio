package latihanuser

import (
	"api_fisioterapi/internal/config"
	"api_fisioterapi/internal/models/users"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateKondisiUser(c *gin.Context) {
	var input users.KondisiUser

	// 1. Bind JSON dari Body Request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Format data tidak valid: " + err.Error()})
		return
	}

	// 2. Ambil ID User dari Context (Hasil dari Middleware Auth)
	userID, exists := c.Get("id_user")
	if !exists {
		c.JSON(401, gin.H{"message": "Unauthorized: Silahkan login terlebih dahulu"})
		return
	}
	input.IDUser = userID.(string)

	// 3. Generate ID Form & Timestamp (Contoh sederhana)
	// Kamu bisa menggunakan UUID atau format: FRM-001
	input.IDForm = "F" + fmt.Sprintf("%d", time.Now().UnixNano())[:3]
	input.CreatedAt = time.Now()

	// 4. Simpan ke Database
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal menyimpan data kondisi"})
		return
	}

	c.JSON(201, gin.H{
		"message": "Data kondisi berhasil disimpan",
		"data":    input,
	})
}