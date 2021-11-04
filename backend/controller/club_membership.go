package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rinradaW/SA-G03/entity"
)

// POST /club_membership
func CreateClubMembership(c *gin.Context) {
	var student entity.Student
	var authority entity.Authority
	var membershipStatus entity.MembershipStatus
	var club entity.Club
	var clubMembership entity.ClubMembership

	// 8:บันทึกข้อมูลเข้าชมรม() ไม่ได้เรียกใช้ตรงๆ แต่เป็นการบันทึกใส่ตัวแปรเพื่อที่ได้นำไปบันทึกใน DB ต่อไป
	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร clubMmebership
	if err := c.ShouldBindJSON(&clubMembership); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา student ด้วย id
	if tx := entity.DB().Where("id = ?", clubMembership.StudentID).First(&student); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "student not found"})
		return
	}

	// 10: ค้นหา authority ด้วย id
	if tx := entity.DB().Where("id = ?", clubMembership.AuthorityID).First(&authority); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "authority not found"})
		return
	}

	// 11: ค้นหา club ด้วย id
	if tx := entity.DB().Where("id = ?", clubMembership.ClubID).First(&club); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "club not found"})
		return
	}

	// 12: ค้นหา membershipStatus ด้วย id
	if tx := entity.DB().Where("id = ?", clubMembership.MembershipStatusID).First(&membershipStatus); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "memebrship status not found"})
		return
	}

	// 13: สร้าง ClubMembership
	cm := entity.ClubMembership{
		Student:          student,                     // โยงความสัมพันธ์กับ Entity Student
		Authority:        authority,                   // โยงความสัมพันธ์กับ Entity Authority
		Club:             club,                        // โยงความสัมพันธ์กับ Entity Club
		MembershipStatus: membershipStatus,            // โยงความสัมพันธ์กับ Entity MembershipStatus
		RegisterDate:     clubMembership.RegisterDate, // ตั้งค่าฟิลด์ registerDate
	}

	// 14: บันทึก()
	if err := entity.DB().Save(&cm).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cm})
}

// GET /club_membership/:id
func GetClubMembership(c *gin.Context) {
	var clubMembership entity.ClubMembership
	id := c.Param("id")

	// preload เป็นการดึงข้อมูลมาจาก FK ของตารางที่ใช้ในคำสั่ง Raw
	if err := entity.DB().Preload("MembershipStatus").Preload("Authority").Preload("Club").Preload("Student").Raw("SELECT * FROM club_memberships WHERE id = ?", id).Find(&clubMembership).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorishere": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": clubMembership})

}

// GET /club_memberships
func ListClubMemberships(c *gin.Context) {
	var clubMembership []entity.ClubMembership

	if err := entity.DB().Preload("MembershipStatus").Preload("Authority").Preload("Club").Preload("Student").Raw("SELECT * FROM club_memberships").Find(&clubMembership).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": clubMembership})

}

// GET /club_memberships/pending
func ListMembershipStatusesPending(c *gin.Context) {
	var clubMembership []entity.ClubMembership

	if err := entity.DB().Preload("MembershipStatus").Preload("Authority").Preload("Club").Preload("Student").Raw("SELECT * FROM club_memberships WHERE membership_status_id = 1").Find(&clubMembership).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": clubMembership})
}

// DELETE /club_membership/:id
func DeleteClubMembership(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM club_memberships WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "club membership not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /club_membership
func UpdateClubMembership(c *gin.Context) {
	var newData entity.ClubMembership
	var oldData entity.ClubMembership

	if err := c.ShouldBindJSON(&newData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", newData.ID).First(&oldData); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "club membership not found"})
		return
	}

	if err := entity.DB().Save(&newData).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newData})
}
