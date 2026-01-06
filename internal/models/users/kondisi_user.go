package users

import "time"

type KondisiUser struct {
	IDForm        string    `json:"id_form" gorm:"column:id_form;primaryKey;size:4"`
	IDUser        string    `gorm:"foreignkey:id_user" json:"id_user"`
	IDBagian      string    `gorm:"foreignkey:id_bagian" json:"id_bagian"`
	LamaNyeriHari int       `json:"lama_nyeri_hari" gorm:"column:lama_nyeri_hari"`
	TingkatNyeri  int       `json:"tingkat_nyeri" gorm:"column:tingkat_nyeri"`
	Catatan       string    `json:"catatan" gorm:"column:catatan;type:text"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`

	
}
