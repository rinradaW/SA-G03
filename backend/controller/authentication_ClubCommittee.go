package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rinradaW/SA-G03/entity"
	"github.com/rinradaW/SA-G03/service"
	"golang.org/x/crypto/bcrypt"
)

// POST /login
func LoginByClubCommittee(c *gin.Context) {
	var payload LoginPayload
	var clubcommittee entity.ClubCommittee

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา user ด้วย ID_Student ที่ผู้ใช้กรอกเข้ามา
	if err := entity.DB().Raw("SELECT * FROM club_committees WHERE id_student = ?", payload.StudentId).Scan(&clubcommittee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบรหัสผ่าน
	err := bcrypt.CompareHashAndPassword([]byte(clubcommittee.Password), []byte(payload.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user credentials"})
		return
	}

	// กำหนดค่า SecretKey, Issuer และระยะเวลาหมดอายุของ Token สามารถกำหนดเองได้
	// SecretKey ใช้สำหรับการ sign ข้อความเพื่อบอกว่าข้อความมาจากตัวเราแน่นอน
	// Issuer เป็น unique id ที่เอาไว้ระบุตัว client
	// ExpirationHours เป็นเวลาหมดอายุของ token

	jwtWrapper := service.JwtWrapper{
		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(clubcommittee.ID_Student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	tokenResponse := LoginResponse{
		Token: signedToken,
		ID:    clubcommittee.ID,
		Stdid: clubcommittee.ID_Student,
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}
