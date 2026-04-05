package database

import (
	"log"

	"sde-tracker-backend/models"
)

type problemData struct {
	Title      string
	URL        string
	Difficulty string
}

type topicData struct {
	Name     string
	Problems []problemData
}

type sheetData struct {
	Name        string
	Slug        string
	Description string
	Topics      []topicData
}

func Seed() {
	var count int64
	DB.Model(&models.Sheet{}).Count(&count)
	if count > 0 {
		log.Println("Database already seeded")
		return
	}

	log.Println("Seeding database with Striver's Sheets...")

	sheets := []sheetData{
		sdeSheet(),
		a2zSheet(),
		blind75Sheet(),
		striver79Sheet(),
		cpSheet(),
		cnSheet(),
		dbmsSheet(),
		osSheet(),
		systemDesignSheet(),
	}

	for i, s := range sheets {
		sheet := models.Sheet{
			Name:        s.Name,
			Slug:        s.Slug,
			Description: s.Description,
			Order:       i + 1,
		}
		DB.Create(&sheet)

		for j, t := range s.Topics {
			topic := models.Topic{
				SheetID: sheet.ID,
				Name:    t.Name,
				Order:   j + 1,
			}
			DB.Create(&topic)

			for k, p := range t.Problems {
				problem := models.Problem{
					TopicID:    topic.ID,
					Title:      p.Title,
					URL:        p.URL,
					Difficulty: p.Difficulty,
					Order:      k + 1,
				}
				DB.Create(&problem)
			}
		}
		log.Printf("Seeded %s with %d topics\n", s.Name, len(s.Topics))
	}

	log.Println("All sheets seeded successfully")
}
