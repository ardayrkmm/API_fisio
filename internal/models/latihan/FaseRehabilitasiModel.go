package latihan

import (
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type FaseRehabilitasi struct {
    IDFase    string    `json:"id_fase" gorm:"primaryKey;size:4"`
    NamaFase  string    `json:"nama_fase"` 
    // Akut, Subakut, Remodelling, Fungsional
    CreatedAt time.Time `json:"created_at"`
}

var charset = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func generateRandom4Char() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 4)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func (f *FaseRehabilitasi) BeforeCreate(tx *gorm.DB) error {
	if f.IDFase == "" {
		for {
			id := generateRandom4Char()

			var count int64
			tx.Model(&FaseRehabilitasi{}).
				Where("id_fase = ?", id).
				Count(&count)

			if count == 0 {
				f.IDFase = id
				break
			}
		}
	}
	return nil
}