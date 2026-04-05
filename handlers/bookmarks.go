package handlers

import (
	"net/http"

	"sde-tracker-backend/database"
	"sde-tracker-backend/models"

	"github.com/gin-gonic/gin"
)

type AddBookmarkRequest struct {
	ProblemID uint `json:"problem_id" binding:"required"`
}

// GetBookmarks returns all bookmarks for the authenticated user
func GetBookmarks(c *gin.Context) {
	userID := c.GetUint("user_id")

	var bookmarks []models.Bookmark
	database.DB.Where("user_id = ?", userID).Preload("Problem").Find(&bookmarks)
	c.JSON(http.StatusOK, bookmarks)
}

// AddBookmark creates a bookmark for a problem
func AddBookmark(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req AddBookmarkRequest
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

	// Check if already bookmarked
	var existing models.Bookmark
	if result := database.DB.Where("user_id = ? AND problem_id = ?", userID, req.ProblemID).First(&existing); result.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Problem already bookmarked"})
		return
	}

	bookmark := models.Bookmark{
		UserID:    userID,
		ProblemID: req.ProblemID,
	}
	database.DB.Create(&bookmark)

	database.DB.Preload("Problem").First(&bookmark, bookmark.ID)
	c.JSON(http.StatusCreated, bookmark)
}

// DeleteBookmark removes a bookmark
func DeleteBookmark(c *gin.Context) {
	userID := c.GetUint("user_id")
	bookmarkID := c.Param("id")

	var bookmark models.Bookmark
	if result := database.DB.Where("id = ? AND user_id = ?", bookmarkID, userID).First(&bookmark); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bookmark not found"})
		return
	}

	database.DB.Delete(&bookmark)
	c.JSON(http.StatusOK, gin.H{"message": "Bookmark removed"})
}
