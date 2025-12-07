package videolatihan

import (
	"api_fisioterapi/internal/config"
	models "api_fisioterapi/internal/models/latihan"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
func toInt(s string) int {
    v, _ := strconv.Atoi(s)
    return v
}

func toFloat(s string) float64 {
    v, _ := strconv.ParseFloat(s, 64)
    return v
}

func CreateListVideo(c *gin.Context) {
	idLatihan := c.PostForm("id_latihan")
	namaGerakan := c.PostForm("nama_gerakan")
	targetSet := c.PostForm("target_set")
	targetRep := c.PostForm("target_repetisi")
	targetWaktu := c.PostForm("target_waktu")

	// upload VIDEO
	video, err := c.FormFile("video")
	if err != nil {
		c.JSON(400, gin.H{"error": "video wajib diupload"})
		return
	}

	filename := "uploads/videos/" + uuid.NewString() + "-" + video.Filename
	c.SaveUploadedFile(video, filename)

	data := models.ListVideoLatihan{
		IDListVideo:    uuid.NewString(),
		IDLatihan:      idLatihan,
		NamaGerakan:    namaGerakan,
		TargetSet:      toInt(targetSet),
		TargetRepetisi: toInt(targetRep),
		TargetWaktu:    toFloat(targetWaktu),
		
		VideoURL:       filename,
		CreatedAt:      time.Now(),
	}

	config.DB.Create(&data)

	c.JSON(201, gin.H{"msg": "list video dibuat", "data": data})
}