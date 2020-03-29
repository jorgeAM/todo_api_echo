package models

import "time"

// User model
type User struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	FirstName       string    `json:"firstName" gorm:"column:firstName;not null"`
	LastName        string    `json:"lastName" gorm:"column:lastName;not null"`
	Email           string    `json:"email" gorm:"unique;not null"`
	Password        string    `json:"password,omitempty" gorm:"not null"`
	ConfirmPassword string    `json:"confirmPassword,omitempty" gorm:"-"`
	CreatedAt       time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt       time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}
