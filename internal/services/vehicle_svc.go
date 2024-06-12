package services

import (
	"context"
	"hop/start_wars/internal/models"
	"hop/start_wars/internal/repositories"

	"go.uber.org/zap"
)

type VehicleService struct{
	starshipRepo repositories.IVehicleRepository
	vehicleRepo repositories.IVehicleRepository
	vehicleExtendRepo repositories.IVehicleExtendRepository
	log *zap.Logger
}

type IVehicleService interface{
	GetTotal(ctx context.Context, model string)(*int, error)
	SetTotal(ctx context.Context, model string, total int) error
	Increment(ctx context.Context, model string) error
	Decrement(ctx context.Context, model string) error
}

func NewVehicleService(log *zap.Logger, starshipRepo repositories.IVehicleRepository, vehicle repositories.IVehicleRepository,
	vehicleExtend repositories.IVehicleExtendRepository)VehicleService{
	return VehicleService{log: log,starshipRepo: starshipRepo, vehicleRepo: vehicle, vehicleExtendRepo: vehicleExtend}
}

func (s VehicleService) GetTotal(ctx context.Context, model string)(*int, error){
	vehicles, err := s.findVehicle(ctx, model)
	if err != nil {
		return nil, err
	}

	var totalVehicle int
	for _, vehicle := range vehicles{
		if vehicle.Name == model {
			totalVehicle++
		}
	}

	vehicle, err := s.vehicleExtendRepo.GetByModel(ctx, model)
	if err != nil {
		return nil, err
	}

	if vehicle != nil {
		totalVehicle += vehicle.Count
	}

	return &totalVehicle, nil
}

func (s VehicleService) findVehicle(ctx context.Context, model string) ([]models.Starship, error){
	allVehicle := make([]models.Starship, 0)
	starship, err := s.starshipRepo.Get(ctx)
	if err != nil {
		return allVehicle, err
	}

	if starship != nil {
		allVehicle = append(allVehicle, *starship...)
	}

	vehicles, err := s.vehicleRepo.Get(ctx)
	if err != nil {
		return allVehicle, err
	}

	if vehicles != nil {
		allVehicle = append(allVehicle, *vehicles...)
	}

	return allVehicle, nil
}

func (s VehicleService)SetTotal(ctx context.Context, model string, total int) error {
	return s.vehicleExtendRepo.SetTotal(ctx, model, total)
}

func (s VehicleService)Increment(ctx context.Context, model string) error {
	//TODO: Se podria pasar la logica de incremento a este nivel, para obtener el vehiculo e incrementar o no la 
	// cantidad segun corresponda, asi de esa forma poder probarlo.
	return s.vehicleExtendRepo.Increment(ctx, model)
}

func (s VehicleService)Decrement(ctx context.Context, model string) error {
	//TODO: idem a la funcion de incremento.
	return s.vehicleExtendRepo.Decrement(ctx, model)
}
