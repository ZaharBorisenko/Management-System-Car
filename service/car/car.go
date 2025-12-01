package car

import (
	"context"

	"github.com/ZaharBorisenko/Management-System-Car/models"
	"github.com/ZaharBorisenko/Management-System-Car/store"
	"github.com/go-playground/validator/v10"
)

type CarService struct {
	store     store.CarStoreInterface
	validator *validator.Validate
}

func NewCarService(store store.CarStoreInterface) *CarService {
	return &CarService{
		store:     store,
		validator: validator.New(),
	}
}

func (c *CarService) GetCarById(ctx context.Context, id string) (*models.Car, error) {
	car, err := c.store.GetCarById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (c *CarService) GetCarByBrand(ctx context.Context, brand string) (*[]models.Car, error) {
	cars, err := c.store.GetCarByBrand(ctx, brand)
	if err != nil {
		return nil, err
	}
	return &cars, nil
}

func (c *CarService) CreateCar(ctx context.Context, req *models.CarRequestDTO) (*models.Car, error) {
	if err := c.validator.Struct(req); err != nil {
		return nil, err
	}

	createdCar, err := c.store.CreateCar(ctx, req)
	if err != nil {
		return nil, err
	}

	return &createdCar, nil
}
