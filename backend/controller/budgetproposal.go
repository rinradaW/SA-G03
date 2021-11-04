package controller

import (
	"net/http"

	"github.com/rinradaW/SA-G03/entity"
	"github.com/gin-gonic/gin"
)

// POST /budgetproposal
func CreateBudgetProposal(c *gin.Context) {
	var activity entity.Activity
	var budgetproposal entity.BudgetProposal
	var budgettype entity.BudgetType
	var budgetcategory entity.BudgetCategory

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร budgetproposal
	if err := c.ShouldBindJSON(&budgetproposal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา activity ด้วย id
	if tx := entity.DB().Where("id = ?", budgetproposal.ActivityID).First(&activity); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "activity not found"})
		return
	}



	// 11: ค้นหา budgettype ด้วย id
	if tx := entity.DB().Where("id = ?", budgetproposal.BudgetTypeID).First(&budgettype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "budgettype not found"})
		return
	}

	// 11: ค้นหา budgetcategory ด้วย id
	if tx := entity.DB().Where("id = ?", budgetproposal.BudgetCategoryID).First(&budgetcategory); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "budgetcategory not found"})
		return
	}


	// 12: สร้าง budgetproposal
	bp := entity.BudgetProposal{      
		Activity:			activity,               
		BudgetType:    		budgettype,               
		BudgetCategory: 	budgetcategory, 
		BudgetPrice:    	budgetproposal.BudgetPrice,
	}

	// 13: บันทึก
	if err := entity.DB().Create(&bp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bp})
}

// GET /budgetproposal/:id
func GetBudgetProposal(c *gin.Context) {
	var budgetproposal entity.BudgetProposal
	id := c.Param("id")
	if err := entity.DB().Preload("Activity").Preload("BudgetType").Preload("BudgetCategory").Raw("SELECT * FROM budget_proposal WHERE id = ?", id).Find(&budgetproposal).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": budgetproposal})
}

// GET /budget_proposal
func ListBudgetProposals(c *gin.Context) {
	var budgetproposals []entity.BudgetProposal
	if err := entity.DB().Preload("Activity").Preload("BudgetType").Preload("BudgetCategory").Raw("SELECT * FROM budget_proposals").Find(&budgetproposals).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": budgetproposals})
}

// DELETE /budget_proposals/:id
func DeleteBudgetProposal(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM budget_proposals WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "budgetproposal not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /budget_proposal
func UpdateBudgetProposal(c *gin.Context) {
	var budgetproposal entity.BudgetProposal
	if err := c.ShouldBindJSON(&budgetproposal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", budgetproposal.ID).First(&budgetproposal); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "budgetproposal not found"})
		return
	}

	if err := entity.DB().Save(&budgetproposal).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": budgetproposal})
}