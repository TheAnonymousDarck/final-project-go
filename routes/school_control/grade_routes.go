package school_control

import (
	"github.com/TheAnonymousDarck/final-project-go/controllers"
	"github.com/gin-gonic/gin"
)

func GradeRoutes(r *gin.Engine) {
	gradeGroup := r.Group("/api/grades")
	{
		// Crear calificación
		gradeGroup.POST("", controllers.CreateGradeHandler)

		// Actualizar calificación
		gradeGroup.PUT("/:grade_id", controllers.UpdateGradeHandler)

		// Eliminar calificación
		gradeGroup.DELETE("/:grade_id", controllers.DeleteGradeHandler)

		// Obtener una calificación específica de un estudiante
		gradeGroup.GET("/:grade_id/student/:student_id", controllers.GetGradeHandler)

		// Obtener todas las calificaciones de un estudiante
		gradeGroup.GET("/student/:student_id", controllers.GetAllGradesForStudentHandler)
	}
}
