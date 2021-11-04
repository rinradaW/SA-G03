package controller


import (

    "net/http"

    "github.com/rinradaW/SA-G03/entity"
    "github.com/gin-gonic/gin"


)

//POST /join_activity_histories
func CreateJoinActivityHistory(c *gin.Context) {

	var joinactivityhistory entity.JoinActivityHistory
	var activity 			entity.Activity 
	var student				entity.Student
	var editor		        entity.ClubCommittee 

	//result from pushing button will be *bind* into joinactivityhistory
	if err := c.ShouldBindJSON(&joinactivityhistory); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
    }

    // 8: search *activity* with id
    if tx := entity.DB().Where("id = ?", joinactivityhistory.ActivityID).First(&activity); tx.RowsAffected == 0 {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "activity not found"})
    	return
    }

    
    // 9: search *student* with id
    if tx := entity.DB().Where("id = ?", joinactivityhistory.StudentID).First(&student); tx.RowsAffected == 0 {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "student not found"})
    	return
    }

    // 10: search *editor* with id
    if tx := entity.DB().Where("id = ?", joinactivityhistory.EditorID).First(&editor); tx.RowsAffected == 0 {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "editor not found"})
    	return
    }

    // 11: create JoinActivityHistory
    jah := entity.JoinActivityHistory{
    	Activity:      activity,            // โยงความสัมพันธ์กับ Entity Activity
    	Student:       student,             // โยงความสัมพันธ์กับ Entity Playlist
    	HourCount:	   joinactivityhistory.HourCount,	 // ตั้งค่าฟิลด์ Hourcount
    	Point:		   joinactivityhistory.Point,		 // ตั้งค่าฟิลด์ Point
    	Editor: editor,       // โยงความสัมพันธ์กับ Entity ClubCommittee
    	Timestamp:     joinactivityhistory.Timestamp,    // ตั้งค่าฟิลด์ Timestamp
    }

    // 12: save
    if err := entity.DB().Create(&jah).Error; err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
    }
    c.JSON(http.StatusOK, gin.H{"data": jah})

}

// GET /joinactivityhistories/:id
func GetJoinActivityHistory(c *gin.Context) {
    var joinactivityhistory entity.JoinActivityHistory
    id := c.Param("id")
    if err := entity.DB().Preload("Activity").Preload("Student").Preload("Editor").Raw("SELECT * FROM join_activity_histories WHERE id = ?", id).Find(&joinactivityhistory).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": joinactivityhistory})

}

// GET /join_activity_histories
func ListJoinActivityHistories(c *gin.Context) {
    var joinactivityhistories []entity.JoinActivityHistory
    if err := entity.DB().Preload("Activity").Preload("Student").Preload("Editor").Raw("SELECT * FROM join_activity_histories").Find(&joinactivityhistories).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": joinactivityhistories})
}

// DELETE /join_activity_histories/:id
func DeleteJoinActivityHistory(c *gin.Context) {
    id := c.Param("id")
    if tx := entity.DB().Exec("DELETE FROM join_activity_histories WHERE id = ?", id); tx.RowsAffected == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "joinactivityhistory not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /join_activity_histories
func UpdateJoinActivityHistory(c *gin.Context) {
    var joinactivityhistory entity.JoinActivityHistory
    if err := c.ShouldBindJSON(&joinactivityhistory); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if tx := entity.DB().Where("id = ?", joinactivityhistory.ID).First(&joinactivityhistory); tx.RowsAffected == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "joinactivityhistory not found"})
        return
    }

    if err := entity.DB().Save(&joinactivityhistory).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": joinactivityhistory})
}


















