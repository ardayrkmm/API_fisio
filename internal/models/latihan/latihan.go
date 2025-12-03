package latihan

import "time"

type Latihan struct {
    IDLatihan   string    `json:"id_latihan" gorm:"column:id_latihan;primaryKey;size:36"`
    NamaLatihan string    `json:"nama_latihan" gorm:"column:nama_latihan"`
    IDKategori  string    `json:"id_kategori" gorm:"column:id_kategori"`
    UrlGambar   string    `json:"url_gambar" gorm:"column:url_gambar"`
    Deskripsi   string    `json:"deskripsi" gorm:"column:deskripsi;type:text"`
    CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`

    
}
