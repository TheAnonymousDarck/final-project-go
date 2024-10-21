package models

type Student struct {
	StudentID int    `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Group     string `gorm:"size:50"`
	Email     string `gorm:"size:100"`
}
