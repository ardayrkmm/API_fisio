package artikel

import (
	"time"
)

type ArtikelView struct {
	IDView    string    `json:"id_view" gorm:"column:id_view;primaryKey;size:4"`
	IDArtikel string    `gorm:"foreignkey:id_artikel" json:"id_artikel"`
	IDUser    string    `json:"id_user" gorm:"column:id_user;size:36"`
	ViewedAt  time.Time `json:"viewed_at" gorm:"column:viewed_at"`
}
