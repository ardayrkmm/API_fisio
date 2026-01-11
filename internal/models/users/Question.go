package users

import "time"

type Question struct {
	ID          string           `gorm:"primaryKey;size:10" json:"id"`
	Title       string           `json:"title"`
	Subtitle    string           `json:"subtitle"`
	MultiSelect bool             `json:"multiSelect"`
	TargetField string `json:"target_field"`
	Options     []QuestionOption `gorm:"foreignKey:QuestionID" json:"options"`
	CreatedAt   time.Time        `json:"created_at"`
}