package models

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"unique;not null;size:50" json:"username"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Email        string    `gorm:"unique;size:100" json:"email"`
	Phone        string    `gorm:"size:20" json:"phone"`
	Role         string    `gorm:"type:enum('admin','teacher','student');default:'student'" json:"role"`
	Status       string    `gorm:"type:enum('active','inactive');default:'active'" json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
