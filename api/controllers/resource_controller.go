package controllers

import (
	"net/http"
	"school-manage/config"
	"school-manage/models"

	"github.com/gin-gonic/gin"
)

// GetClassrooms 获取所有教室
func GetClassrooms(c *gin.Context) {
	var classrooms []models.Classroom
	config.DB.Find(&classrooms)
	c.JSON(http.StatusOK, gin.H{"data": classrooms})
}

// CreateClassroom 创建教室
func CreateClassroom(c *gin.Context) {
	var input models.Classroom
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	classroom := models.Classroom{
		Name:      input.Name,
		Capacity:  input.Capacity,
		Equipment: input.Equipment,
		Status:    "available",
	}

	config.DB.Create(&classroom)
	c.JSON(http.StatusOK, gin.H{"data": classroom})
}

// UpdateClassroom 更新教室
func UpdateClassroom(c *gin.Context) {
	id := c.Param("id")
	var classroom models.Classroom

	if err := config.DB.First(&classroom, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Classroom not found"})
		return
	}

	var input models.Classroom
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&classroom).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": classroom})
}

// DeleteClassroom 删除教室
func DeleteClassroom(c *gin.Context) {
	id := c.Param("id")
	var classroom models.Classroom

	if err := config.DB.First(&classroom, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Classroom not found"})
		return
	}

	config.DB.Delete(&classroom)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
