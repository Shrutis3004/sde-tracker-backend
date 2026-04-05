package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"sde-tracker-backend/database"
	"sde-tracker-backend/models"

	"github.com/gin-gonic/gin"
)

type MockInterviewProblem struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	URL        string `json:"url"`
	Difficulty string `json:"difficulty"`
	TopicName  string `json:"topic_name"`
}

func GetMockInterview(c *gin.Context) {
	difficulty := c.Query("difficulty") // Easy, Medium, Hard, or Mixed
	countStr := c.DefaultQuery("count", "5")
	topicName := c.Query("topic") // optional topic filter

	count, err := strconv.Atoi(countStr)
	if err != nil || count < 1 || count > 20 {
		count = 5
	}

	var problems []MockInterviewProblem

	query := database.DB.Table("problems").
		Select("problems.id, problems.title, problems.url, problems.difficulty, topics.name as topic_name").
		Joins("LEFT JOIN topics ON topics.id = problems.topic_id")

	if difficulty != "" && difficulty != "Mixed" {
		query = query.Where("problems.difficulty = ?", difficulty)
	}

	if topicName != "" && topicName != "All" {
		query = query.Where("topics.name LIKE ?", "%"+topicName+"%")
	}

	// Get all matching problems
	var allProblems []MockInterviewProblem
	query.Find(&allProblems)

	if len(allProblems) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"problems": []MockInterviewProblem{},
			"count":    0,
			"duration": 0,
		})
		return
	}

	// Shuffle and pick
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(allProblems), func(i, j int) {
		allProblems[i], allProblems[j] = allProblems[j], allProblems[i]
	})

	if count > len(allProblems) {
		count = len(allProblems)
	}
	problems = allProblems[:count]

	// Calculate suggested duration based on difficulty
	totalMinutes := 0
	for _, p := range problems {
		switch p.Difficulty {
		case "Easy":
			totalMinutes += 10
		case "Medium":
			totalMinutes += 15
		case "Hard":
			totalMinutes += 25
		default:
			totalMinutes += 15
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"problems": problems,
		"count":    len(problems),
		"duration": totalMinutes,
	})
}

// GetTopicsList returns distinct topic names for the interview picker
func GetTopicsList(c *gin.Context) {
	var topics []string
	database.DB.Model(&models.Topic{}).Distinct("name").Order("name").Pluck("name", &topics)
	c.JSON(http.StatusOK, topics)
}
