package school_control

import (
	"github.com/TheAnonymousDarck/final-project-go/controllers"
	"github.com/gin-gonic/gin"
)

func StudentRoutes(router *gin.Engine) {
	studentRoutes := router.Group("/api/students")
	{
		studentRoutes.POST("", controllers.CreateStudent)
		studentRoutes.GET("", controllers.GetAllStudents)
		studentRoutes.GET("/:student_id", controllers.GetStudent)
		studentRoutes.PUT("/:student_id", controllers.UpdateStudent)
		studentRoutes.DELETE("/:student_id", controllers.DeleteStudent)
	}
}