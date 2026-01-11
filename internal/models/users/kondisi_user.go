package users

import "time"

type KondisiUser struct {
    IDForm string `json:"id_form" gorm:"column:id_form;primaryKey;size:4"`

    IDUser   string `json:"id_user" gorm:"column:id_user;index"`
    IDBagian string `json:"id_bagian" gorm:"column:id_bagian;index"`

    LamaNyeriHari int `json:"lama_nyeri_hari" gorm:"column:lama_nyeri_hari"`
    TingkatNyeri  int `json:"tingkat_nyeri" gorm:"column:tingkat_nyeri"`

    JenisKeluhan string `json:"jenis_keluhan" gorm:"column:jenis_keluhan"`
    // nyeri | kaku | lemah | tidak_stabil | bengkak

    

    CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}
