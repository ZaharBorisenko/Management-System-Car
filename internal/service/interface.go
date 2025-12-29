package service

import (
	"context"
	"github.com/ZaharBorisenko/Management-System-Car/internal/models/dto"
)

type CarServiceInterface interface {
	GetAllCar(ctx context.Context) ([]dto.CarResponse, error)
	GetCarById(ctx context.Context, id string) (dto.CarResponse, error)
	GetCarByVinCode(ctx context.Context, vin string) (dto.CarResponse, error)
	GetCarByBrand(ctx context.Context, brand string) ([]dto.CarResponse, error)

	CreateCar(ctx context.Context, req *dto.CarCreateRequest) (dto.CarResponse, error)
	UpdateCar(ctx context.Context, req *dto.CarUpdateRequest, id string) error
	DeleteCar(ctx context.Context, id string) error
}

type EngineServiceInterface interface {
	GetAllEngine(ctx context.Context) ([]dto.EngineResponse, error)
	GetEngineById(ctx context.Context, id string) (dto.EngineResponse, error)
	CreateEngine(ctx context.Context, req *dto.EngineCreateRequest) (dto.EngineResponse, error)
	UpdateEngine(ctx context.Context, req *dto.EngineUpdateRequest, id string) error
	DeleteEngine(ctx context.Context, id string) error
}
