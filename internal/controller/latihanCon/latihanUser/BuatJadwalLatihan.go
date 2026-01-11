package latihanuser

import (
	"api_fisioterapi/internal/config"
	"api_fisioterapi/internal/models/latihan"
	models "api_fisioterapi/internal/models/latihan"
	ss "api_fisioterapi/internal/models/users"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GenerateJadwalOtomatis(c *gin.Context) {

	// ✅ Ambil userID dari middleware
	val, exists := c.Get("userID")
	if !exists || val == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := val.(string)
	if !ok || userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	// 2. Ambil kondisi terakhir user
	var kondisi ss.KondisiUser
	if err := config.DB.
		Where("id_user = ?", userID).
		Order("created_at DESC").
		First(&kondisi).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Silahkan isi form kondisi terlebih dahulu",
		})
		return
	}

	// 3. Validasi tingkat nyeri
	if kondisi.TingkatNyeri > 7 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ke_dokter",
			"message": "Tingkat nyeri berat, silahkan konsultasi ke dokter",
		})
		return
	}

	// 4. Ambil latihan sesuai bagian
	var latihanTersedia []models.Latihan
	if err := config.DB.
		Preload("ListVideos").
		Where("id_bagian = ?", kondisi.IDBagian).
		Find(&latihanTersedia).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mencari latihan",
		})
		return
	}

	// 5. Simpan jadwal
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		for _, lat := range latihanTersedia {
			jadwal := models.JadwalLatihanUser{
				IDJadwal:  uuid.New().String()[:4],
				IDUser:    userID,
				IDLatihan: lat.IDLatihan,
				Tanggal:   time.Now().AddDate(0, 0, 1),
				Status:    "locked",
				CreatedAt: time.Now(),
			}
			if err := tx.Create(&jadwal).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal membuat jadwal otomatis",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Jadwal rehabilitasi berhasil dibuat",
	})
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

func GetJadwalUser(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	week := c.DefaultQuery("week", "1")

	var jadwal []models.JadwalLatihanUser
	if err := config.DB.
		Preload("Latihan").
		Where("id_user = ?", userID).
		Order("tanggal ASC").
		Find(&jadwal).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil jadwal"})
		return
	}

	if len(jadwal) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Jadwal belum tersedia"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"week":   week,
		"jadwal": jadwal,
	})
}
