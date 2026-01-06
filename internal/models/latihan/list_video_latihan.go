package latihan

import (
	"time"
)

type ListVideoLatihan struct {
    IDListVideo    string    `json:"id_list_video" gorm:"column:id_list_video;primaryKey;size:4"`
    IDLatihan      string    `gorm:"foreignkey:id_latihan" json:"id_latihan"`
    NamaGerakan    string    `json:"nama_gerakan" gorm:"column:nama_gerakan"`
    VideoURL       string    `json:"video_url" gorm:"column:video_url"`
    TargetSet      int       `json:"target_set" gorm:"column:target_set"`
    TargetRepetisi int       `json:"target_repetisi" gorm:"column:target_repetisi"`
    TargetWaktu    float64   `json:"target_waktu" gorm:"column:target_waktu;type:double"`
    CreatedAt      time.Time `json:"created_at" gorm:"column:created_at"`

  
}
