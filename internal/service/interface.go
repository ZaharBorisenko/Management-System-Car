package service

import (
	"context"

	"github.com/ZaharBorisenko/Management-System-Car/internal/models"
)

type CarServiceInterface interface {
	GetAllCar(ctx context.Context) (*[]models.Car, error)
	GetCarById(ctx context.Context, id string) (*models.Car, error)
	GetCarByVinCode(ctx context.Context, vinCode string) (*models.Car, error)
	GetCarByBrand(ctx context.Context, brand string) (*[]models.Car, error)
	CreateCar(ctx context.Context, req *models.CarRequestDTO) (*models.Car, error)
}

type EngineServiceInterface interface {
	GetEngineById(ctx context.Context, id string) (*models.Engine, error)
	CreateEngine(ctx context.Context, req *models.EngineRequestDTO) (*models.Engine, error)
}
