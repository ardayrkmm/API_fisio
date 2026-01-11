package latihanuser

import (
	"api_fisioterapi/internal/config"
	"api_fisioterapi/internal/models/latihan"
	"api_fisioterapi/internal/models/users"
	"api_fisioterapi/internal/services"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


func CreateKondisiUser(c *gin.Context) {
    type JawabanInput struct {
    QuestionID string `json:"question_id"`
    OptionID   string `json:"option_id"`
}

type KondisiRequest struct {
    IDBagian string          `json:"id_bagian"`
    Answers  []JawabanInput `json:"answers"`
}

    var req KondisiRequest
    var kondisi users.KondisiUser

    // 1. Bind JSON
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid"})
        return
    }

val, exists := c.Get("userID")

    if !exists || val == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Sesi tidak valid"})
        return
    }

    userID, ok := val.(string)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Tipe data User ID tidak sesuai"})
        return
    }

    kondisi.IDForm = services.GenerateRandom4Digit()
    kondisi.IDUser = userID
    kondisi.IDBagian = req.IDBagian
    kondisi.CreatedAt = time.Now()

    err := config.DB.Transaction(func(tx *gorm.DB) error {

        // === MAPPING ANSWER KE KONDISI USER ===
        for _, ans := range req.Answers {
            var question users.Question
            var option users.QuestionOption

            if err := tx.First(&question, "id = ?", ans.QuestionID).Error; err != nil {
                return err
            }

            if err := tx.First(&option, "id = ?", ans.OptionID).Error; err != nil {
                return err
            }

            switch question.TargetField {
            case "tingkat_nyeri":
                kondisi.TingkatNyeri = option.Nilai
            case "lama_nyeri_hari":
                kondisi.LamaNyeriHari = option.Nilai
            case "jenis_keluhan":
                kondisi.JenisKeluhan = option.Label
            case "id_bagian":
                var bagian users.BagianTubuh
                err := tx.
                          Where("LOWER(nama_bagian) = LOWER(?)", option.Label).
                          First(&bagian).Error

                if err != nil {
                    if err == gorm.ErrRecordNotFound {
            return fmt.Errorf("bagian_tubuh_tidak_valid")
                    }
                    return err
                    }

                  kondisi.IDBagian = bagian.IDBagian

    

            }
        }

        // === SIMPAN KONDISI USER ===
        if err := tx.Create(&kondisi).Error; err != nil {
            return err
        }

        // === VALIDASI NYERI BERAT ===
        if kondisi.TingkatNyeri > 7 {
            return fmt.Errorf("nyeri_berat")
        }

        // === AMBIL LATIHAN SESUAI BAGIAN ===
        var daftarLatihan []latihan.Latihan
        if err := tx.Where("id_bagian = ?", kondisi.IDBagian).
            Find(&daftarLatihan).Error; err != nil {
            return err
        }

        // === BUAT JADWAL OTOMATIS ===
        for _, lat := range daftarLatihan {
            for day := 1; day <= 3; day++ {
                jadwal := latihan.JadwalLatihanUser{
                    IDJadwal:  uuid.New().String()[:4],
                    IDUser:    kondisi.IDUser,
                    IDLatihan: lat.IDLatihan,
                    Tanggal:   time.Now().AddDate(0, 0, day),
                    Status:    "pending",
                    CreatedAt: time.Now(),
                }
                if err := tx.Create(&jadwal).Error; err != nil {
                    return err
                }
            }
        }

        return nil
    })

    // === RESPONSE ===
    if err != nil {
        if err.Error() == "nyeri_berat" {
            c.JSON(http.StatusOK, gin.H{
                "status":  "refer_to_doctor",
                "message": "Tingkat nyeri Anda kategori berat. Mohon konsultasi ke dokter spesialis.",
            })
            return
        }

        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Gagal memproses data: " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Assessment berhasil. Jadwal latihan telah dibuat otomatis.",
        "data":    kondisi,
    })
}