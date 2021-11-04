package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rinradaW/SA-G03/entity"
)

// POST /reservestatuss
func CreateReserveStatus(c *gin.Context) {
	var reservestatus entity.ReserveStatus
	if err := c.ShouldBindJSON(&reservestatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&reservestatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": reservestatus})
}

// GET /reservestatus/:id
func GetReserveStatus(c *gin.Context) {
	var reservestatus entity.ReserveStatus
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM reserve_statuses WHERE id = ?", id).Scan(&reservestatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reservestatus})
}

// GET /reservestatuss
func ListReserveStatuss(c *gin.Context) {
	var reservestatus []entity.ReserveStatus
	if err := entity.DB().Raw("SELECT * FROM reserve_statuses").Scan(&reservestatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reservestatus})
}

// DELETE /reservestatuss/:id
func DeleteReserveStatus(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM reserve_statuses WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "reserve_statuses not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /reservestatuss
func UpdateReserveStatus(c *gin.Context) {
	var reservestatus entity.ReserveStatus
	if err := c.ShouldBindJSON(&reservestatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", reservestatus.ID).First(&reservestatus); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "reserve_statuses not found"})
		return
	}

	if err := entity.DB().Save(&reservestatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reservestatus})
}
