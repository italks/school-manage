package models

import "time"

type Classroom struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:50;not null" json:"name"`
	Capacity  int       `gorm:"default:30" json:"capacity"`
	Equipment string    `gorm:"type:text" json:"equipment"`
	Status    string    `gorm:"type:enum('available','maintenance','closed');default:'available'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
