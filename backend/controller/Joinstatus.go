package controller

import (
	"net/http"

	"github.com/rinradaW/SA-G03/entity"
	"github.com/gin-gonic/gin"
)

// POST /joinstatuss
func CreateJoinstatus(c *gin.Context) {
	var joinstatus entity.Joinstatus
	if err := c.ShouldBindJSON(&joinstatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&joinstatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": joinstatus})
}

// GET /joinstatus/:id
func GetJoinstatus(c *gin.Context) {
	var joinstatus entity.Joinstatus
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM joinstatuses WHERE id = ?", id).Scan(&joinstatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": joinstatus})
}

// GET /joinstatuses
func ListJoinstatuses(c *gin.Context) {
	var joinstatuss []entity.Joinstatus
	if err := entity.DB().Raw("SELECT * FROM joinstatuses").Scan(&joinstatuss).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": joinstatuss})
}

// DELETE /joinstatuss/:id
func DeleteJoinstatus(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM joinstatuses WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "joinstatus not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /joinstatuss
func UpdateJoinstatus(c *gin.Context) {
	var joinstatus entity.Joinstatus
	if err := c.ShouldBindJSON(&joinstatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", joinstatus.ID).First(&joinstatus); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "joinstatus not found"})
		return
	}

	if err := entity.DB().Save(&joinstatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": joinstatus})
}
