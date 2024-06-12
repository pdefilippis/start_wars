package handlers

import "github.com/gin-gonic/gin"

func CreateVehicleEndpoints(router *gin.Engine, vehicleManager VehicleManager){
	endpoints := router.Group("/api/v1")

	endpoints.GET("/vehicle/totals/:name", vehicleManager.GetTotal)
	endpoints.PUT("/vehicle/totals", vehicleManager.SetTotal)
	endpoints.POST("/vehicle/increment/:name", vehicleManager.Increment)
	endpoints.POST("/vehicle/decrement/:name", vehicleManager.Decrement)
}