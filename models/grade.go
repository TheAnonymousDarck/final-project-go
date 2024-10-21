package models

type Grade struct {
	GradeID   int `gorm:"primaryKey"`
	StudentID int `gorm:"index"`
	SubjectID int `gorm:"index"`
	Grade     float32
}
