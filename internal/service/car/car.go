package car

import (
	"context"
	"log/slog"

	"github.com/ZaharBorisenko/Management-System-Car/internal/models"
	"github.com/ZaharBorisenko/Management-System-Car/internal/store"
	"github.com/go-playground/validator/v10"
)

type Service struct {
	store     store.CarStoreInterface
	validator *validator.Validate
	logger    *slog.Logger
}

func NewCarService(store store.CarStoreInterface, logger *slog.Logger) *Service {
	return &Service{
		store:     store,
		validator: validator.New(),
		logger:    logger,
	}
}

func (c *Service) GetCarById(ctx context.Context, id string) (*models.Car, error) {
	car, err := c.store.GetCarById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (c *Service) GetCarByBrand(ctx context.Context, brand string) (*[]models.Car, error) {
	cars, err := c.store.GetCarByBrand(ctx, brand)
	if err != nil {
		return nil, err
	}
	return &cars, nil
}

func (c *Service) CreateCar(ctx context.Context, req *models.CarRequestDTO) (*models.Car, error) {
	if err := c.validator.Struct(req); err != nil {
		return nil, err
	}

	createdCar, err := c.store.CreateCar(ctx, req)
	if err != nil {
		return nil, err
	}

	return &createdCar, nil
}
