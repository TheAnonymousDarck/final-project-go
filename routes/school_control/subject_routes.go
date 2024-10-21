package school_control

import (
	"github.com/TheAnonymousDarck/final-project-go/controllers"
	"github.com/gin-gonic/gin"
)

func SubjectRoutes(r *gin.Engine) {
	subjectGroup := r.Group("/api/subjects")
	{
		// Crear materia
		subjectGroup.POST("", controllers.CreateSubjectHandler)

		// Obtener materia por ID
		subjectGroup.GET("/:subject_id", controllers.GetSubjectHandler)

		// Actualizar materia
		subjectGroup.PUT("/:subject_id", controllers.UpdateSubjectHandler)

		// Eliminar materia
		subjectGroup.DELETE("/:subject_id", controllers.DeleteSubjectHandler)
	}
}
