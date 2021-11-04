package controller

import (
	"net/http"

	"github.com/rinradaW/SA-G03/entity"
	"github.com/gin-gonic/gin"
)

// POST /budgettype
func CreateBudgetType(c *gin.Context) {
	var budgettype entity.BudgetType
	if err := c.ShouldBindJSON(&budgettype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&budgettype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": budgettype})
}

// GET /budgettype/:id
func GetBudgetType(c *gin.Context) {
	var budgettype entity.BudgetType
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM budget_types WHERE id = ?", id).Scan(&budgettype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": budgettype})
}

// GET /budgettypes
func ListBudgetTypes(c *gin.Context) {
	var budgettypes []entity.BudgetType
	if err := entity.DB().Raw("SELECT * FROM budget_types").Scan(&budgettypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": budgettypes})
}

// DELETE /budgettypes/:id
func DeleteBudgetType(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM budget_types WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "budgettype not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /budgettype
func UpdateBudgetType(c *gin.Context) {
	var budgettype entity.BudgetType
	if err := c.ShouldBindJSON(&budgettype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", budgettype.ID).First(&budgettype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "budgettype not found"})
		return
	}

	if err := entity.DB().Save(&budgettype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": budgettype})
}