package controllers

import (
	"net/http"
	"school-manage/config"
	"school-manage/models"

	"github.com/gin-gonic/gin"
)

// GetUsers 获取用户列表
func GetUsers(c *gin.Context) {
	role := c.Query("role")
	var users []models.User
	query := config.DB.Model(&models.User{})
	
	if role != "" {
		query = query.Where("role = ?", role)
	}

	query.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}
