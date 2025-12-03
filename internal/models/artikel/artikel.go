package artikel

import "time"
type Artikel struct {
    IDArtikel  string    `json:"id_artikel" gorm:"column:id_artikel;primaryKey;size:36"`
    Judul      string    `json:"judul" gorm:"column:judul"`
    Deskripsi  string    `json:"deskripsi" gorm:"column:deskripsi"`
    IDTags     string    `json:"id_tags" gorm:"column:id_tags"`
    CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`

   
}
