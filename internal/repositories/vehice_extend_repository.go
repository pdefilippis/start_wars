package repositories

import (
	"context"
	"errors"
	"hop/start_wars/internal/datastore"
	"hop/start_wars/internal/models"

	"gorm.io/gorm"
)

type VehicleExtendRepository struct{
	datastore.IDataStore
}


type IVehicleExtendRepository interface{
	SetTotal(ctx context.Context, model string, total int) error
	Increment(ctx context.Context, model string) error
	Decrement(ctx context.Context, model string) error
	GetByModel(ctx context.Context, model string) (*models.Vehicle, error)
}

func NewVehicleExtendRepository(store datastore.IDataStore)VehicleExtendRepository{
	return VehicleExtendRepository{store}
}

func(r VehicleExtendRepository)SetTotal(ctx context.Context, model string, total int) error {
	vehicle, err := r.GetByModel(ctx, model)
	if err != nil {
		return err
	}

	if vehicle == nil {
		vehicle = &models.Vehicle{
			Model: model,
			Count: total,
		}
		return r.create(ctx, vehicle)
	}
	
	vehicle.Count = total
	return r.save(ctx, vehicle)
}

func(r VehicleExtendRepository)Increment(ctx context.Context, model string) error{
	vehicle, err := r.GetByModel(ctx, model)
	if err != nil {
		return err
	}

	if vehicle == nil {
		vehicle = &models.Vehicle{Model: model, Count: 1}
		return r.create(ctx, vehicle)
	}

	vehicle.Count++
	return r.save(ctx, vehicle)
}

func (r VehicleExtendRepository)Decrement(ctx context.Context, model string) error{
	vehicle, err := r.GetByModel(ctx, model)
	if err != nil {
		return err
	}

	if vehicle == nil {
		vehicle = &models.Vehicle{Model: model, Count: -1}
		return r.create(ctx, vehicle)
	}

	vehicle.Count--
	return r.save(ctx, vehicle)
}

func (r VehicleExtendRepository) GetByModel(ctx context.Context, model string) (*models.Vehicle, error){
	var vehicle models.Vehicle
	err := r.GetDB().Where("model = ?", model).First(&vehicle).Error
	if errors.Is(err, gorm.ErrRecordNotFound){
		return nil, nil
	}

	return &vehicle, err
}

func (r VehicleExtendRepository) create(ctx context.Context, vehicle *models.Vehicle) error {
	return r.GetDB().Create(vehicle).Error
}

func (r VehicleExtendRepository) save(ctx context.Context, vehicle *models.Vehicle) error {
	return r.GetDB().Save(vehicle).Error
}