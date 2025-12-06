package usersCon

import (
	ds "api_fisioterapi/internal/config"
	userModel "api_fisioterapi/internal/models/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetKondisiByUser(c *gin.Context) {
	// Ambil ID user dari JWT
	idUser := c.GetString("id_user")

	var kondisi []userModel.KondisiUser

	// Ambil data berdasarkan id_user
	if err := ds.DB.Where("id_user = ?", idUser).Find(&kondisi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil data kondisi",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Berhasil mengambil data",
		"data":    kondisi,
	})
}