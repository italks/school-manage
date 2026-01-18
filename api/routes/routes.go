package routes

import (
	"school-manage/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		users := api.Group("/users")
		{
			users.GET("", controllers.GetUsers)
		}

		resources := api.Group("/resources")
		{
			resources.GET("/classrooms", controllers.GetClassrooms)
			resources.POST("/classrooms", controllers.CreateClassroom)
			resources.PUT("/classrooms/:id", controllers.UpdateClassroom)
			resources.DELETE("/classrooms/:id", controllers.DeleteClassroom)
		}

		schedules := api.Group("/schedules")
		{
			schedules.GET("", controllers.GetSchedules)
			schedules.POST("", controllers.CreateSchedule)
			schedules.DELETE("/:id", controllers.DeleteSchedule)
		}
	}
}
