package latihan

type JadwalLatihanDetail struct {
	IDDetail    string `json:"id_detail" gorm:"column:id_detail;primaryKey;size:36"`
	IDJadwal    string `gorm:"foreignkey:id_jadwal" json:"id_jadwal"`
	IDListVideo string `gorm:"foreignkey:id_list_video" json:"id_list_video"`
	Urutan      int    `json:"urutan" gorm:"column:urutan"`
}
