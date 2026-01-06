package users

type BagianTubuh struct {
	IDBagian   string `json:"id_bagian" gorm:"column:id_bagian;primaryKey;size:4"`
	NamaBagian string `json:"nama_bagian" gorm:"column:nama_bagian"`
}
