package services

import (
	"github.com/TheAnonymousDarck/final-project-go/models"
	"gorm.io/gorm"
)

// CreateStudent Crear estudiante
func CreateStudent(db *gorm.DB, student *models.Student) error {
	return db.Create(student).Error
}

// DeleteStudent Eliminar estudiante por ID
func DeleteStudent(db *gorm.DB, studentID string) error {
	return db.Delete(&models.Student{}, studentID).Error
}

// UpdateStudent Actualizar estudiante
func UpdateStudent(db *gorm.DB, studentID string, student *models.Student) (*models.Student, error) {
	var existingStudent models.Student
	if err := db.First(&existingStudent, studentID).Error; err != nil {
		return nil, err
	}

	existingStudent.Name = student.Name
	existingStudent.Group = student.Group
	existingStudent.Email = student.Email

	if err := db.Save(&existingStudent).Error; err != nil {
		return nil, err
	}

	return &existingStudent, nil
}

// GetAllStudents Obtener todos los estudiantes
func GetAllStudents(db *gorm.DB) ([]models.Student, error) {
	var students []models.Student
	if err := db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

// GetStudent Obtener un estudiante por ID
func GetStudent(db *gorm.DB, studentID string) (*models.Student, error) {
	var student models.Student
	if err := db.First(&student, studentID).Error; err != nil {
		return nil, err
	}
	return &student, nil
}
