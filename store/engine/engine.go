package engine

import (
	"context"
	"database/sql"

	"github.com/ZaharBorisenko/Management-System-Car/models"
)

type EngineStore struct {
	db *sql.DB
}

func NewEngineStore(db *sql.DB) *EngineStore {
	return &EngineStore{db: db}
}

// === GET ===
func (e *EngineStore) GetAllEngine(ctx context.Context) ([]models.Engine, error) {
	return []models.Engine{}, nil
}

func (e *EngineStore) GetEngineById(ctx context.Context, id string) (models.Engine, error) {
	return models.Engine{}, nil
}
func (e *EngineStore) GetEngineByEngineType(ctx context.Context, engineType string) (models.Engine, error) {
	return models.Engine{}, nil
}

// === CREATE ===
func (e *EngineStore) CreateEngine(ctx context.Context, req *models.EngineRequestDTO) (models.Engine, error) {
	return models.Engine{}, nil
}

// === UPDATE ===
func (e *EngineStore) UpdateEngine(ctx context.Context, req *models.EngineRequestDTO, id string) (models.Engine, error) {
	return models.Engine{}, nil
}

// === DELETE ===
func (e *EngineStore) DeleteEngine(ctx context.Context, id string) (models.Engine, error) {
	return models.Engine{}, nil
}
