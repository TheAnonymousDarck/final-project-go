package services

import (
	"github.com/TheAnonymousDarck/final-project-go/models"
	"github.com/TheAnonymousDarck/final-project-go/repositories"
	"gorm.io/gorm"
)

// CreateSubject Crea una nueva materia
func CreateSubject(db *gorm.DB, subject *models.Subject) error {
	return repositories.CreateSubject(db, subject)
}

// GetSubject Obtiene una materia por su ID
func GetSubject(db *gorm.DB, subjectID int) (models.Subject, error) {
	return repositories.GetSubject(db, subjectID)
}

// UpdateSubject Actualiza una materia existente
func UpdateSubject(db *gorm.DB, subject *models.Subject) error {
	return repositories.UpdateSubject(db, subject)
}

// DeleteSubject Elimina una materia por su ID
func DeleteSubject(db *gorm.DB, subjectID int) error {
	return repositories.DeleteSubject(db, subjectID)
}
