package handlers

import (
	"net/http"

	"sde-tracker-backend/database"
	"sde-tracker-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetSheets returns all sheets
func GetSheets(c *gin.Context) {
	var sheets []models.Sheet
	database.DB.Order("[order]").Find(&sheets)
	c.JSON(http.StatusOK, sheets)
}

// GetSheetTopics returns all topics with problems for a specific sheet
func GetSheetTopics(c *gin.Context) {
	sheetSlug := c.Param("slug")

	var sheet models.Sheet
	if result := database.DB.Where("slug = ?", sheetSlug).First(&sheet); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sheet not found"})
		return
	}

	difficulty := c.Query("difficulty")

	var topics []models.Topic
	database.DB.Where("sheet_id = ?", sheet.ID).Preload("Problems", func(db *gorm.DB) *gorm.DB {
		if difficulty != "" {
			db = db.Where("difficulty = ?", difficulty)
		}
		return db.Order("[order]")
	}).Order("[order]").Find(&topics)

	c.JSON(http.StatusOK, gin.H{
		"sheet":  sheet,
		"topics": topics,
	})
}

// GetTopics returns all topics with their problems (legacy - returns all)
func GetTopics(c *gin.Context) {
	var topics []models.Topic
	database.DB.Preload("Problems", func(db *gorm.DB) *gorm.DB {
		return db.Order("[order]")
	}).Order("[order]").Find(&topics)
	c.JSON(http.StatusOK, topics)
}

// GetTopicProblems returns problems for a specific topic
func GetTopicProblems(c *gin.Context) {
	topicID := c.Param("id")

	var problems []models.Problem
	database.DB.Where("topic_id = ?", topicID).Order("[order]").Find(&problems)
	c.JSON(http.StatusOK, problems)
}
