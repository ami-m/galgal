package controllers

import (
	"dropit/app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BookDelivery Creates new Delivery object in DB
func BookDelivery(c *gin.Context) {
	delivery := models.Delivery{}
	if err := c.BindJSON(&delivery); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := delivery.SetRecord(delivery.TimeSlotID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Delete Deletes delviry (Soft delete) and changes status to 'CANCELED'
func Delete(c *gin.Context) {
	delivery := models.Delivery{}
	deliveryId, _ := strconv.Atoi(c.Param("deliveryId"))
	err := delivery.DeleteRecord(int64(deliveryId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true})
}

// Daily Returns all Daily Deliveries
func Daily(c *gin.Context) {
	delivery := models.Delivery{}
	deliveries, err := delivery.GetRecords(models.DAILY)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deliveries)
}

// Returns all Weekly deliveries
func Weekly(c *gin.Context) {
	delivery := models.Delivery{}
	deliveries, err := delivery.GetRecords(models.WEEKLY)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deliveries)
}
