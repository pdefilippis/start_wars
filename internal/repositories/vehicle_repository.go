package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"hop/start_wars/internal/models"
	"io/ioutil"
	"net/http"
)


type VehicleRepository struct{
	rootService string
	endpoint string
}

type IVehicleRepository interface{
	Get(ctx context.Context)(*[]models.Starship, error)
}

func NewVehicleRepository(rootService string, endpoint string) VehicleRepository{
	return VehicleRepository{rootService: rootService, endpoint: endpoint}
}

func (r VehicleRepository) Get(ctx context.Context)(*[]models.Starship, error) {
	startship, err := r.getStarship(ctx, fmt.Sprintf("%s/%s", r.rootService, r.endpoint))
	return &startship, err
}

func (r VehicleRepository) getStarship(ctx context.Context, url string)([]models.Starship, error){
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result models.StarshipResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	starship := make([]models.Starship, 0)
	starship = append(starship, result.Results...)
	if result.Next != nil {
		data, err := r.getStarship(ctx, *result.Next)
		if err != nil {
			return nil, err
		}

		starship = append(starship, data...)
		
	}

	return starship, nil
}