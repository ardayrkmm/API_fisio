package users

import (
	"time"
)


type HistoryAktifitas struct {
	IDHistoryAktifitas string    `json:"id_history_aktifitas" gorm:"column:id_history_aktifitas;primaryKey;size:36"`
	IDLatihan          string    `gorm:"foreignkey:id_latihan" json:"id_latihan"`
	IDUser             string    `gorm:"foreignkey:id_user" json:"id_user"`
	Tanggal            time.Time `json:"tanggal" gorm:"column:tanggal"`
	SetTercapai        int       `json:"set_tercapai" gorm:"column:set_tercapai"`
	RepetisiTercapai   int       `json:"repetisi_tercapai" gorm:"column:repetisi_tercapai"`
	DurasiAktual       float64   `json:"durasi_aktual" gorm:"column:durasi_aktual;type:double"`
	NilaiAkurasi       float64   `json:"nilai_akurasi" gorm:"column:nilai_akurasi;type:double"`
	NilaiLatihan       string    `json:"nilai_latihan" gorm:"column:nilai_latihan"`
	CreatedAt          time.Time `json:"created_at" gorm:"column:created_at"`

	
}
