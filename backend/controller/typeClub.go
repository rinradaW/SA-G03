package controller

import (
	"net/http"

	"github.com/rinradaW/SA-G03/entity"
	"github.com/gin-gonic/gin"
)

// POST /typeClubs
func CreateTypeClub(c *gin.Context) {
	var typeClub entity.TypeClub
	if err := c.ShouldBindJSON(&typeClub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&typeClub).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": typeClub})
}

// GET /typeClub/:id
func GetTypeClub(c *gin.Context) {
	var typeClub entity.TypeClub
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM type_clubs WHERE id = ?", id).Scan(&typeClub).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": typeClub})
}

// GET /typeClubs
func ListTypeClubs(c *gin.Context) {
	var typeClubs []entity.TypeClub
	if err := entity.DB().Raw("SELECT * FROM type_clubs").Scan(&typeClubs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": typeClubs})
}

// DELETE /typeClubs/:id
func DeleteTypeClub(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM type_clubs WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "typeClub not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /typeClubs
func UpdateTypeClub(c *gin.Context) {
	var typeClub entity.TypeClub
	if err := c.ShouldBindJSON(&typeClub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", typeClub.ID).First(&typeClub); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "typeClub not found"})
		return
	}

	if err := entity.DB().Save(&typeClub).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": typeClub})
}