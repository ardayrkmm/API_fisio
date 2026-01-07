package latihanuser

import (
	"api_fisioterapi/internal/config"
	"api_fisioterapi/internal/models/latihan"
	"api_fisioterapi/internal/models/users"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func GenerateJadwalOtomatis(kondisi users.KondisiUser) {
	// 1. Tentukan intensitas berdasarkan tingkat nyeri
	// Nyeri 7-10: 3x seminggu (Fase Akut)
	// Nyeri 1-6: 4x seminggu (Fase Pemulihan)
	frekuensi := 4
	if kondisi.TingkatNyeri >= 7 {
		frekuensi = 3
	}

	// 2. Ambil Video Latihan yang sesuai dengan Bagian Tubuh
	var videos []latihan.ListVideoLatihan
	config.DB.Table("list_video_latihans").
		Joins("JOIN latihans ON latihans.id_latihan = list_video_latihans.id_latihan").
		Where("latihans.id_bagian = ?", kondisi.IDBagian).
		Find(&videos)

	if len(videos) == 0 {
		return
	}

	// 3. Loop untuk 3 Minggu (21 Hari)
	for minggu := 0; minggu < 3; minggu++ {
		for i := 0; i < frekuensi; i++ {
			// Jeda antar hari (misal latihan tiap hari ke-1, 3, 5, dst)
			hariTambah := (minggu * 7) + (i * 2)
			tanggalLatihan := kondisi.CreatedAt.AddDate(0, 0, hariTambah)

			// A. Simpan ke JadwalLatihanUser
			jadwalID := "J" + fmt.Sprintf("%d%d", kondisi.IDUser, tanggalLatihan.Unix())[:3]
			jadwalBaru := latihan.JadwalLatihanUser{
				IDJadwal:  jadwalID,
				IDUser:    kondisi.IDUser,
				Tanggal:   tanggalLatihan,
					Status:    "locked",
				CreatedAt: time.Now(),
			}
			if minggu == 0 && i == 0 {
				jadwalBaru.Status = "unlocked"
			}

			config.DB.Create(&jadwalBaru)

			// B. Masukkan gerakan video ke JadwalLatihanDetail
			for urutan, v := range videos {
				detail := latihan.JadwalLatihanDetail{
					IDDetail:    fmt.Sprintf("D%d%d", urutan, time.Now().UnixNano())[:4],
					IDJadwal:    jadwalID,
					IDListVideo: v.IDListVideo,
					Urutan:      urutan + 1,
				}
				config.DB.Create(&detail)
			}
		}
	}
}


func SelesaikanLatihan(c *gin.Context) {
	idJadwal := c.Param("id_jadwal")

	tx := config.DB.Begin()

	var jadwal latihan.JadwalLatihanUser
	if err := tx.Where("id_jadwal = ?", idJadwal).First(&jadwal).Error; err != nil {
		tx.Rollback()
		c.JSON(404, gin.H{"error": "jadwal tidak ditemukan"})
		return
	}

	// ❗ VALIDASI STATUS
	if jadwal.Status != "unlocked" {
		tx.Rollback()
		c.JSON(400, gin.H{"error": "latihan belum dapat dikerjakan"})
		return
	}

	// 1️⃣ set done
	if err := tx.Model(&jadwal).Update("status", "done").Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "gagal update status"})
		return
	}

	// 2️⃣ unlock next
	var next latihan.JadwalLatihanUser
	err := tx.
		Where("id_user = ? AND tanggal > ? AND status = ?", jadwal.IDUser, jadwal.Tanggal, "locked").
		Order("tanggal asc").
		First(&next).Error

	if err == nil {
		tx.Model(&next).Update("status", "unlocked")
	}

	tx.Commit()

	c.JSON(200, gin.H{
		"message": "latihan selesai",
	})
}