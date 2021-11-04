package controller

import (
	"net/http"

	"github.com/rinradaW/SA-G03/entity"
	"github.com/gin-gonic/gin"
)

// POST /budgetcategory
func CreateBudgetCategory(c *gin.Context) {
	var budgetcategory entity.BudgetCategory
	if err := c.ShouldBindJSON(&budgetcategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&budgetcategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": budgetcategory})
}

// GET /budgetcategory/:id
func GetBudgetCategory(c *gin.Context) {
	var budgetcategory entity.BudgetCategory
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM budget_categories WHERE id = ?", id).Scan(&budgetcategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": budgetcategory})
}

// GET /budgetcategories
func ListBudgetCategories(c *gin.Context) {
	var budgetcategories []entity.BudgetCategory
	if err := entity.DB().Raw("SELECT * FROM budget_categories").Scan(&budgetcategories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": budgetcategories})
}

// DELETE /budgetcategories/:id
func DeleteBudgetCategory(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM budget_categories WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "budgetcategory not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /budgetcategory
func UpdateBudgetCategory(c *gin.Context) {
	var budgetcategory entity.BudgetCategory
	if err := c.ShouldBindJSON(&budgetcategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", budgetcategory.ID).First(&budgetcategory); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "budgetcategory not found"})
		return
	}

	if err := entity.DB().Save(&budgetcategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": budgetcategory})
}
