package latihan

import (
	"time"
)


type JadwalLatihanUser struct {
    IDJadwal string    `json:"id_jadwal" gorm:"column:id_jadwal;primaryKey;size:4"`
    IDUser   string    `gorm:"foreignkey:id_user" json:"id_user"`
    Tanggal  time.Time `json:"tanggal" gorm:"column:tanggal"`
    IDLatihan  string    `json:"id_latihan"` 
    CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
    Status      string    `json:"status" gorm:"column:status"`

    
}
