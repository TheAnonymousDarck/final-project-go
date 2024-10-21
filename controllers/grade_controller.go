package controllers

import (
	"net/http"
	"strconv"

	"github.com/TheAnonymousDarck/final-project-go/models"
	"github.com/TheAnonymousDarck/final-project-go/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateGradeHandler Maneja la creación de una calificación
func CreateGradeHandler(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var grade models.Grade
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateGrade(db, &grade); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la calificación"})
		return
	}

	c.JSON(http.StatusOK, grade)
}

// UpdateGradeHandler Maneja la actualización de una calificación
func UpdateGradeHandler(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	gradeID, err := strconv.Atoi(c.Param("grade_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de calificación inválido"})
		return
	}

	var grade models.Grade
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedGrade, err := services.UpdateGrade(db, gradeID, &grade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la calificación"})
		return
	}

	c.JSON(http.StatusOK, updatedGrade)
}

// DeleteGradeHandler Maneja la eliminación de una calificación
func DeleteGradeHandler(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	gradeID, err := strconv.Atoi(c.Param("grade_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de calificación inválido"})
		return
	}

	if err := services.DeleteGrade(db, gradeID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la calificación"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Calificación eliminada exitosamente"})
}

// GetGradeHandler Maneja la obtención de una calificación
func GetGradeHandler(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	gradeID, err := strconv.Atoi(c.Param("grade_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de calificación inválido"})
		return
	}

	studentID, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de estudiante inválido"})
		return
	}

	grade, err := services.GetGrade(db, gradeID, studentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Calificación no encontrada"})
		return
	}

	c.JSON(http.StatusOK, grade)
}

// GetAllGradesForStudentHandler Maneja la obtención de todas las calificaciones de un estudiante
func GetAllGradesForStudentHandler(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	studentID, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de estudiante inválido"})
		return
	}

	grades, err := services.GetAllGradesForStudent(db, studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las calificaciones"})
		return
	}

	c.JSON(http.StatusOK, grades)
}
