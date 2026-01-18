package controllers

import (
	"net/http"
	"school-manage/config"
	"school-manage/models"
	"time"

	"github.com/gin-gonic/gin"
)

// GetSchedules 获取排课列表
func GetSchedules(c *gin.Context) {
	var schedules []models.Schedule
	// Preload associations
	config.DB.Preload("Teacher").Preload("Classroom").Find(&schedules)
	c.JSON(http.StatusOK, gin.H{"data": schedules})
}

// CreateScheduleInput 创建排课输入结构
type CreateScheduleInput struct {
	TeacherID   uint      `json:"teacher_id" binding:"required"`
	ClassroomID uint      `json:"classroom_id" binding:"required"`
	CourseName  string    `json:"course_name" binding:"required"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	MaxStudents int       `json:"max_students"`
}

// CreateSchedule 创建排课
func CreateSchedule(c *gin.Context) {
	var input CreateScheduleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 冲突检测 - 老师
	var teacherCount int64
	config.DB.Model(&models.Schedule{}).
		Where("teacher_id = ? AND status != 'cancelled'", input.TeacherID).
		Where("NOT (start_time >= ? OR end_time <= ?)", input.EndTime, input.StartTime).
		Count(&teacherCount)

	if teacherCount > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Teacher schedule conflict"})
		return
	}

	// 冲突检测 - 教室
	var classroomCount int64
	config.DB.Model(&models.Schedule{}).
		Where("classroom_id = ? AND status != 'cancelled'", input.ClassroomID).
		Where("NOT (start_time >= ? OR end_time <= ?)", input.EndTime, input.StartTime).
		Count(&classroomCount)

	if classroomCount > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Classroom schedule conflict"})
		return
	}

	schedule := models.Schedule{
		TeacherID:   input.TeacherID,
		ClassroomID: input.ClassroomID,
		CourseName:  input.CourseName,
		StartTime:   input.StartTime,
		EndTime:     input.EndTime,
		MaxStudents: input.MaxStudents,
		Status:      "scheduled",
	}

	if schedule.MaxStudents == 0 {
		schedule.MaxStudents = 20
	}

	if err := config.DB.Create(&schedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": schedule})
}

// DeleteSchedule 删除/取消排课
func DeleteSchedule(c *gin.Context) {
	id := c.Param("id")
	var schedule models.Schedule

	if err := config.DB.First(&schedule, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	schedule.Status = "cancelled"
	config.DB.Save(&schedule)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
