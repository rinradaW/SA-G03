package controller

import(
	"github.com/rinradaW/SA-G03/entity"
	"github.com/gin-gonic/gin"
	"net/http"
) 

// POST /authority
func CreateAuthority(c *gin.Context) {
	var authority entity.Authority

	if err := c.ShouldBindJSON(&authority); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return	
	}

	if err := entity.DB().Create(&authority).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": authority})
}

// GET /authority/:id
func GetAuthority(c *gin.Context) {
	var authority entity.Authority
	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM authorities WHERE id = ?", id).Scan(&authority).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": authority})
}

// GET /authorities
func ListAuthorities(c *gin.Context) {
	var authorities []entity.Authority

	if err := entity.DB().Raw("SELECT * FROM authorities").Scan(&authorities).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": authorities})
}

// DELETE /authority
func DeleteAuthority(c *gin.Context) {
	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM authorities WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "authority not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /authority
func UpdateAuthority(c *gin.Context) {
	var authority entity.Authority

	if err := c.ShouldBindJSON(&authority); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", authority.ID).First(&authority); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "authority not found"})
		return
	}

	if err := entity.DB().Save(&authority).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": authority})
}