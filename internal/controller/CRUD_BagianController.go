package controller

import (
	"api_fisioterapi/internal/config"
	"api_fisioterapi/internal/models/users"
	"api_fisioterapi/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BagianTubuhRequest struct {
	NamaBagian string `json:"nama_bagian"`
}

func CreateBagianTubuh(c *gin.Context) {
	var req BagianTubuhRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "format request tidak valid"})
		return
	}

	if req.NamaBagian == "" {
		c.JSON(400, gin.H{"error": "nama_bagian wajib diisi"})
		return
	}

	data := users.BagianTubuh{
		IDBagian:   services.GenerateRandom4Digit(),
		NamaBagian: req.NamaBagian,
	}

	if err := config.DB.Create(&data).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"message": "bagian tubuh berhasil dibuat",
		"data":    data,
	})
}

func GetAllBagianTubuh(c *gin.Context) {
	var data []users.BagianTubuh

	if err := config.DB.Find(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
func GetBagianTubuhByID(c *gin.Context) {
	id := c.Param("id")

	var data users.BagianTubuh
	if err := config.DB.First(&data, "id_bagian = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
func UpdateBagianTubuh(c *gin.Context) {
	id := c.Param("id")
	nama := c.PostForm("nama_bagian")

	if nama == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nama_bagian wajib diisi"})
		return
	}

	var data users.BagianTubuh
	if err := config.DB.First(&data, "id_bagian = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data tidak ditemukan"})
		return
	}

	data.NamaBagian = nama

	if err := config.DB.Save(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "bagian tubuh berhasil diperbarui",
		"data":    data,
	})
}
