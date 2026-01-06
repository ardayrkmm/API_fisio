package latihanuser

import (
	"api_fisioterapi/internal/config"
	"api_fisioterapi/internal/models/latihan"

	"github.com/gin-gonic/gin"
)

func GetJadwalPerMinggu(c *gin.Context) {
	userID, _ := c.Get("id_user")

	// 1. Ambil semua jadwal milik user, urutkan berdasarkan tanggal tertua
	var semuaJadwal []latihan.JadwalLatihanUser
	config.DB.Where("id_user = ?", userID).Order("tanggal ASC").Find(&semuaJadwal)

	if len(semuaJadwal) == 0 {
		c.JSON(404, gin.H{"message": "Jadwal belum dibuat"})
		return
	}

	// 2. Gunakan Map untuk mengelompokkan (Minggu 1, 2, 3)
	// Kita hitung selisih hari dari tanggal latihan pertama
	tanggalMulai := semuaJadwal[0].Tanggal

	// Response terstruktur
	type Mingguan struct {
		Minggu   int           `json:"minggu"`
		Fase     string        `json:"fase"`
		Latihans []interface{} `json:"daftar_hari"`
	}

	var responMingguan []Mingguan
	// Inisialisasi 3 minggu (sesuai permintaan kamu)
	for i := 1; i <= 3; i++ {
		fase := "Pemulihan"
		if i == 1 {
			fase = "Fase Akut (Nyeri Berat)"
		}
		if i == 3 {
			fase = "Fase Penguatan"
		}

		responMingguan = append(responMingguan, Mingguan{
			Minggu:   i,
			Fase:     fase,
			Latihans: []interface{}{},
		})
	}

	// 3. Masukkan jadwal ke grup minggunya masing-masing
	for _, j := range semuaJadwal {
		diff := int(j.Tanggal.Sub(tanggalMulai).Hours() / 24)
		indexMinggu := diff / 7 // 0-6 hari = minggu 1, 7-13 = minggu 2, dst

		if indexMinggu < 3 { // Batasi hanya 3 minggu
			// Ambil detail singkat (nama latihan saja untuk ringkasan)
			var detail []latihan.JadwalLatihanDetail
			config.DB.Where("id_jadwal = ?", j.IDJadwal).Find(&detail)

			responMingguan[indexMinggu].Latihans = append(responMingguan[indexMinggu].Latihans, gin.H{
				"id_jadwal":      j.IDJadwal,
				"tanggal":        j.Tanggal.Format("2006-01-02"),
				"jumlah_gerakan": len(detail),
			})
		}
	}

	c.JSON(200, gin.H{
		"user_id": userID,
		"program": responMingguan,
	})
}