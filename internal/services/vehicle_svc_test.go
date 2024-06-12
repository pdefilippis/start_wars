package services

import (
	"context"
	"errors"
	"hop/start_wars/internal/models"
	"hop/start_wars/internal/repositories/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetTotal(t *testing.T){
	vehicleData := []models.Starship{
		{Name: "Sand Crawler", Model: "Digger Crawler"},
		{Name: "T-16 skyhopper", Model: "T-16 skyhopper"},
		{Name: "X-34 landspeeder", Model: "X-34 landspeeder"},
		{Name: "TIE/LN starfighter", Model: "Twin Ion Engine/Ln Starfighter"},
		{Name: "Snowspeeder", Model: "t-47 airspeeder"},
		{Name: "TIE bomber", Model: "TIE/sa bomber"},
	}

	starshipData := []models.Starship{
		{Name: "Naboo star skiff", Model: "J-type star skiff"},
		{Name: "Jedi Interceptor", Model: "Eta-2 Actis-class light interceptor"},
		{Name: "arc-170", Model: "Aggressive Reconnaissance-170 starfighte"},
		{Name: "Banking clan frigte", Model: "Munificent-class star frigate"},
		{Name: "Belbullab-22 starfighter", Model: "Belbullab-22 starfighter"},
	}

	testCase := []struct {
		Name string
		Args struct{
			Model string
		}
		VehicleRepo struct{
			Data *[]models.Starship
			Error error
		}
		StarshipData struct{
			Data *[]models.Starship
			Error error
		}
		Expected struct{
			Count int
			Error error
		}
	}{
		{
			Name: "Deberia devolver 1, encontrando solamente 'Sand Crawler'", 
			Args:struct{Model string}{Model: "Sand Crawler"},
			VehicleRepo: struct{Data *[]models.Starship; Error error}{
				Data: &vehicleData,
				Error: nil,
			},
			StarshipData: struct{Data *[]models.Starship; Error error}{
				Data: &starshipData,
				Error: nil,
			},
			Expected: struct{Count int; Error error}{
				Count: 1,
				Error: nil,
			},
		},
		{
			Name: "Deberia devolver 0 al no encontrar ninguno", 
			Args:struct{Model string}{Model: "SandCrawler"},
			VehicleRepo: struct{Data *[]models.Starship; Error error}{
				Data: &vehicleData,
				Error: nil,
			},
			StarshipData: struct{Data *[]models.Starship; Error error}{
				Data: &starshipData,
				Error: nil,
			},
			Expected: struct{Count int; Error error}{
				Count: 0,
				Error: nil,
			},
		},
		{
			Name: "Deberia devolver error al fallar vehicleRepo", 
			Args:struct{Model string}{Model: "SandCrawler"},
			VehicleRepo: struct{Data *[]models.Starship; Error error}{
				Data: nil,
				Error: errors.New("No se puedo acceder al repositorio"),
			},
			StarshipData: struct{Data *[]models.Starship; Error error}{
				Data: &starshipData,
				Error: nil,
			},
			Expected: struct{Count int; Error error}{
				Count: 0,
				Error: errors.New("No se puedo acceder al repositorio"),
			},
		},
	}

	for _, test := range testCase {
		t.Run(test.Name, func(t *testing.T){
			ctrl := gomock.NewController(t)
			vehicleExtendRepo := mock_repositories.NewMockIVehicleExtendRepository(ctrl)
			starshipRepo := mock_repositories.NewMockIVehicleRepository(ctrl)
			vehicleRepo := mock_repositories.NewMockIVehicleRepository(ctrl)

			vehicleRepo.EXPECT().Get(gomock.Any()).Return(&vehicleData, nil)
			starshipRepo.EXPECT().Get(gomock.Any()).Return(starshipData, nil)

			vehicleSvc := NewVehicleService(nil, starshipRepo, vehicleRepo, vehicleExtendRepo)
			total, err := vehicleSvc.GetTotal(context.Background(), test.Args.Model)
			
			assert.Equal(t, test.Expected.Count, total)
			if err != nil || test.Expected.Error != nil {
				assert.Equal(t, test.Expected.Error.Error(), err.Error())
			}
		})
	}

}

//TODO: Falta agregar los test unitarios para los metodos (findVehicle, SetTotal, Increment, Decrement)