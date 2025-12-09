package engine

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/ZaharBorisenko/Management-System-Car/internal/models"
	"github.com/ZaharBorisenko/Management-System-Car/internal/store/helper"
	"github.com/google/uuid"
)

type Store struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewEngineStore(db *sql.DB, logger *slog.Logger) *Store {
	return &Store{db: db, logger: logger}
}

const ENGINE_SELECT = `
SELECT
	id, description, displacement, no_of_cylinders, car_range, horse_power,
	torque, engine_type, emission_class, created_at, updated_at
FROM engines
`

// === GET ===
func (e *Store) GetAllEngine(ctx context.Context) ([]models.Engine, error) {
	return []models.Engine{}, nil
}

func (e *Store) GetEngineById(ctx context.Context, id string) (models.Engine, error) {
	query := ENGINE_SELECT + " WHERE id = $1"

	row := e.db.QueryRowContext(ctx, query, id)
	engine, err := helper.ScanEngine(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return engine, fmt.Errorf("engine not found")
		}
		return engine, err
	}

	return engine, nil
}

func (e *Store) GetEngineByEngineType(ctx context.Context, engineType string) (models.Engine, error) {
	return models.Engine{}, nil
}

// === CREATE ===
func (e *Store) CreateEngine(ctx context.Context, req *models.EngineRequestDTO) (models.Engine, error) {
	createdEngine := models.Engine{}

	query := `
INSERT INTO engines (
	id, description, displacement, no_of_cylinders, car_range, horse_power,
	torque, engine_type, emission_class, created_at, updated_at
) VALUES (
	$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11
)
RETURNING id, description, displacement, no_of_cylinders, car_range, horse_power,
          torque, engine_type, emission_class, created_at, updated_at
`

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
		&createdEngine.ID,
		&createdEngine.Description,
		&createdEngine.Displacement,
		&createdEngine.NoOfCylinders,
		&createdEngine.CarRange,
		&createdEngine.HorsePower,
		&createdEngine.Torque,
		&createdEngine.EngineType,
		&createdEngine.EmissionClass,
		&createdEngine.CreatedAt,
		&createdEngine.UpdatedAt,
	)

	if err != nil {
		return models.Engine{}, err
	}

	return createdEngine, nil
}

// === UPDATE ===
func (e *Store) UpdateEngine(ctx context.Context, req *models.EngineRequestDTO, id string) (models.Engine, error) {
	return models.Engine{}, nil
}

// === DELETE ===
func (e *Store) DeleteEngine(ctx context.Context, id string) (models.Engine, error) {
	return models.Engine{}, nil
}
