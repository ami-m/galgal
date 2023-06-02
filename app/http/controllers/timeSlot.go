package controllers

import (
	"dropit/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllAvailableTimeSlots returns all available time slots
func GetAllAvailableTimeSlots(c *gin.Context) {
	address := models.Address{}
	if err := c.BindJSON(&address); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := models.TimeSlot{}.GetRecords(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
