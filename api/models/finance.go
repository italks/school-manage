package models

import "time"

type StudentWallet struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	StudentID     uint      `gorm:"unique;not null" json:"student_id"`
	Student       User      `gorm:"foreignKey:StudentID" json:"student"`
	Balance       float64   `gorm:"type:decimal(10,2);default:0.00" json:"balance"`
	TotalRecharge float64   `gorm:"type:decimal(10,2);default:0.00" json:"total_recharge"`
	TotalConsume  float64   `gorm:"type:decimal(10,2);default:0.00" json:"total_consume"`
	UpdatedAt     time.Time `json:"updated_at"`
}
