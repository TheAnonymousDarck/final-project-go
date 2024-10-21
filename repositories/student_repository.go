package repositories

import (
	"github.com/TheAnonymousDarck/final-project-go/models"
	"gorm.io/gorm"
)

func CreateStudent(db *gorm.DB, student *models.Student) error {
	return db.Create(student).Error
}

func GetStudent(db *gorm.DB, studentID int) (models.Student, error) {
	var student models.Student
	err := db.First(&student, studentID).Error
	return student, err
}

func GetAllStudents(db *gorm.DB) ([]models.Student, error) {
	var students []models.Student
	err := db.Find(&students).Error
	return students, err
}

func UpdateStudent(db *gorm.DB, student *models.Student) error {
	return db.Save(student).Error
}

func DeleteStudent(db *gorm.DB, studentID int) error {
	return db.Delete(&models.Student{}, studentID).Error
}
