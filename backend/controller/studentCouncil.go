package controller

import (
	"net/http"

	"github.com/rinradaW/SA-G03/entity"
	"github.com/gin-gonic/gin"
)

// POST /studentCouncil
func CreateStudentCouncil(c *gin.Context) {
	var studentCouncil entity.StudentCouncil
	if err := c.ShouldBindJSON(&studentCouncil); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&studentCouncil).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": studentCouncil})
}

// GET /studentCouncil/:student_id
func GetStudentCouncilWithStudentID(c *gin.Context) {
	var studentCouncil entity.StudentCouncil
	student_id := c.Param("ID_Student")
	if err := entity.DB().Raw("SELECT * FROM student_councils WHERE ID_Student = ?", student_id).Scan(&studentCouncil).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": studentCouncil})
}

// GET /studentCouncil/:id
func GetStudentCouncil(c *gin.Context) {
	var studentCouncil entity.StudentCouncil
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM student_councils WHERE id = ?", id).Scan(&studentCouncil).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": studentCouncil})
}

// GET /studentCouncils
func ListStudentCouncils(c *gin.Context) {
	var studentCouncils []entity.StudentCouncil
	if err := entity.DB().Raw("SELECT * FROM student_councils").Scan(&studentCouncils).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": studentCouncils})
}

// DELETE /studentCouncils/:id
func DeleteStudentCouncil(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM student_councils WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Student Council not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /studentCouncils
func UpdateStudentCouncil(c *gin.Context) {
	var studentCouncil entity.StudentCouncil
	if err := c.ShouldBindJSON(&studentCouncil); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", studentCouncil.ID).First(&studentCouncil); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Student Council not found"})
		return
	}

	if err := entity.DB().Save(&studentCouncil).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": studentCouncil})
}
