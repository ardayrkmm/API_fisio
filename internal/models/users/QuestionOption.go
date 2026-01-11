package users

type QuestionOption struct {
	ID         string `gorm:"primaryKey;size:10" json:"id"`
	QuestionID string `json:"question_id"`
	Nilai      int    `json:"nilai"`
	Label      string `json:"label"`
}