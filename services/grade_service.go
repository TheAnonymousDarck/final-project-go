package services

import (
	"github.com/TheAnonymousDarck/final-project-go/models"
	"github.com/TheAnonymousDarck/final-project-go/repositories"
	"gorm.io/gorm"
)

// CreateGrade Crear calificación
func CreateGrade(db *gorm.DB, grade *models.Grade) error {
	return repositories.CreateGrade(db, grade)
}

// UpdateGrade Actualizar calificación
func UpdateGrade(db *gorm.DB, gradeID int, grade *models.Grade) (*models.Grade, error) {
	// Buscar la calificación existente
	existingGrade, err := repositories.GetGradeByStudentAndSubject(db, gradeID, grade.StudentID)
	if err != nil {
		return nil, err
	}

	// Actualizar los campos
	existingGrade.SubjectID = grade.SubjectID
	existingGrade.Grade = grade.Grade

	// Guardar los cambios
	if err := repositories.UpdateGrade(db, &existingGrade); err != nil {
		return nil, err
	}

	return &existingGrade, nil
}

// DeleteGrade Eliminar calificación por ID
func DeleteGrade(db *gorm.DB, gradeID int) error {
	return repositories.DeleteGrade(db, gradeID)
}

// GetGrade Obtener una calificación por ID y estudiante
func GetGrade(db *gorm.DB, gradeID, studentID int) (*models.Grade, error) {
	grade, err := repositories.GetGradeByStudentAndSubject(db, gradeID, studentID)
	if err != nil {
		return nil, err
	}
	return &grade, nil
}

// GetAllGradesForStudent Obtener todas las calificaciones de un estudiante
func GetAllGradesForStudent(db *gorm.DB, studentID int) ([]models.Grade, error) {
	return repositories.GetAllGradesByStudent(db, studentID)
}
