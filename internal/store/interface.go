package store

import (
	"context"
	"github.com/ZaharBorisenko/Management-System-Car/internal/models/dto"
	"github.com/ZaharBorisenko/Management-System-Car/internal/models/entity"
)

type CarStoreInterface interface {
	GetAllCar(ctx context.Context) ([]dto.CarResponse, error)
	GetCarById(ctx context.Context, id string) (dto.CarResponse, error)
	GetCarByVinCode(ctx context.Context, vinCode string) (dto.CarResponse, error)
	GetCarByBrand(ctx context.Context, brand string) ([]dto.CarResponse, error)

	CreateCar(ctx context.Context, req *dto.CarCreateRequest) (dto.CarResponse, error)
	UpdateCar(ctx context.Context, req *dto.CarUpdateRequest, id string) error
	DeleteCar(ctx context.Context, id string) error
}
type EngineStoreInterface interface {
	GetAllEngine(ctx context.Context) ([]entity.Engine, error)
	GetEngineById(ctx context.Context, id string) (entity.Engine, error)
	CreateEngine(ctx context.Context, e entity.Engine) (entity.Engine, error)
	UpdateEngine(ctx context.Context, req *dto.EngineUpdateRequest, id string) error
	DeleteEngine(ctx context.Context, id string) error
}
