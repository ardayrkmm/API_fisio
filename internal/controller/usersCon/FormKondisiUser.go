package usersCon

import (
	ds "api_fisioterapi/internal/config"
	userModel "api_fisioterapi/internal/models/users"
	"api_fisioterapi/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)
type KondisiRequest struct {
	IDUser        string `json:"id_user" binding:"required"`
	IDBagian      string `json:"id_bagian" binding:"required"`
	LamaNyeriHari int    `json:"lama_nyeri_hari" binding:"required"`
	TingkatNyeri  int    `json:"tingkat_nyeri" binding:"required"`
	Catatan       string `json:"catatan"`
}
func BuatKondisiUser(c *gin.Context) {
idUser := c.GetString("id_user")
var req KondisiRequest
if err := c.ShouldBindJSON(&req); err != nil {
	c.JSON(400, gin.H{"error": "Invalid request", "message": err.Error()})
	return
}

kondisi :=userModel.KondisiUser{
	IDUser:        idUser,
	IDBagian:      req.IDBagian,
	IDForm: utils.RandomString(3) ,
	LamaNyeriHari: req.LamaNyeriHari,
	TingkatNyeri:  req.TingkatNyeri,
	Catatan:       req.Catatan,
	CreatedAt:     time.Now(),
}

if err := ds.DB.Create(&kondisi).Error; err != nil {
	c.JSON(500, gin.H{"error": "Failed to create kondisi user", "message": err.Error()})
	return	}
c.JSON(201, gin.H{"message": "Kondisi user created successfully", "data": kondisi})

}