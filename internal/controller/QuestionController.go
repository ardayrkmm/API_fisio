package controller

import (
	"api_fisioterapi/internal/config"
	models "api_fisioterapi/internal/models/users"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateQuestion(c *gin.Context) {
	var req models.Question

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Payload tidak valid"})
		return
	}

	req.ID = uuid.New().String()[:8]

	for i := range req.Options {
		req.Options[i].ID = uuid.New().String()[:8]
		req.Options[i].QuestionID = req.ID
	}

	if err := config.DB.Create(&req).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, req)
}
func GetQuestions(c *gin.Context) {
	var questions []models.Question

	if err := config.DB.
		Preload("Options").
		Find(&questions).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, questions)
}
func UpdateQuestion(c *gin.Context) {
	id := c.Param("id")

	var question models.Question
	if err := config.DB.First(&question, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Question tidak ditemukan"})
		return
	}

	var req models.Question
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Payload tidak valid"})
		return
	}

	question.Title = req.Title
	question.Subtitle = req.Subtitle
	question.MultiSelect = req.MultiSelect

	config.DB.Where("question_id = ?", id).
		Delete(&models.QuestionOption{})

	for i := range req.Options {
		req.Options[i].ID = uuid.New().String()[:8]
		req.Options[i].QuestionID = id
	}

	config.DB.Save(&question)
	config.DB.Create(&req.Options)

	c.JSON(200, question)
}
func DeleteQuestion(c *gin.Context) {
	id := c.Param("id")

	config.DB.Where("question_id = ?", id).
		Delete(&models.QuestionOption{})

	config.DB.Delete(&models.Question{}, "id = ?", id)

	c.JSON(200, gin.H{"message": "Question berhasil dihapus"})
}
