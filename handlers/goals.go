package handlers

import (
	"net/http"
	"time"

	"sde-tracker-backend/database"
	"sde-tracker-backend/models"

	"github.com/gin-gonic/gin"
)

type SetGoalRequest struct {
	Target int `json:"target" binding:"required,min=1"`
}

// GetGoal returns the user's daily goal target
func GetGoal(c *gin.Context) {
	userID := c.GetUint("user_id")

	var goal models.DailyGoal
	result := database.DB.Where("user_id = ?", userID).First(&goal)
	if result.RowsAffected == 0 {
		// Return default
		c.JSON(http.StatusOK, gin.H{"target": 5})
		return
	}

	c.JSON(http.StatusOK, goal)
}

// SetGoal sets the user's daily goal target
func SetGoal(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req SetGoalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var goal models.DailyGoal
	result := database.DB.Where("user_id = ?", userID).First(&goal)

	if result.RowsAffected == 0 {
		goal = models.DailyGoal{
			UserID: userID,
			Target: req.Target,
		}
		database.DB.Create(&goal)
	} else {
		goal.Target = req.Target
		database.DB.Save(&goal)
	}

	c.JSON(http.StatusOK, goal)
}

// GetHeatmap returns the last 365 days of DailySolve data
func GetHeatmap(c *gin.Context) {
	userID := c.GetUint("user_id")

	startDate := time.Now().AddDate(0, 0, -365).Format("2006-01-02")

	var solves []models.DailySolve
	database.DB.Where("user_id = ? AND date >= ?", userID, startDate).
		Order("date").
		Find(&solves)

	c.JSON(http.StatusOK, solves)
}

// IncrementDailySolve increments today's solve count for a user.
// Called internally when a problem is marked as solved.
func IncrementDailySolve(userID uint) {
	today := time.Now().Format("2006-01-02")

	var solve models.DailySolve
	result := database.DB.Where("user_id = ? AND date = ?", userID, today).First(&solve)

	if result.RowsAffected == 0 {
		solve = models.DailySolve{
			UserID: userID,
			Date:   today,
			Count:  1,
		}
		database.DB.Create(&solve)
	} else {
		solve.Count++
		database.DB.Save(&solve)
	}
}
