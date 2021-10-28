package controller

import (
	"net/http"

	"github.com/rinradaW/SA-G03/entity"
	"github.com/gin-gonic/gin"
)

// POST /clubs
func CreateClub(c *gin.Context) {
	var club entity.Club
	if err := c.ShouldBindJSON(&club); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&club).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": club})
}

// GET /clubs
// List all clubs
func ListClubs(c *gin.Context) {
	var clubs []entity.Club
	if err := entity.DB().Raw("SELECT * FROM clubs").Scan(&clubs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": clubs})
}

// GET /club/:id
// Get club by id
func GetClub(c *gin.Context) {
	var club entity.Club
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM clubs WHERE id = ?", id).Scan(&club).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": club})
}

// PATCH /clubs
func UpdateClub(c *gin.Context) {
	var club entity.Club
	if err := c.ShouldBindJSON(&club); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", club.ID).First(&club); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "club not found"})
		return
	}

	if err := entity.DB().Save(&club).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": club})
}

// DELETE /clubs/:id
func DeleteClub(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM clubs WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "club not found"})
		return
	}
	/*
		if err := entity.DB().Where("id = ?", id).Delete(&entity.Club{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"data": id})
}