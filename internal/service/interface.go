package service

import (
	"context"

	"github.com/ZaharBorisenko/Management-System-Car/internal/models"
)

type CarServiceInterface interface {
	GetAllCar(ctx context.Context) ([]models.Car, error)
	GetCarById(ctx context.Context, id string) (*models.Car, error)
	GetCarByVinCode(ctx context.Context, vinCode string) (*models.Car, error)
	GetCarByBrand(ctx context.Context, brand string) ([]models.Car, error)
	CreateCar(ctx context.Context, req *models.CarRequestDTO) (*models.Car, error)
	DeleteCar(ctx context.Context, id string) error
	UpdateCar(ctx context.Context, req *models.CarUpdateDTO, id string) error
}

type EngineServiceInterface interface {
	GetAllEngine(ctx context.Context) ([]models.Engine, error)
	GetEngineById(ctx context.Context, id string) (*models.Engine, error)
	CreateEngine(ctx context.Context, req *models.EngineRequestDTO) (*models.Engine, error)
}
