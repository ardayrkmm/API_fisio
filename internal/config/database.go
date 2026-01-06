package config

import (
	"fmt"
	"log"
	"os"

	artikelModel "api_fisioterapi/internal/models/artikel"
	latihanModel "api_fisioterapi/internal/models/latihan"
	userModel "api_fisioterapi/internal/models/users"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env tidak ditemukan, pakai environment sistem")
	}

	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_CHARSET"),
		os.Getenv("DB_LOC"),
	)


	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Gagal connect ke database:", err)
	}

	log.Println("✅ Database connected")

	
	migrate()
}

func migrate() {
	err := DB.AutoMigrate(
		&userModel.User{},
		&artikelModel.Artikel{},
		&artikelModel.ArtikelView{},
		&artikelModel.GaleriGambar{},
		&userModel.Notifikasi{},
		&userModel.HistoryAktifitas{},
		&userModel.Question{},
		&userModel.QuestionOption{},
		&latihanModel.Latihan{},
		&latihanModel.ListVideoLatihan{},
		&userModel.KondisiUser{},
		&userModel.BagianTubuh{},
		&latihanModel.JadwalLatihanUser{},
		&latihanModel.JadwalLatihanDetail{},
		&latihanModel.Category{},
	)

	if err != nil {
		log.Println("❌ Migrasi gagal:", err)
	} else {
		log.Println("🚀 Migrasi berhasil!")
	}
}