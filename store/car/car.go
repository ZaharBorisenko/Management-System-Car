package car

import (
	"context"
	"database/sql"

	"github.com/ZaharBorisenko/Management-System-Car/models"
)

type CarStore struct {
	db *sql.DB
}

func NewCarStore(db *sql.DB) *CarStore {
	return &CarStore{db: db}
}

// === GET ===
func (s *CarStore) GetAllCar(ctx context.Context) ([]models.Car, error) {
	return []models.Car{}, nil
}
func (s *CarStore) GetCarById(ctx context.Context, id string) (models.Car, error) {
	return models.Car{}, nil
}
func (s *CarStore) GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error) {
	return []models.Car{}, nil
}
func (s *CarStore) GetCarByBodyType(ctx context.Context, bodyType string) ([]models.Car, error) {
	return []models.Car{}, nil
}
func (s *CarStore) GetCarByColor(ctx context.Context, color string) ([]models.Car, error) {
	return []models.Car{}, nil
}
func (s *CarStore) GetCarByFuelType(ctx context.Context, fuelType string) ([]models.Car, error) {
	return []models.Car{}, nil
}
func (s *CarStore) GetCarByVinCode(ctx context.Context, vinCode string) (models.Car, error) {
	return models.Car{}, nil
}

// === CREATE ===
func (s *CarStore) CreateCar(ctx context.Context, req *models.CarRequestDTO) (models.Car, error) {
	return models.Car{}, nil
}

// === UPDATE ===
func (s *CarStore) UpdateCar(ctx context.Context, req *models.CarRequestDTO, id string) (models.Car, error) {
	return models.Car{}, nil
}

// === DELETE ===
func (s *CarStore) DeleteCar(ctx context.Context, id string) (models.Car, error) {
	return models.Car{}, nil
}
