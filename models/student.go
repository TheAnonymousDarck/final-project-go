package models

type Student struct {
    StudentID int     `gorm:"primaryKey;autoIncrement"`
    Name      string  
    Group     string  
    Email     string  

    // Relación uno a muchos con Grade
    // Grades []Grade `gorm:"foreignKey:StudentID"`
}
