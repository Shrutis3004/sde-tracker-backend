package models

type DailyGoal struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id" gorm:"uniqueIndex;not null"`
	Target int  `json:"target" gorm:"default:5"`
}

type DailySolve struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	UserID uint   `json:"user_id" gorm:"index;not null"`
	Date   string `json:"date" gorm:"index;not null"` // YYYY-MM-DD
	Count  int    `json:"count" gorm:"default:0"`
}
