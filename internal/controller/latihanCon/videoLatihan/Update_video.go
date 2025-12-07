package videolatihan

import (
	"api_fisioterapi/internal/config"
	models "api_fisioterapi/internal/models/latihan"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UpdateVideo(c *gin.Context) {
	id := c.Param("id_list_video")

	var video models.ListVideoLatihan
	if err := config.DB.First(&video, "id_list_video = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"error": "video tidak ditemukan"})
		return
	}

	// form text
	namaGerakan := c.PostForm("nama_gerakan")
	targetSet := c.PostForm("target_set")
	targetRep := c.PostForm("target_repetisi")
	targetWaktu := c.PostForm("target_waktu")

	// UPDATE nilai jika dikirim
	if namaGerakan != "" {
		video.NamaGerakan = namaGerakan
	}
	if targetSet != "" {
		video.TargetSet = toInt(targetSet)
	}
	if targetRep != "" {
		video.TargetRepetisi = toInt(targetRep)
	}
	if targetWaktu != "" {
		video.TargetWaktu = toFloat(targetWaktu)
	}

	// cek apakah ada upload video baru
	file, err := c.FormFile("video")
	if err == nil {
		filename := "uploads/videos/" + uuid.NewString() + "-" + file.Filename
		c.SaveUploadedFile(file, filename)
		video.VideoURL = filename
	}

	config.DB.Save(&video)

	c.JSON(200, gin.H{
		"msg":   "video berhasil diupdate",
		"video": video,
	})
}
