package users

type QuestionOption struct {
	ID         string `gorm:"primaryKey;size:10" json:"id"`
	QuestionID string `json:"question_id"`
	Label      string `json:"label"`
}