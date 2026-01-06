package latihanuser

import (
	"api_fisioterapi/internal/config"
	"api_fisioterapi/internal/models/latihan"
	"time"

	"github.com/gin-gonic/gin"
)

func GetJadwalHariIni(c *gin.Context) {
	userID, _ := c.Get("id_user")
	today := time.Now().Format("2006-01-02")

	// 1. Cari Jadwal User hari ini
	var jadwal latihan.JadwalLatihanUser
	if err := config.DB.Where("id_user = ? AND DATE(tanggal) = ?", userID, today).First(&jadwal).Error; err != nil {
		c.JSON(200, gin.H{"message": "Hari ini jadwal istirahat", "latihan": []interface{}{}})
		return
	}

	// 2. Ambil Latihan yang terkait dengan jadwal ini (lewat JadwalLatihanDetail)
	// Kita ambil ID Latihannya dulu
	var detailJadwal []latihan.JadwalLatihanDetail
	config.DB.Where("id_jadwal = ?", jadwal.IDJadwal).Order("urutan ASC").Find(&detailJadwal)

	// 3. Buat Response Terstruktur
	var result []gin.H

	// Ambil semua Video unik dari detail jadwal, lalu kelompokkan berdasarkan Latihan
	for _, d := range detailJadwal {
		var video latihan.ListVideoLatihan
		config.DB.Where("id_list_video = ?", d.IDListVideo).First(&video)

		var lat latihan.Latihan
		config.DB.Where("id_latihan = ?", video.IDLatihan).First(&lat)

		// Cek apakah Latihan ini sudah ada di dalam result?
		found := false
		for i, item := range result {
			if item["id_latihan"] == lat.IDLatihan {
				// Jika sudah ada, tambahkan videonya ke array video yang sudah ada
				result[i]["list_video"] = append(result[i]["list_video"].([]latihan.ListVideoLatihan), video)
				found = true
				break
			}
		}

		// Jika belum ada, buat entri latihan baru
		if !found {
			result = append(result, gin.H{
				"id_latihan":   lat.IDLatihan,
				"nama_latihan": lat.NamaLatihan,
				"deskripsi":    lat.Deskripsi,
				"gambar":       lat.UrlGambar,
				"list_video":   []latihan.ListVideoLatihan{video},
			})
		}
	}

	c.JSON(200, gin.H{
		"id_jadwal": jadwal.IDJadwal,
		"tanggal":   jadwal.Tanggal.Format("2006-01-02"),
		"program":   result,
	})
}