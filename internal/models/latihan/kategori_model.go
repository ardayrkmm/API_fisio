package latihan

import "time"

type Category struct {
	IDKategori   uint      `gorm:"primaryKey;autoIncrement" json:"idkategori"`
	Nama         string    `gorm:"type:varchar(100);not null" json:"nama"`
	CreatedStamp time.Time `gorm:"autoCreateTime" json:"createdstamp"`
}