package main

import (
	"github.com/TheAnonymousDarck/final-project-go/cmd/config"
	"github.com/TheAnonymousDarck/final-project-go/database"
	"github.com/TheAnonymousDarck/final-project-go/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Cargar configuración
	config.LoadConfig()

	// Conectar a la base de datos
	database.ConnectDatabase()

	// Iniciar Gin
	router := gin.Default()

	//router.LoadHTMLGlob("cmd/templates/**/*")

	router.Use(func(c *gin.Context) {
		c.Set("db", database.DB)
		c.Next()
	})

	// Definir las rutas
	routes.SetupRoutes(router)

	// Ejecutar el servidor
	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
