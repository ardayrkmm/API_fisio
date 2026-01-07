package latihanuser

import (
	"api_fisioterapi/internal/config"
	models "api_fisioterapi/internal/models/latihan"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category models.FaseRehabilitasi

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := config.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal menyimpan data",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Kategori berhasil ditambahkan",
		"data":    category,
	})
}

func GetCategories(c *gin.Context) {
	var categories []models.FaseRehabilitasi

	if err := config.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengambil data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}
func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	var category models.FaseRehabilitasi

	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Kategori tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.FaseRehabilitasi

	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Kategori tidak ditemukan",
		})
		return
	}

	var input models.FaseRehabilitasi
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	category.NamaFase = input.NamaFase

	config.DB.Save(&category)

	c.JSON(http.StatusOK, gin.H{
		"message": "Kategori berhasil diupdate",
		"data":    category,
	})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.FaseRehabilitasi

	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Kategori tidak ditemukan",
		})
		return
	}

	config.DB.Delete(&category)

	c.JSON(http.StatusOK, gin.H{
		"message": "Kategori berhasil dihapus",
	})
}
