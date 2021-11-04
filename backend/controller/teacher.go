package controller

import (
	"net/http"

	"github.com/rinradaW/SA-G03/entity"
	"github.com/gin-gonic/gin"
)

// POST /Teachers
func CreateTeacher(c *gin.Context) {
	var teacher entity.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&teacher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": teacher})
}

// GET /teacher/:id
func GetTeacher(c *gin.Context) {
	var teacher entity.Teacher
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM teachers WHERE id = ?", id).Scan(&teacher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": teacher})
}

// GET /teachers
func ListTeachers(c *gin.Context) {
	var teachers []entity.Teacher
	if err := entity.DB().Raw("SELECT * FROM teachers").Scan(&teachers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": teachers})
}

// DELETE /teachers/:id
func DeleteTeacher(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM teachers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "teacher not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /teachers
func UpdateTeacher(c *gin.Context) {
	var teacher entity.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", teacher.ID).First(&teacher); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "teacher not found"})
		return
	}

	if err := entity.DB().Save(&teacher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": teacher})
}