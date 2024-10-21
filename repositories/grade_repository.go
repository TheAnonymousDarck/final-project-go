package repositories

import (
	"github.com/TheAnonymousDarck/final-project-go/models"
	"gorm.io/gorm"
)

func CreateGrade(db *gorm.DB, grade *models.Grade) error {
	return db.Create(grade).Error
}

func GetGradeByStudentAndSubject(db *gorm.DB, gradeID, studentID int) (models.Grade, error) {
	var grade models.Grade
	err := db.Where("grade_id = ? AND student_id = ?", gradeID, studentID).First(&grade).Error
	return grade, err
}

func GetAllGradesByStudent(db *gorm.DB, studentID int) ([]models.Grade, error) {
	var grades []models.Grade
	err := db.Where("student_id = ?", studentID).Find(&grades).Error
	return grades, err
}

func UpdateGrade(db *gorm.DB, grade *models.Grade) error {
	return db.Save(grade).Error
}

func DeleteGrade(db *gorm.DB, gradeID int) error {
	return db.Delete(&models.Grade{}, gradeID).Error
}
