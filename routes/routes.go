package routes

import (
	"github.com/TheAnonymousDarck/final-project-go/routes/school_control"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/")
	// Otras rutas...

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	school_control.StudentRoutes(router)
	school_control.SubjectRoutes(router)
	school_control.GradeRoutes(router)
}
