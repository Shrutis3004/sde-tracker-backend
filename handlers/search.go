package handlers

import (
	"net/http"

	"sde-tracker-backend/database"
	"sde-tracker-backend/models"

	"github.com/gin-gonic/gin"
)

type SearchResult struct {
	models.Problem
	TopicName string `json:"topic_name"`
}

// SearchProblems searches problems by title (case-insensitive)
func SearchProblems(c *gin.Context) {
	q := c.Query("q")
	if q == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	var results []SearchResult
	database.DB.Raw(`
		SELECT p.*, t.name as topic_name
		FROM problems p
		JOIN topics t ON t.id = p.topic_id
		WHERE LOWER(p.title) LIKE LOWER(?)
		ORDER BY p.title
	`, "%"+q+"%").Scan(&results)

	c.JSON(http.StatusOK, results)
}
