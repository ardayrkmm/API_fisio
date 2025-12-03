package users

import (
	"time"
)

type User struct {
    IDUser           string    `json:"id_user" gorm:"column:id_user;primaryKey;size:36"`
    Nama             string    `json:"nama" gorm:"column:nama"`
    Email            string    `json:"email" gorm:"column:email"`
    Role             string    `json:"role" gorm:"column:role"`
    NoTelepon        int       `json:"no_telepon" gorm:"column:no_telepon"`
    VerifikasiStatus int       `json:"verifikasi_status" gorm:"column:verifikasi_status"`
    Password         string    `json:"password" gorm:"column:password"`
    CreatedAt        time.Time `json:"created_at" gorm:"column:created_at"`

   
}
