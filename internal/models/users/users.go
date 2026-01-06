package users

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
    IDUser           string    `json:"id_user" gorm:"column:id_user;primaryKey;size:4"`
    Nama             string    `json:"nama" gorm:"column:nama"`
    Email            string    `json:"email" gorm:"column:email"`
    Role             string    `json:"role" gorm:"column:role"`
    NoTelepon        string    `json:"no_telepon" gorm:"column:no_telepon"`
    VerifikasiStatus int       `json:"verifikasi_status" gorm:"column:verifikasi_status"`
    Password         string    `json:"password" gorm:"column:password"`
    CreatedAt        time.Time `json:"created_at" gorm:"column:created_at"`

   
}

type PublicUser struct {
	IDUser    string `json:"id_user"`
	Nama      string `json:"nama"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	NoTelepon string    `json:"no_telepon"`
}
func (u *User) CheckPassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}


func (u *User) ToPublicUser() PublicUser {
	return PublicUser{
		IDUser:    u.IDUser,
		Nama:      u.Nama,
		Email:     u.Email,
		Role:      u.Role,
		NoTelepon: u.NoTelepon,
	}
}
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// 1. Generate 4 Karakter Acak untuk IDUser (Hexadecimal)
	// Kita ambil 2 byte karena 1 byte = 2 karakter hex
	b := make([]byte, 2)
	if _, err := rand.Read(b); err != nil {
		return err
	}
	u.IDUser = hex.EncodeToString(b) // Menghasilkan contoh: "a1b2", "3f4e"

	// 2. Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return nil
}