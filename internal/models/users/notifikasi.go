package users

import "time"

type Notifikasi struct {
    IDNotifikasi string    `json:"id_notifikasi" gorm:"column:id_notifikasi;primaryKey;size:36"`
    IDUser       string    `gorm:"foreignkey:id_user" json:"id_user"`
    Judul        string    `json:"judul" gorm:"column:judul"`
    Pesan        string    `json:"pesan" gorm:"column:pesan;type:text"`
    Tipe         string    `json:"tipe" gorm:"column:tipe"`
    StatusBaca   bool      `json:"status_baca" gorm:"column:status_baca"`
    CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`

    
}
