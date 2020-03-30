package models

import "time"

// Todo model
type Todo struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	UserID      *uint     `json:"-" gorm:"column:userId"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}
