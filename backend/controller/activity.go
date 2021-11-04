package controller

import (
	"net/http"

	"github.com/rinradaW/SA-G03/entity"
	"github.com/gin-gonic/gin"
)

// POST /students
func CreateActivity(c *gin.Context) {
	var activity entity.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&activity).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": activity})
}

// GET /activities
// List all activities
func ListActivities(c *gin.Context) {
	var activities []entity.Activity
	if err := entity.DB().Raw("SELECT * FROM activities").Scan(&activities).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": activities})
}

//clubandactivity
func ClubwithActivity(c *gin.Context){
	var activity []entity.Activity
	ClubID := c.Param("ClubID")
	if err := entity.DB().Preload("Club").Raw("SELECT * FROM activities WHERE club_id = ?", ClubID).Find(&activity).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": activity})
}

// GET /activity/:id
// Get activity by id
func GetActivity(c *gin.Context) {
	var activity entity.Activity
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM activities WHERE id = ?", id).Scan(&activity).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": activity})
}

// PATCH /activities
func UpdateActivity(c *gin.Context) {
	var activity entity.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", activity.ID).First(&activity); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "activity not found"})
		return
	}

	if err := entity.DB().Save(&activity).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": activity})
}

// DELETE /activities/:id
func DeleteActivity(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM activities WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "activity not found"})
		return
	}
	/*
		if err := entity.DB().Where("id = ?", id).Delete(&entity.Activity{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"data": id})
}