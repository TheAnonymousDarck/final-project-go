package database

import (
	"github.com/TheAnonymousDarck/final-project-go/cmd/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := config.DBUser + ":" + config.DBPassword + "@tcp(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	log.Println("Conectado a la base de datos MySQL")
}
