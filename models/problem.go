package models

import "time"

type Sheet struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" gorm:"not null"`
	Slug        string  `json:"slug" gorm:"uniqueIndex;not null"`
	Description string  `json:"description"`
	Order       int     `json:"order" gorm:"not null"`
	Topics      []Topic `json:"topics,omitempty" gorm:"foreignKey:SheetID"`
}

type Topic struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	SheetID  uint      `json:"sheet_id" gorm:"not null;index"`
	Name     string    `json:"name" gorm:"not null"`
	Order    int       `json:"order" gorm:"not null"`
	Problems []Problem `json:"problems,omitempty" gorm:"foreignKey:TopicID"`
}

type Problem struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	TopicID    uint   `json:"topic_id" gorm:"not null;index"`
	Title      string `json:"title" gorm:"not null"`
	URL        string `json:"url"`
	Difficulty string `json:"difficulty" gorm:"not null"`
	Order      int    `json:"order" gorm:"not null"`
}

type UserProgress struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	UserID    uint       `json:"user_id" gorm:"uniqueIndex:idx_user_problem;not null"`
	ProblemID uint       `json:"problem_id" gorm:"uniqueIndex:idx_user_problem;not null"`
	Status    string     `json:"status" gorm:"default:unsolved"`
	Notes     string     `json:"notes"`
	SolvedAt  *time.Time `json:"solved_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Problem   Problem    `json:"problem,omitempty" gorm:"foreignKey:ProblemID"`
}
