package engine

import (
	"context"

	"github.com/ZaharBorisenko/Management-System-Car/models"
	"github.com/ZaharBorisenko/Management-System-Car/store"
)

type EngineService struct {
	store store.EngineStoreInterface
}

func NewEngineService(store store.EngineStoreInterface) *EngineService {
	return &EngineService{
		store: store,
	}
}

func (e EngineService) GetEngineById(ctx context.Context, id string) (*models.Engine, error) {
	engine, err := e.store.GetEngineById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &engine, nil
}

func (e EngineService) CreateEngine(ctx context.Context, req *models.EngineRequestDTO) (*models.Engine, error) {
	createdEngine, err := e.store.CreateEngine(ctx, req)
	if err != nil {
		return nil, err
	}
	return &createdEngine, nil
}
