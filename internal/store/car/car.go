package car

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

func NewCarStore(db *sql.DB, logger *slog.Logger) *Store {
	return &Store{db: db, logger: logger}
}

const CAR_SELECT = `
SELECT
    c.id, c.description, c.year, c.brand, c.model, c.fuel_type,
    c.price, c.vin, c.mileage, c.transmission, c.color, c.body_type,
    c.created_at, c.updated_at,
    e.id, e.displacement, e.no_of_cyclinders, e.carRange,
    e.horse_power, e.torque, e.engine_type, e.emission_class,
    e.created_at, e.updated_at
FROM car AS c
LEFT JOIN engine AS e ON c.engine_id = e.id
`

// === GET ===
func (s *Store) GetAllCar(ctx context.Context) ([]models.Car, error) {
	return []models.Car{}, nil
}
func (s *Store) GetCarById(ctx context.Context, id string) (models.Car, error) {
	query := CAR_SELECT + "WHERE c.id = $1"

	row := s.db.QueryRowContext(ctx, query, id)
	car, err := helper.ScanCar(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return car, fmt.Errorf("car not found")
		}
		return car, err
	}

	return car, nil
}

func (s *Store) GetCarByBrand(ctx context.Context, brand string) ([]models.Car, error) {
	query := CAR_SELECT + "WHERE c.brand = $1"

	rows, err := s.db.QueryContext(ctx, query, brand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cars := []models.Car{}
	for rows.Next() {
		car, err := helper.ScanCar(rows)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cars, nil
}
func (s *Store) GetCarByBodyType(ctx context.Context, bodyType string) ([]models.Car, error) {
	return []models.Car{}, nil
}
func (s *Store) GetCarByColor(ctx context.Context, color string) ([]models.Car, error) {
	return []models.Car{}, nil
}
func (s *Store) GetCarByFuelType(ctx context.Context, fuelType string) ([]models.Car, error) {
	return []models.Car{}, nil
}
func (s *Store) GetCarByVinCode(ctx context.Context, vinCode string) (models.Car, error) {
	return models.Car{}, nil
}

// === CREATE ===
func (s *Store) CreateCar(ctx context.Context, req *models.CarRequestDTO) (models.Car, error) {
	createdCar := models.Car{}
	engineID := uuid.UUID{}

	err := s.db.QueryRowContext(ctx, "SELECT id from engine WHERE id = $1", req.Engine.ID).Scan(&engineID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return createdCar, errors.New("id engine does not exists in the engine table")
		}
		return createdCar, nil
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
RETURNING id, description, year, brand, model, fuel_type, price, vin, mileage,
          transmission, color, body_type, created_at, updated_at
`
	err = s.db.QueryRowContext(ctx, query,
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
		return models.Car{}, err
	}

	return createdCar, nil
}

// === UPDATE ===
func (s *Store) UpdateCar(ctx context.Context, req *models.CarRequestDTO, id string) (models.Car, error) {
	return models.Car{}, nil
}

// === DELETE ===
func (s *Store) DeleteCar(ctx context.Context, id string) (models.Car, error) {
	return models.Car{}, nil
}
