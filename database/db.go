package database

import (
	"github.com/TheAnonymousDarck/final-project-go/cmd/config"
	"github.com/TheAnonymousDarck/final-project-go/models"
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

	migrateModels()
}

func migrateModels() {
	if err := DB.AutoMigrate(&models.Student{}, &models.Subject{},&models.Grade{}  ); err != nil {
        log.Fatalf("Error al migrar la base de datos: %v", err)
    }
	// err := DB.AutoMigrate(&models.Grade{},&models.Student{}, &models.Subject{} )
	// err := db.AutoMigrate(&models.Student{})
	// if err != nil {
	// 	log.Fatalf("Error al migrar la base de datos: %v", err)
	// }
	// err := db.AutoMigrate(&models.Grade{})
	// if err != nil {
	// 	log.Fatalf("Error al migrar la base de datos: %v", err)
	// }
	// err := db.AutoMigrate(&models.Subject{})
	// if err != nil {
	// 	log.Fatalf("Error al migrar la base de datos: %v", err)
	// }
}
