package routes

import (
	"dropit/app/http/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {
	v1 := app.Group("/v1")
	{
		v1.POST("/resolve-address", controllers.ResolveAddress)
		v1.POST("/timeslots", controllers.GetAllAvailableTimeSlots)
		v1.POST("/deliveries", controllers.BookDelivery)
		v1.DELETE("/deliveries/:deliveryId", controllers.Delete)
		v1.GET("/deliveries/daily", controllers.Daily)
		v1.GET("/deliveries/weekly", controllers.Weekly)
	}
}
