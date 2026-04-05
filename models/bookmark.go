package models

import "time"

type Bookmark struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"uniqueIndex:idx_user_bookmark;not null"`
	ProblemID uint      `json:"problem_id" gorm:"uniqueIndex:idx_user_bookmark;not null"`
	CreatedAt time.Time `json:"created_at"`
	Problem   Problem   `json:"problem,omitempty" gorm:"foreignKey:ProblemID"`
}
