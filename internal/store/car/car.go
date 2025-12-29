package car

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ZaharBorisenko/Management-System-Car/internal/models/dto"
	"github.com/ZaharBorisenko/Management-System-Car/internal/models/mapper"
	"github.com/ZaharBorisenko/Management-System-Car/internal/myErr"
	"github.com/ZaharBorisenko/Management-System-Car/internal/store"
	"github.com/lib/pq"
	"log/slog"
	"reflect"
	"strings"
	"time"

	"github.com/ZaharBorisenko/Management-System-Car/internal/store/helper"
	"github.com/google/uuid"
)

type Store struct {
	db          *sql.DB
	logger      *slog.Logger
	engineStore store.EngineStoreInterface
}

func NewCarStore(
	db *sql.DB,
	engineStore store.EngineStoreInterface,
	logger *slog.Logger,
) *Store {
	return &Store{
		db:          db,
		engineStore: engineStore,
		logger:      logger,
	}
}

const CAR_SELECT = `
SELECT
    c.id, c.description, c.year, c.model, c.fuel_type,
    c.price, c.vin, c.mileage, c.transmission, c.color, c.body_type,
    c.created_at, c.updated_at,

    b.id, b.name, b.created_at, b.updated_at,

    e.id, e.horse_power, e.engine_type, e.created_at, e.updated_at
FROM cars c
JOIN brands b ON b.id = c.brand_id
LEFT JOIN engines e ON e.id = c.engine_id
`

func pgErrorCode(err error) pq.ErrorCode {
	var pgErr *pq.Error
	if errors.As(err, &pgErr) {
		return pgErr.Code
	}
	return ""
}

func (s *Store) GetAllCar(ctx context.Context) ([]dto.CarResponse, error) {
	query := CAR_SELECT
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cars := make([]dto.CarResponse, 0)

	for rows.Next() {
		car, err := helper.ScanCar(rows)
		if err != nil {
			return nil, err
		}
		cars = append(cars, mapper.ToCarResponse(car))
	}

	if len(cars) == 0 {
		return nil, myErr.ErrNotFound
	}

	return cars, nil
}

func (s *Store) GetCarById(ctx context.Context, id string) (dto.CarResponse, error) {
	query := CAR_SELECT + " WHERE c.id = $1"
	row := s.db.QueryRowContext(ctx, query, id)

	car, err := helper.ScanCar(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.CarResponse{}, myErr.ErrNotFound
		}
		return dto.CarResponse{}, err
	}

	return mapper.ToCarResponse(car), nil
}

func (s *Store) GetCarByBrand(ctx context.Context, brand string) ([]dto.CarResponse, error) {
	query := CAR_SELECT + " WHERE b.id = $1"

	rows, err := s.db.QueryContext(ctx, query, brand)
	if err != nil {
		return nil, fmt.Errorf("get cars by brand query: %w", err)
	}
	defer rows.Close()

	cars := make([]dto.CarResponse, 0)

	for rows.Next() {
		car, err := helper.ScanCar(rows)
		if err != nil {
			return nil, fmt.Errorf("scan car:  %w", err)
		}
		cars = append(cars, mapper.ToCarResponse(car))
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error:  %w", err)
	}

	if len(cars) == 0 {
		return nil, myErr.ErrNotFound
	}

	return cars, nil
}

func (s *Store) GetCarByVinCode(ctx context.Context, vinCode string) (dto.CarResponse, error) {
	query := CAR_SELECT + "WHERE c.vin = $1"

	row := s.db.QueryRowContext(ctx, query, vinCode)
	car, err := helper.ScanCar(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.CarResponse{}, myErr.ErrNotFound
		}
		return dto.CarResponse{}, err
	}

	return mapper.ToCarResponse(car), nil
}

func (s *Store) CreateCar(ctx context.Context, req *dto.CarCreateRequest) (dto.CarResponse, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return dto.CarResponse{}, err
	}
	defer tx.Rollback()

	// check engine
	var engineExists bool
	err = tx.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM engines WHERE id = $1)", req.EngineID).Scan(&engineExists)

	if err != nil {
		return dto.CarResponse{}, fmt.Errorf("check engine existence: %w", err)
	}
	if !engineExists {
		return dto.CarResponse{}, myErr.ErrEngineNotFound
	}

	// check brand
	var brandExists bool
	err = tx.QueryRowContext(
		ctx,
		"SELECT EXISTS(SELECT 1 FROM brands WHERE id = $1)",
		req.BrandID,
	).Scan(&brandExists)
	if err != nil {
		return dto.CarResponse{}, err
	}
	if !brandExists {
		return dto.CarResponse{}, myErr.ErrNotFound
	}

	carID := uuid.New()
	now := time.Now()

	query := `
INSERT INTO cars (
	id, description, year, model, fuel_type,
	engine_id, brand_id, price, vin, mileage,
	transmission, color, body_type, created_at, updated_at
) VALUES (
	$1,$2,$3,$4,$5,
	$6,$7,$8,$9,$10,
	$11,$12,$13,$14,$15
)
`

	_, err = tx.ExecContext(ctx, query,
		carID,
		req.Description,
		req.Year,
		req.Model,
		req.FuelType,
		req.EngineID,
		req.BrandID,
		req.Price,
		req.VIN,
		req.Mileage,
		req.Transmission,
		req.Color,
		req.BodyType,
		now,
		now,
	)
	if err != nil {
		switch pgErrorCode(err) {
		case "23505":
			return dto.CarResponse{}, myErr.ErrDuplicateVIN
		case "23503":
			return dto.CarResponse{}, myErr.ErrConflict
		default:
			return dto.CarResponse{}, err
		}
	}

	if err := tx.Commit(); err != nil {
		return dto.CarResponse{}, err
	}

	return s.GetCarById(ctx, carID.String())
}

func (s *Store) UpdateCar(ctx context.Context, req *dto.CarUpdateRequest, id string) error {
	type field struct {
		column string
		value  any
		valid  bool
	}

	fields := []field{
		{"description", req.Description, req.Description != nil},
		{"year", req.Year, req.Year != nil},
		{"model", req.Model, req.Model != nil},
		{"fuel_type", req.FuelType, req.FuelType != nil},
		{"engine_id", req.EngineID, req.EngineID != nil},
		{"brand_id", req.BrandID, req.BrandID != nil},
		{"price", req.Price, req.Price != nil},
		{"vin", req.VIN, req.VIN != nil},
		{"mileage", req.Mileage, req.Mileage != nil},
		{"transmission", req.Transmission, req.Transmission != nil},
		{"color", req.Color, req.Color != nil},
		{"body_type", req.BodyType, req.BodyType != nil},
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
UPDATE cars
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

func (s *Store) DeleteCar(ctx context.Context, id string) error {
	query := `DELETE FROM cars WHERE id = $1`

	result, err := s.db.ExecContext(ctx, query, id)
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
