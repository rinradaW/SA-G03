package controller

import (
	"net/http"

	"github.com/rinradaW/SA-G03/entity"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
)

// POST /clubcommittees
func CreateClubCommittee(c *gin.Context) {
	var clubcommittee entity.ClubCommittee
	if err := c.ShouldBindJSON(&clubcommittee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(clubcommittee.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}
	clubcommittee.Password = string(bytes)


	if err := entity.DB().Create(&clubcommittee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": clubcommittee})
}

// GET /clubcommittees
// List all clubcommittees
func ListClubCommittees(c *gin.Context) {
	var clubcommittees []entity.ClubCommittee
	if err := entity.DB().Raw("SELECT * FROM club_committees").Scan(&clubcommittees).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": clubcommittees})
}

// GET /clubcommittee/:id
// Get clubcommittee by id
func GetClubCommittee(c *gin.Context) {
	var clubcommittee entity.ClubCommittee
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM club_committees WHERE id = ?", id).Scan(&clubcommittee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": clubcommittee})
}

// GET /club_committees/std/:ID_Student
func ClubCommitteefromstudentid(c *gin.Context) {
	var clubcommittee entity.ClubCommittee
	student := c.Param("ID_Student") //user.go
	if err := entity.DB().Preload("Club").Raw("SELECT * FROM club_committees WHERE ID_Student = ?", student).Find(&clubcommittee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} //preload object Club ตารางของ Club ใน type ClubCommittee struct
	c.JSON(http.StatusOK, gin.H{"data": clubcommittee})
}

// PATCH /clubcommittees
func UpdateClubCommittee(c *gin.Context) {
	var clubcommittee entity.ClubCommittee
	if err := c.ShouldBindJSON(&clubcommittee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", clubcommittee.ID).First(&clubcommittee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "clubcommittee not found"})
		return
	}

	if err := entity.DB().Save(&clubcommittee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": clubcommittee})
}

// DELETE /club_committees/:id
func DeleteClubCommittee(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM club_committees WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "club_committee not found"})
		return
	}
	/*
		if err := entity.DB().Where("id = ?", id).Delete(&entity.ClubCommittee{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"data": id})
}