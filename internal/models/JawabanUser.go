package models

import "time"

type JawabanUser struct {
	ID         string `gorm:"primaryKey;size:10"`
	IDUser     string `json:"id_user"`
	QuestionID string `json:"question_id"`
	OptionID   string `json:"option_id"`
	CreatedAt  time.Time
}
