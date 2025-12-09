package car

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ZaharBorisenko/Management-System-Car/internal/myErr"
	"github.com/lib/pq"
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

func NewCarStore(db *sql.DB, logger *slog.Logger) *Store {
	return &Store{db: db, logger: logger}
}

const CAR_SELECT = `
SELECT
    c.id, c.description, c.year, c.brand, c.model, c.fuel_type,
    c.price, c.vin, c.mileage, c.transmission, c.color, c.body_type,
    c.created_at, c.updated_at,
    e.id, e.description, e.displacement, e.no_of_cylinders, e.car_range,
    e.horse_power, e.torque, e.engine_type, e.emission_class,
    e.created_at, e.updated_at
FROM car AS c
LEFT JOIN engine AS e ON c.engine_id = e.id
`

func pgErrorCode(err error) pq.ErrorCode {
	var pgErr *pq.Error
	if errors.As(err, &pgErr) {
		return pgErr.Code
	}
	return ""
}

func (s *Store) GetAllCar(ctx context.Context) ([]models.Car, error) {
	query := CAR_SELECT

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	cars := make([]models.Car, 0)

	for rows.Next() {
		car, err := helper.ScanCar(rows)
		if err != nil {
			return nil, fmt.Errorf("scan car:  %w", err)
		}
		cars = append(cars, car)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	if len(cars) == 0 {
		return nil, myErr.ErrNotFound
	}

	return cars, nil
}

func (s *Store) GetCarById(ctx context.Context, id string) (*models.Car, error) {
	query := CAR_SELECT + " WHERE c.id = $1"
	row := s.db.QueryRowContext(ctx, query, id)

	car, err := helper.ScanCar(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, myErr.ErrNotFound
		}
		return nil, fmt.Errorf("get car by id query: %w", err)
	}

	return &car, nil
}

func (s *Store) GetCarByBrand(ctx context.Context, brand string) ([]models.Car, error) {
	query := CAR_SELECT + " WHERE c.brand = $1"

	rows, err := s.db.QueryContext(ctx, query, brand)
	if err != nil {
		return nil, fmt.Errorf("get cars by brand query: %w", err)
	}
	defer rows.Close()

	cars := make([]models.Car, 0)

	for rows.Next() {
		car, err := helper.ScanCar(rows)
		if err != nil {
			return nil, fmt.Errorf("scan car:  %w", err)
		}
		cars = append(cars, car)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error:  %w", err)
	}

	if len(cars) == 0 {
		return nil, myErr.ErrNotFound
	}

	return cars, nil
}

func (s *Store) GetCarByVinCode(ctx context.Context, vinCode string) (*models.Car, error) {
	query := CAR_SELECT + "WHERE c.vin = $1"

	row := s.db.QueryRowContext(ctx, query, vinCode)
	car, err := helper.ScanCar(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, myErr.ErrNotFound
		}
		return nil, fmt.Errorf("get car by vinCode query: %w", err)
	}

	return &car, nil
}

func (s *Store) CreateCar(ctx context.Context, req *models.CarRequestDTO) (models.Car, error) {
	//begin transaction
	tx, err := s.db.BeginTx(ctx, nil) // nil = default option (READ COMMITTED)
	if err != nil {
		return models.Car{}, fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()

	//check engine
	var engineExists bool
	err = tx.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM engine WHERE id = $1)", req.Engine.ID).Scan(&engineExists)

	if err != nil {
		return models.Car{}, fmt.Errorf("check engine existence: %w", err)
	}
	if !engineExists {
		return models.Car{}, myErr.ErrEngineNotFound
	}

	newCar := models.Car{
		ID:           uuid.New(),
		Description:  req.Description,
		Year:         req.Year,
		Brand:        req.Brand,
		Model:        req.Model,
		FuelType:     req.FuelType,
		Engine:       req.Engine,
		Price:        req.Price,
		VIN:          req.VIN,
		Mileage:      req.Mileage,
		Transmission: req.Transmission,
		Color:        req.Color,
		BodyType:     req.BodyType,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	query := `
INSERT INTO car (
    id, description, year, brand, model, fuel_type, engine_id,
    price, vin, mileage, transmission, color, body_type, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7,
    $8, $9, $10, $11, $12, $13, $14, $15
)
RETURNING id, description, year, brand, model, fuel_type, engine_id, price, vin, mileage,
          transmission, color, body_type, created_at, updated_at`

	createdCar := models.Car{}

	err = tx.QueryRowContext(ctx, query,
		newCar.ID,
		newCar.Description,
		newCar.Year,
		newCar.Brand,
		newCar.Model,
		newCar.FuelType,
		newCar.Engine.ID,
		newCar.Price,
		newCar.VIN,
		newCar.Mileage,
		newCar.Transmission,
		newCar.Color,
		newCar.BodyType,
		newCar.CreatedAt,
		newCar.UpdatedAt,
	).Scan(
		&createdCar.ID,
		&createdCar.Description,
		&createdCar.Year,
		&createdCar.Brand,
		&createdCar.Model,
		&createdCar.FuelType,
		&createdCar.Engine.ID,
		&createdCar.Price,
		&createdCar.VIN,
		&createdCar.Mileage,
		&createdCar.Transmission,
		&createdCar.Color,
		&createdCar.BodyType,
		&createdCar.CreatedAt,
		&createdCar.UpdatedAt,
	)

	if err != nil {
		switch pgErrorCode(err) {
		case "23505": // Duplicate unique field
			if strings.Contains(err.Error(), "car_vin_key") ||
				strings.Contains(strings.ToLower(err.Error()), "vin") {
				return models.Car{}, myErr.ErrDuplicateVIN
			}
			return models.Car{}, myErr.ErrConflict

		case "23514": // err CHECK (year < 1886, mileage < 0)
			return models.Car{}, myErr.ErrInvalidInput

		case "23503": // foreign key
			return models.Car{}, myErr.ErrEngineNotFound

		default:
			return models.Car{}, fmt.Errorf("insert car: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return models.Car{}, fmt.Errorf("commit transaction: %w", err)
	}

	return createdCar, nil
}

func (s *Store) UpdateCar(ctx context.Context, req *models.CarUpdateDTO, id string) error {
	type field struct {
		column string
		value  any
		valid  bool
	}

	fields := []field{
		{"description", req.Description, req.Description != nil},
		{"year", req.Year, req.Year != nil},
		{"brand", req.Brand, req.Brand != nil},
		{"model", req.Model, req.Model != nil},
		{"fuel_type", req.FuelType, req.FuelType != nil},
		{"engine_id", req.EngineID, req.EngineID != nil},
		{"price", req.Price, req.Price != nil},
		{"vin", req.VIN, req.VIN != nil},
		{"mileage", req.Mileage, req.Mileage != nil},
		{"transmission", req.Transmission, req.Transmission != nil},
		{"color", req.Color, req.Color != nil},
		{"body_type", req.BodyType, req.BodyType != nil},
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
        UPDATE car
        SET %s
        WHERE id = $%d
    `, strings.Join(setParts, ", "), argID)

	args = append(args, id)

	result, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("db update error: %w", err)
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return myErr.ErrNotFound
	}

	return nil
}

func (s *Store) DeleteCar(ctx context.Context, id string) error {
	query := "DELETE FROM car WHERE id = $1"

	result, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
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
