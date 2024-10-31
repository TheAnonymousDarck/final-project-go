package controllers

import (
	"github.com/TheAnonymousDarck/final-project-go/models"
	"github.com/TheAnonymousDarck/final-project-go/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// CreateStudent Crear un estudiante
func CreateStudent(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateStudent(db, &student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, student)
}

// DeleteStudent Eliminar un estudiante
func DeleteStudent(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	studentID := c.Param("student_id")

	if err := services.DeleteStudent(db, studentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Estudiante eliminado con éxito"})
}

// UpdateStudent Actualizar un estudiante
func UpdateStudent(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	studentID := c.Param("student_id")

	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedStudent, err := services.UpdateStudent(db, studentID, &student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedStudent)
}

// GetAllStudents Obtener todos los estudiantes en formato JSON
func GetAllStudents(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	students, err := services.GetAllStudents(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

// GetStudent Obtener un solo estudiante por ID en formato JSON
func GetStudent(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	studentID := c.Param("student_id")

	student, err := services.GetStudent(db, studentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Estudiante no encontrado"})
		return
	}

	c.JSON(http.StatusOK, student)
}