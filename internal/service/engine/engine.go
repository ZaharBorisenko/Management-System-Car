package engine

import (
	"context"
	"log/slog"

	"github.com/ZaharBorisenko/Management-System-Car/internal/models"
	"github.com/ZaharBorisenko/Management-System-Car/internal/store"
	"github.com/go-playground/validator/v10"
)

type Service struct {
	store     store.EngineStoreInterface
	validator *validator.Validate
	logger    *slog.Logger
}

func NewEngineService(store store.EngineStoreInterface, logger *slog.Logger) *Service {
	return &Service{
		store:     store,
		validator: validator.New(),
		logger:    logger,
	}
}
func (e Service) GetAllEngine(ctx context.Context) ([]models.Engine, error) {
	engines, err := e.store.GetAllEngine(ctx)
	if err != nil {
		return nil, err
	}
	return engines, nil
}

func (e Service) GetEngineById(ctx context.Context, id string) (*models.Engine, error) {
	engine, err := e.store.GetEngineById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &engine, nil
}

func (e Service) CreateEngine(ctx context.Context, req *models.EngineRequestDTO) (*models.Engine, error) {
	if err := e.validator.Struct(req); err != nil {
		return nil, err
	}

	createdEngine, err := e.store.CreateEngine(ctx, req)
	if err != nil {
		return nil, err
	}
	return &createdEngine, nil
}
