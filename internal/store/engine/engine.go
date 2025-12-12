package engine

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ZaharBorisenko/Management-System-Car/internal/myErr"
	"log/slog"
	"reflect"
	"strings"
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

func (e *Store) GetAllEngine(ctx context.Context) ([]models.Engine, error) {
	query := ENGINE_SELECT

	rows, err := e.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	engines := make([]models.Engine, 0)

	for rows.Next() {
		engine, err := helper.ScanEngine(rows)
		if err != nil {
			return nil, fmt.Errorf("scan engine:  %w", err)
		}
		engines = append(engines, engine)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	if len(engines) == 0 {
		return nil, myErr.ErrNotFound
	}

	return engines, nil
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

func (e *Store) UpdateEngine(ctx context.Context, req *models.EngineUpdateDTO, id string) error {
	type field struct {
		column string
		value  any
		valid  bool
	}

	fields := []field{
		{"description", req.Description, req.Description != nil},
		{"displacement", req.Displacement, req.Displacement != nil},
		{"no_of_cylinders", req.NoOfCylinders, req.NoOfCylinders != nil},
		{"car_range", req.CarRange, req.CarRange != nil},
		{"horse_power", req.HorsePower, req.HorsePower != nil},
		{"torque", req.Torque, req.Torque != nil},
		{"engine_type", req.EngineType, req.EngineType != nil},
		{"emission_class", req.EmissionClass, req.EmissionClass != nil},
	}

	setParts := []string{}
	args := []any{}
	argID := 1

	for _, f := range fields {
		if f.valid {
			setParts = append(setParts, fmt.Sprintf("%s = $%d", f.column, argID))
			args = append(args, reflect.ValueOf(f.value).Elem().Interface())
			argID++
		}
	}

	if len(setParts) == 0 {
		return nil
	}

	query := fmt.Sprintf(`
        UPDATE engines
        SET %s
        WHERE id = $%d
    `, strings.Join(setParts, ", "), argID)

	args = append(args, id)

	result, err := e.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("db update engine error: %w", err)
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return myErr.ErrNotFound
	}

	return nil
}

func (e *Store) DeleteEngine(ctx context.Context, id string) error {
	query := "DELETE FROM engines WHERE id = $1"
	result, err := e.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting engine: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return myErr.ErrNotFound
	}
	return nil
}
