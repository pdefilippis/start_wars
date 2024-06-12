package handlers

import (
	"hop/start_wars/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VehicleManager struct{
	log *zap.Logger
	vehicleSvc services.IVehicleService
}

func NewVehicleManager(log *zap.Logger, vehicleSvc services.IVehicleService) VehicleManager{
	return VehicleManager{ log: log, vehicleSvc: vehicleSvc}
}

func(m VehicleManager) GetTotal(ctx *gin.Context){
	total, err := m.vehicleSvc.GetTotal(ctx, ctx.Params.ByName("name"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"total": total} )
}

func(m VehicleManager) SetTotal(ctx *gin.Context){
	type TotalVehicle struct{
		Total int `json:"total"`
		Model string `json:"model"`
	}

	var totalVehicle TotalVehicle

	if err := ctx.ShouldBindJSON(&totalVehicle); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err := m.vehicleSvc.SetTotal(ctx, totalVehicle.Model, totalVehicle.Total)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	
	ctx.JSON(http.StatusOK, nil)
}

func (m VehicleManager) Increment(ctx *gin.Context){
	err := m.vehicleSvc.Increment(ctx, ctx.Param("name"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	
	ctx.JSON(http.StatusOK, nil)
}

func (m VehicleManager) Decrement(ctx *gin.Context){
	err := m.vehicleSvc.Decrement(ctx, ctx.Param("name"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	
	ctx.JSON(http.StatusOK, nil)
}