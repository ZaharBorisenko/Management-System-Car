package engine

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ZaharBorisenko/Management-System-Car/internal/models/dto"
	"reflect"
	"strings"
	"time"

	"github.com/ZaharBorisenko/Management-System-Car/internal/models/entity"
	"github.com/ZaharBorisenko/Management-System-Car/internal/myErr"
	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
}

func NewEngineStore(db *sql.DB) *Store {
	return &Store{db: db}
}

const ENGINE_SELECT = `
SELECT
	id, description, displacement, no_of_cylinders, car_range,
	horse_power, torque, engine_type, emission_class,
	created_at, updated_at
FROM engines
`

func (s *Store) scan(row *sql.Row) (entity.Engine, error) {
	var e entity.Engine
	err := row.Scan(
		&e.ID, &e.Description, &e.Displacement, &e.NoOfCylinders,
		&e.CarRange, &e.HorsePower, &e.Torque,
		&e.EngineType, &e.EmissionClass,
		&e.CreatedAt, &e.UpdatedAt,
	)
	return e, err
}

func (s *Store) GetAllEngine(ctx context.Context) ([]entity.Engine, error) {
	rows, err := s.db.QueryContext(ctx, ENGINE_SELECT)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []entity.Engine
	for rows.Next() {
		var e entity.Engine
		if err := rows.Scan(
			&e.ID, &e.Description, &e.Displacement, &e.NoOfCylinders,
			&e.CarRange, &e.HorsePower, &e.Torque,
			&e.EngineType, &e.EmissionClass,
			&e.CreatedAt, &e.UpdatedAt,
		); err != nil {
			return nil, err
		}
		result = append(result, e)
	}

	if len(result) == 0 {
		return nil, myErr.ErrNotFound
	}
	return result, nil
}

func (s *Store) GetEngineById(ctx context.Context, id string) (entity.Engine, error) {
	row := s.db.QueryRowContext(ctx, ENGINE_SELECT+" WHERE id=$1", id)
	e, err := s.scan(row)
	if errors.Is(err, sql.ErrNoRows) {
		return e, myErr.ErrNotFound
	}
	return e, err
}

func (s *Store) CreateEngine(ctx context.Context, e entity.Engine) (entity.Engine, error) {
	e.ID = uuid.New()
	e.CreatedAt = time.Now()
	e.UpdatedAt = e.CreatedAt

	query := `
INSERT INTO engines (
	id, description, displacement, no_of_cylinders, car_range,
	horse_power, torque, engine_type, emission_class, created_at, updated_at
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
`

	_, err := s.db.ExecContext(ctx, query,
		e.ID, e.Description, e.Displacement, e.NoOfCylinders,
		e.CarRange, e.HorsePower, e.Torque,
		e.EngineType, e.EmissionClass,
		e.CreatedAt, e.UpdatedAt,
	)
	return e, err
}

func (s *Store) UpdateEngine(ctx context.Context, req *dto.EngineUpdateRequest, id string) error {
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

	setParts := make([]string, 0)
	args := make([]any, 0)
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

	// updated_at
	setParts = append(setParts, fmt.Sprintf("updated_at = $%d", argID))
	args = append(args, time.Now())
	argID++

	query := fmt.Sprintf(`
UPDATE engines
SET %s
WHERE id = $%d
`, strings.Join(setParts, ", "), argID)

	args = append(args, id)

	result, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return myErr.ErrNotFound
	}

	return nil
}

func (s *Store) DeleteEngine(ctx context.Context, id string) error {
	res, err := s.db.ExecContext(ctx, "DELETE FROM engines WHERE id=$1", id)
	if err != nil {
		return err
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return myErr.ErrNotFound
	}
	return nil
}
