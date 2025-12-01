package engine

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ZaharBorisenko/Management-System-Car/models"
	"github.com/ZaharBorisenko/Management-System-Car/store/helper"
	"github.com/google/uuid"
)

type EngineStore struct {
	db *sql.DB
}

func NewEngineStore(db *sql.DB) *EngineStore {
	return &EngineStore{db: db}
}

const ENGINE_SELECT = `
SELECT
	id, description, displacemen, no_of_cyclinders, carRange, horse_power,
	torque, engine_type, emission_class, created_at, updated_at FROM engine`

// === GET ===
func (e *EngineStore) GetAllEngine(ctx context.Context) ([]models.Engine, error) {
	return []models.Engine{}, nil
}

func (e *EngineStore) GetEngineById(ctx context.Context, id string) (models.Engine, error) {
	query := ENGINE_SELECT + "WHERE id = $1"

	row := e.db.QueryRowContext(ctx, query, id)
	engine, err := helper.ScanEngine(row)

	if err != nil {
		if err == sql.ErrNoRows {
			return engine, fmt.Errorf("engine not found")
		}
		return engine, err
	}

	return engine, nil
}
func (e *EngineStore) GetEngineByEngineType(ctx context.Context, engineType string) (models.Engine, error) {
	return models.Engine{}, nil
}

// === CREATE ===
func (e *EngineStore) CreateEngine(ctx context.Context, req *models.EngineRequestDTO) (models.Engine, error) {
	createdEgine := models.Engine{}

	query := `INSERT INTO engine
	(id, description, displacemen, no_of_cyclinders, carRange, horse_power,
	torque, engine_type, emission_class, created_at, updated_at FROM engine)
	VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	newEngine := models.Engine{
		ID:            uuid.New(),
		Description:   req.Description,
		Displacement:  req.Displacement,
		NoOfCylinders: req.NoOfCylinders,
		CarRange:      req.CarRange,
		HorsePower:    req.HorsePower,
		Torque:        req.Torque,
		EngineType:    req.EngineType,
		EmissionClass: req.EmissionClass,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err := e.db.QueryRowContext(ctx, query,
		newEngine.ID,
		newEngine.Description,
		newEngine.Displacement,
		newEngine.NoOfCylinders,
		newEngine.CarRange,
		newEngine.HorsePower,
		newEngine.Torque,
		newEngine.EngineType,
		newEngine.EmissionClass,
		newEngine.CreatedAt,
		newEngine.UpdatedAt,
	).Scan(
		&createdEgine.ID,
		&createdEgine.Description,
		&createdEgine.Displacement,
		&createdEgine.NoOfCylinders,
		&createdEgine.CarRange,
		&createdEgine.HorsePower,
		&createdEgine.Torque,
		&createdEgine.EngineType,
		&createdEgine.EmissionClass,
		&createdEgine.CreatedAt,
		&createdEgine.UpdatedAt,
	)

	if err != nil {
		return models.Engine{}, err
	}

	return createdEgine, nil

}

// === UPDATE ===
func (e *EngineStore) UpdateEngine(ctx context.Context, req *models.EngineRequestDTO, id string) (models.Engine, error) {
	return models.Engine{}, nil
}

// === DELETE ===
func (e *EngineStore) DeleteEngine(ctx context.Context, id string) (models.Engine, error) {
	return models.Engine{}, nil
}
