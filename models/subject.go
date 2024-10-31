package models

type Subject struct {
    SubjectID int     `gorm:"primaryKey;autoIncrement"`
    Name      string  

    // Relación uno a muchos con Grade
    // Grades []Grade `gorm:"foreignKey:SubjectID"`
}
