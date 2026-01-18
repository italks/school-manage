package models

import (
	"time"
)

type Schedule struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TeacherID   uint      `gorm:"not null" json:"teacher_id"`
	Teacher     User      `gorm:"foreignKey:TeacherID" json:"teacher"`
	ClassroomID uint      `gorm:"not null" json:"classroom_id"`
	Classroom   Classroom `gorm:"foreignKey:ClassroomID" json:"classroom"`
	CourseName  string    `gorm:"size:100;not null" json:"course_name"`
	StartTime   time.Time `gorm:"not null;index" json:"start_time"`
	EndTime     time.Time `gorm:"not null;index" json:"end_time"`
	MaxStudents int       `gorm:"default:20" json:"max_students"`
	Status      string    `gorm:"type:enum('scheduled','ongoing','completed','cancelled');default:'scheduled'" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type Booking struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	StudentID   uint      `gorm:"not null" json:"student_id"`
	Student     User      `gorm:"foreignKey:StudentID" json:"student"`
	ScheduleID  uint      `gorm:"not null" json:"schedule_id"`
	Schedule    Schedule  `gorm:"foreignKey:ScheduleID" json:"schedule"`
	BookingTime time.Time `gorm:"autoCreateTime" json:"booking_time"`
	Status      string    `gorm:"type:enum('booked','attended','absent','cancelled');default:'booked'" json:"status"`
	Cost        float64   `gorm:"type:decimal(10,2);default:0.00" json:"cost"`
	CheckinTime *time.Time `json:"checkin_time"`
}
