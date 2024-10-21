package controllers

import (
	"net/http"
	"strconv"

	"github.com/TheAnonymousDarck/final-project-go/models"
	"github.com/TheAnonymousDarck/final-project-go/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateSubjectHandler Maneja la creación de una nueva materia
func CreateSubjectHandler(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var subject models.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateSubject(db, &subject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la materia"})
		return
	}

	c.JSON(http.StatusOK, subject)
}

// GetSubjectHandler Maneja la obtención de una materia por ID
func GetSubjectHandler(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	subjectID, err := strconv.Atoi(c.Param("subject_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de materia inválido"})
		return
	}

	subject, err := services.GetSubject(db, subjectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Materia no encontrada"})
		return
	}

	c.JSON(http.StatusOK, subject)
}

// UpdateSubjectHandler Maneja la actualización de una materia
func UpdateSubjectHandler(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	subjectID, err := strconv.Atoi(c.Param("subject_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de materia inválido"})
		return
	}

	var subject models.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subject.SubjectID = int(uint(subjectID)) // Asegurarse de asignar el ID correcto

	if err := services.UpdateSubject(db, &subject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la materia"})
		return
	}

	c.JSON(http.StatusOK, subject)
}

// DeleteSubjectHandler Maneja la eliminación de una materia
func DeleteSubjectHandler(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	subjectID, err := strconv.Atoi(c.Param("subject_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de materia inválido"})
		return
	}

	if err := services.DeleteSubject(db, subjectID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la materia"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Materia eliminada exitosamente"})
}
