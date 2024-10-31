package models

type Grade struct {
    GradeID   int     `gorm:"primaryKey;autoIncrement" json:"grade_id"`   // Asegúrate de que el json:"grade_id" sea el mismo que en tu JSON
    StudentID int     `json:"student_id"`  // Asegúrate de que el json:"student_id" sea el mismo que en tu JSON
    SubjectID int     `json:"subject_id"`  // Asegúrate de que el json:"subject_id" sea el mismo que en tu JSON
    Grade     float32 `json:"grade"` 

    // Relaciones
    // Student Student `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    // Subject Subject `gorm:"foreignKey:SubjectID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
