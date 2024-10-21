package repositories

import (
	"github.com/TheAnonymousDarck/final-project-go/models"
	"gorm.io/gorm"
)

func CreateSubject(db *gorm.DB, subject *models.Subject) error {
	return db.Create(subject).Error
}

func GetSubject(db *gorm.DB, subjectID int) (models.Subject, error) {
	var subject models.Subject
	err := db.First(&subject, subjectID).Error
	return subject, err
}

func UpdateSubject(db *gorm.DB, subject *models.Subject) error {
	return db.Save(subject).Error
}

func DeleteSubject(db *gorm.DB, subjectID int) error {
	return db.Delete(&models.Subject{}, subjectID).Error
}
