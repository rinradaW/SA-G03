package controller

import (
	"net/http"

	"github.com/rinradaW/SA-G03/entity"
	"github.com/gin-gonic/gin"
)

// POST /watch_activitys
func CreateJoining(c *gin.Context) {

	var joining entity.Joining
	var joinstatus entity.Joinstatus
	var student entity.Student
	var activity entity.Activity

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 9 จะถูก bind เข้าตัวแปร joining
	if err := c.ShouldBindJSON(&joining); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 10: ค้นหา student ด้วย id
	if tx := entity.DB().Where("id = ?", joining.StudentID).First(&student); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "student not found"})
		return
	}

	// 11: ค้นหา activity ด้วย id
	if tx := entity.DB().Where("id = ?", joining.ActivityID).First(&activity); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "activity not found"})
		return
	}
	// 12: ค้นหา joinstatus ด้วย id
	if tx := entity.DB().Where("id = ?", joining.JoinstatusID).First(&joinstatus); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "joinstatus not found"})
		return
	}

	// 13: สร้าง Joining
	wv := entity.Joining{
		Joinstatus:   joinstatus,           // โยงความสัมพันธ์กับ Entity Joinstatus
		Activity:     activity,             // โยงความสัมพันธ์กับ Entity Activity
		Student:      student,              // โยงความสัมพันธ์กับ Entity Student
		Joining_time: joining.Joining_time, // ตั้งค่าฟิลด์ DateTime
	}

	// 14: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /joinings/:id
func GetJoining(c *gin.Context) {
	var joining entity.Joining
	id := c.Param("id")
	if err := entity.DB().Preload("Joinstatus").Preload("Student").Preload("Activity").Raw("SELECT * FROM joinings WHERE id = ?", id).Find(&joining).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": joining})
}

// GET /joinings
func ListJoinings(c *gin.Context) {
	var joinings []entity.Joining
	if err := entity.DB().Preload("Joinstatus").Preload("Student").Preload("Activity").Raw("SELECT * FROM joinings").Find(&joinings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": joinings})
}

// DELETE /joinings/:id
func DeleteJoining(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM joinings WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "joining not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /watch_activitys
func UpdateJoining(c *gin.Context) {
	var joining entity.Joining
	if err := c.ShouldBindJSON(&joining); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", joining.ID).First(&joining); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "joining not found"})
		return
	}

	if err := entity.DB().Save(&joining).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": joining})
}
