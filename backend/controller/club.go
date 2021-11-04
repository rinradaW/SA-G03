package controller

import (
	"net/http"

	"github.com/rinradaW/SA-G03/entity"
	"github.com/gin-gonic/gin"
)

// POST /club
func CreateClub(c *gin.Context) {
	var club entity.Club
	var adviser entity.Teacher
	var adder entity.StudentCouncil
	var typeClub entity.TypeClub

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 7 จะถูก bind เข้าตัวแปร Club
	if err := c.ShouldBindJSON(&club); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//8 type club
	if tx := entity.DB().Where("id = ?", club.TypeClubID).First(&typeClub); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type club not found"})
		return
	}

	//9 teacher
	if tx := entity.DB().Where("id = ?", club.AdviserID).First(&adviser); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "adviser not found"})
		return
	}

	//10 Student Council
	if tx := entity.DB().Where("id = ?", club.AdderID).First(&adder); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "student council not found"})
		return
	}

	//11 create
	newClub := entity.Club{
		Adder:    adder,
		Adviser:  adviser,
		TypeClub: typeClub,
		Name:     club.Name,
	}

	//12 save
	if err := entity.DB().Create(&newClub).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": newClub})
}

// GET /club/:id
func GetClub(c *gin.Context) {
	var club entity.Club
	id := c.Param("id")
	if err := entity.DB().Preload("Adder").Preload("Adviser").Preload("TypeClub").Raw("SELECT * FROM clubs WHERE id = ?", id).Find(&club).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": club})
}

// GET /clubs
func ListClubs(c *gin.Context) {
	var clubs []entity.Club
	if err := entity.DB().Preload("Adder").Preload("Adviser").Preload("TypeClub").Raw("SELECT * FROM clubs").Find(&clubs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": clubs})
}

// DELETE /club/:id
func DeleteClub(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM clubs WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "teacher not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
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


// GET /clubs/student_id/:id --By Ohm--
func ListClubByStudentID(c *gin.Context) {
	var clubs []entity.Club
	id := c.Param("id")

	// ค้นหา club ทั้งหมดที่นักษาที่กำลังใช้งานระบบไม่ได้เป็นสมาชิกอยู่ 
	if err := entity.DB().Raw("SELECT * FROM clubs WHERE id NOT IN (SELECT club_id FROM club_memberships WHERE student_id = ?)", id).Find(&clubs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorishere": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": clubs})

}

// GET /clubname/clubcommittee/:id <- Parameter --by bank
func GetClubwithClubCommittee(c *gin.Context) {
	var clubcommittee entity.ClubCommittee
	var club entity.Club
	id := c.Param("id")

	if tx := entity.DB().Raw("SELECT * FROM club_committees WHERE id = ?", id).Find(&clubcommittee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "clubcommittees not found"})
		//  find club_committees from id
		return
	}
	if err := entity.DB().Raw("SELECT * FROM clubs WHERE id = ?", clubcommittee.ClubID).Find(&club).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//  ค้นหาข้างบนเจอ ใช้ clubcommittee.ClubID ค้นหา club
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": club})
}