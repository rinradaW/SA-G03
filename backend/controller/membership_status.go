package controller 

import(
	"github.com/rinradaW/SA-G03/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

// POST /membership_status
func CreateMembershipStatus(c *gin.Context) {
	var membershipStatus entity.MembershipStatus

	if err := c.ShouldBindJSON(&membershipStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
		return
	}

	if err := entity.DB().Create(&membershipStatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": membershipStatus})
}

// GET /membership_status/:id
func GetMembershipStatus(c *gin.Context) {
	var membershipStatus entity.MembershipStatus
	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM membership_statuses WHERE id = ?", id).Scan(&membershipStatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": membershipStatus})	
}

// GET /membership_statuses
func ListMembershipStatuses(c *gin.Context) {
	var membershipStatus []entity.MembershipStatus

	if err := entity.DB().Raw("SELECT * FROM membership_statuses").Scan(&membershipStatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": membershipStatus})	
}


// DELETE /membership_status
func DeleteMembershipStatus(c *gin.Context) {
	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM membership_statuses WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "membership not found"})	
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})	
}

// PATCH /membership_status
func UpdateMembershipStatus(c *gin.Context) {
	var membershipStatus entity.MembershipStatus

	if err := c.ShouldBindJSON(&membershipStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
		return
	}

	if tx := entity.DB().Where("id = ?", membershipStatus.ID).First(&membershipStatus); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "membership not found"})	
		return
	}

	if err := entity.DB().Save(&membershipStatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": membershipStatus})	

}