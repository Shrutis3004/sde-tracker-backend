package handlers

import (
	"net/http"
	"time"

	"sde-tracker-backend/database"
	"sde-tracker-backend/models"

	"github.com/gin-gonic/gin"
)

type UpdateProgressRequest struct {
	ProblemID uint   `json:"problem_id" binding:"required"`
	Status    string `json:"status" binding:"required,oneof=unsolved solved revisit"`
	Notes     string `json:"notes"`
}

// GetProgress returns all progress for the authenticated user
func GetProgress(c *gin.Context) {
	userID := c.GetUint("user_id")

	var progress []models.UserProgress
	database.DB.Where("user_id = ?", userID).Preload("Problem").Find(&progress)
	c.JSON(http.StatusOK, progress)
}

// UpdateProgress creates or updates progress for a problem
func UpdateProgress(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req UpdateProgressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check problem exists
	var problem models.Problem
	if result := database.DB.First(&problem, req.ProblemID); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Problem not found"})
		return
	}

	var progress models.UserProgress
	result := database.DB.Where("user_id = ? AND problem_id = ?", userID, req.ProblemID).First(&progress)

	now := time.Now()

	if result.RowsAffected == 0 {
		// Create new progress
		progress = models.UserProgress{
			UserID:    userID,
			ProblemID: req.ProblemID,
			Status:    req.Status,
			Notes:     req.Notes,
		}
		if req.Status == "solved" {
			progress.SolvedAt = &now
		}
		database.DB.Create(&progress)
	} else {
		// Update existing
		progress.Status = req.Status
		progress.Notes = req.Notes
		if req.Status == "solved" && progress.SolvedAt == nil {
			progress.SolvedAt = &now
		}
		database.DB.Save(&progress)
	}

	// Update streak and daily solve count if solved
	if req.Status == "solved" {
		updateStreak(userID)
		IncrementDailySolve(userID)
	}

	database.DB.Preload("Problem").First(&progress, progress.ID)
	c.JSON(http.StatusOK, progress)
}

// GetStats returns dashboard stats for the user
func GetStats(c *gin.Context) {
	userID := c.GetUint("user_id")

	var totalProblems int64
	database.DB.Model(&models.Problem{}).Count(&totalProblems)

	var solved int64
	database.DB.Model(&models.UserProgress{}).Where("user_id = ? AND status = ?", userID, "solved").Count(&solved)

	var revisit int64
	database.DB.Model(&models.UserProgress{}).Where("user_id = ? AND status = ?", userID, "revisit").Count(&revisit)

	// Topic-wise stats
	type TopicStat struct {
		TopicID   uint   `json:"topic_id"`
		TopicName string `json:"topic_name"`
		Total     int64  `json:"total"`
		Solved    int64  `json:"solved"`
	}

	var topicStats []TopicStat
	database.DB.Raw(`
		SELECT t.id as topic_id, t.name as topic_name,
			COUNT(p.id) as total,
			COUNT(CASE WHEN up.status = 'solved' THEN 1 END) as solved
		FROM topics t
		LEFT JOIN problems p ON p.topic_id = t.id
		LEFT JOIN user_progresses up ON up.problem_id = p.id AND up.user_id = ?
		GROUP BY t.id, t.name
		ORDER BY t.[order]
	`, userID).Scan(&topicStats)

	var user models.User
	database.DB.First(&user, userID)

	c.JSON(http.StatusOK, gin.H{
		"total_problems": totalProblems,
		"solved":         solved,
		"revisit":        revisit,
		"unsolved":       totalProblems - solved - revisit,
		"streak":         user.Streak,
		"topic_stats":    topicStats,
	})
}

// GetSheetStats returns stats for a specific sheet
func GetSheetStats(c *gin.Context) {
	userID := c.GetUint("user_id")
	sheetSlug := c.Param("slug")

	var sheet models.Sheet
	if result := database.DB.Where("slug = ?", sheetSlug).First(&sheet); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sheet not found"})
		return
	}

	var totalProblems int64
	database.DB.Raw(`
		SELECT COUNT(p.id) FROM problems p
		JOIN topics t ON t.id = p.topic_id
		WHERE t.sheet_id = ?
	`, sheet.ID).Scan(&totalProblems)

	var solved int64
	database.DB.Raw(`
		SELECT COUNT(up.id) FROM user_progresses up
		JOIN problems p ON p.id = up.problem_id
		JOIN topics t ON t.id = p.topic_id
		WHERE t.sheet_id = ? AND up.user_id = ? AND up.status = 'solved'
	`, sheet.ID, userID).Scan(&solved)

	var revisit int64
	database.DB.Raw(`
		SELECT COUNT(up.id) FROM user_progresses up
		JOIN problems p ON p.id = up.problem_id
		JOIN topics t ON t.id = p.topic_id
		WHERE t.sheet_id = ? AND up.user_id = ? AND up.status = 'revisit'
	`, sheet.ID, userID).Scan(&revisit)

	type TopicStat struct {
		TopicID   uint   `json:"topic_id"`
		TopicName string `json:"topic_name"`
		Total     int64  `json:"total"`
		Solved    int64  `json:"solved"`
	}

	var topicStats []TopicStat
	database.DB.Raw(`
		SELECT t.id as topic_id, t.name as topic_name,
			COUNT(p.id) as total,
			COUNT(CASE WHEN up.status = 'solved' THEN 1 END) as solved
		FROM topics t
		LEFT JOIN problems p ON p.topic_id = t.id
		LEFT JOIN user_progresses up ON up.problem_id = p.id AND up.user_id = ?
		WHERE t.sheet_id = ?
		GROUP BY t.id, t.name
		ORDER BY t.[order]
	`, userID, sheet.ID).Scan(&topicStats)

	var user models.User
	database.DB.First(&user, userID)

	c.JSON(http.StatusOK, gin.H{
		"sheet":          sheet,
		"total_problems": totalProblems,
		"solved":         solved,
		"revisit":        revisit,
		"unsolved":       totalProblems - solved - revisit,
		"streak":         user.Streak,
		"topic_stats":    topicStats,
	})
}

func updateStreak(userID uint) {
	var user models.User
	database.DB.First(&user, userID)

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	if user.LastSolve != nil {
		lastSolveDate := time.Date(user.LastSolve.Year(), user.LastSolve.Month(), user.LastSolve.Day(), 0, 0, 0, 0, user.LastSolve.Location())
		diff := today.Sub(lastSolveDate).Hours() / 24

		if diff == 1 {
			// Consecutive day — increment streak
			user.Streak++
		} else if diff > 1 {
			// Streak broken — reset
			user.Streak = 1
		}
		// diff == 0 means same day, don't change streak
	} else {
		user.Streak = 1
	}

	user.LastSolve = &now
	database.DB.Save(&user)
}
