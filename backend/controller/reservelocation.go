package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rinradaW/SA-G03/entity"
)
// POST /reserve_locations
func CreateReserveLocation(c *gin.Context) {

	var reserveLocation entity.ReserveLocation
	var reservestatus entity.ReserveStatus
	var location entity.Location
	var activity entity.Activity
	var clubcommittee entity.ClubCommittee

	if err := c.ShouldBindJSON(&reserveLocation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 10: ค้นหา location ด้วย id
	if tx := entity.DB().Where("id = ?", reserveLocation.LocationID).First(&location); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "location not found"})
		return
	}
	// 11: ค้นหา activity ด้วย id
	if tx := entity.DB().Where("id = ?", reserveLocation.ActivityID).First(&activity); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "activity not found"})
		return
	}

	// 12: ค้นหา ClubCommittee ด้วย id
	if tx := entity.DB().Where("id = ?", reserveLocation.RequestID).First(&clubcommittee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "location not found"})
		return
	}
	// 13: ค้นหา reservestatus ด้วย id
	if tx := entity.DB().Where("id = ?", reserveLocation.ReserveStatusID).First(&reservestatus); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "reservestatus not found"})
		return
	}
	// 12: สร้าง ReserveLocation
	rl := entity.ReserveLocation{

		Location:      location,
		Activity:      activity,
		Request:       clubcommittee,
		ReserveStatus: reservestatus,
		DateStart:     reserveLocation.DateStart,
		DateEnd:       reserveLocation.DateEnd,
	}
	// 13: บันทึก
	if err := entity.DB().Create(&rl).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rl})
}

// GET /reserveLocation/:id
func GetReserveLocation(c *gin.Context) {
	var reserveLocation entity.ReserveLocation
	id := c.Param("id")
	if err := entity.DB().Preload("ReserveStatus").Preload("Location").Preload("Activity").Preload("Request").Raw("SELECT * FROM reserve_locations WHERE id = ?", id).Find(&reserveLocation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": reserveLocation})
}

// GET /reserve_locations
func ListReserveLocations(c *gin.Context) {
	var reserveLocations []entity.ReserveLocation
	if err := entity.DB().Preload("ReserveStatus").Preload("Location").Preload("Activity").Preload("Request").Raw("SELECT * FROM reserve_locations").Find(&reserveLocations).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reserveLocations})
}

// GET /reserve_location/location/:id
func ListReserveLocationsFromLocation(c *gin.Context) {
	var reserveLocations []entity.ReserveLocation
	id := c.Param("id")
	if err := entity.DB().Preload("ReserveStatus").Preload("Location").Preload("Activity").Preload("Request").Raw("SELECT * FROM reserve_locations WHERE location_id = ?", id).Find(&reserveLocations).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reserveLocations})
}
// GET /reserve_location/reserve_status/:id
func ListReserveLocationsFromReserveStatus(c *gin.Context) {
	var reserveLocations []entity.ReserveLocation
	id := c.Param("id")
	if err := entity.DB().Preload("ReserveStatus").Preload("Location").Preload("Activity").Preload("Request").Raw("SELECT * FROM reserve_locations WHERE reserve_status_id != ?", id).Find(&reserveLocations).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": reserveLocations})
}
// DELETE /reserve_locations/:id
func DeleteReserveLocation(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM reserve_locations WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "reserveLocation not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /reserve_locations
func UpdateReserveLocation(c *gin.Context) {
	var reserveLocation entity.ReserveLocation
	if err := c.ShouldBindJSON(&reserveLocation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", reserveLocation.ID).First(&reserveLocation); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "reserveLocation not found"})
		return
	}

	if err := entity.DB().Save(&reserveLocation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reserveLocation})
}
