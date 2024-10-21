package models

type Subject struct {
	SubjectID int    `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
}
