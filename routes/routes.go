package routes

import (
	"sde-tracker-backend/config"
	"sde-tracker-backend/handlers"
	"sde-tracker-backend/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine, cfg *config.Config) {
	authHandler := handlers.NewAuthHandler(cfg)

	api := r.Group("/api")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)
	}

	protected := api.Group("/")
	protected.Use(middleware.AuthRequired(cfg))
	{
		protected.GET("/profile", authHandler.GetProfile)

		// Sheets
		protected.GET("/sheets", handlers.GetSheets)
		protected.GET("/sheets/:slug", handlers.GetSheetTopics)
		protected.GET("/sheets/:slug/stats", handlers.GetSheetStats)

		// Topics & Problems
		protected.GET("/topics", handlers.GetTopics)
		protected.GET("/topics/:id/problems", handlers.GetTopicProblems)

		// Progress
		protected.GET("/progress", handlers.GetProgress)
		protected.PUT("/progress", handlers.UpdateProgress)

		// Stats
		protected.GET("/stats", handlers.GetStats)

		// Bookmarks
		protected.GET("/bookmarks", handlers.GetBookmarks)
		protected.POST("/bookmarks", handlers.AddBookmark)
		protected.DELETE("/bookmarks/:id", handlers.DeleteBookmark)

		// Goals
		protected.GET("/goals", handlers.GetGoal)
		protected.PUT("/goals", handlers.SetGoal)
		protected.GET("/goals/heatmap", handlers.GetHeatmap)

		// Search
		protected.GET("/search", handlers.SearchProblems)

		// Hints
		protected.GET("/hints/:id", handlers.GetHints)

		// Mock Interview
		protected.GET("/mock-interview", handlers.GetMockInterview)
		protected.GET("/topics-list", handlers.GetTopicsList)
	}
}
