package artikel

import "time"

type GaleriGambar struct {
	IDGambar  string    `json:"id_gambar" gorm:"column:id_gambar;primaryKey;size:4"`
	UrlFile   string    `json:"url_file" gorm:"column:url_file"`
	IDArtikel string    `gorm:"foreignkey:id_artikel" json:"id_artikel"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`

	// Relasi
	
}
