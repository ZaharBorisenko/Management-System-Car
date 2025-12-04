package helper

import (
	"github.com/ZaharBorisenko/Management-System-Car/internal/models"
	"log"
	"strings"
)

type scannable interface {
	Scan(dest ...any) error
}

func ScanCar(row scannable) (models.Car, error) {
	car := models.Car{}
	engine := models.Engine{}

	log.Printf("Starting ScanCar...")

	err := row.Scan(
		&car.ID,
		&car.Description,
		&car.Year,
		&car.Brand,
		&car.Model,
		&car.FuelType,
		&car.Price,
		&car.VIN,
		&car.Mileage,
		&car.Transmission,
		&car.Color,
		&car.BodyType,
		&car.CreatedAt,
		&car.UpdatedAt,

		&engine.ID,
		&engine.Description,
		&engine.Displacement,
		&engine.NoOfCylinders,
		&engine.CarRange,
		&engine.HorsePower,
		&engine.Torque,
		&engine.EngineType,
		&engine.EmissionClass,
		&engine.CreatedAt,
		&engine.UpdatedAt,
	)

	if err != nil {
		log.Printf("ScanCar error: %v, type: %T", err, err)
		// Дополнительная диагностика
		if strings.Contains(err.Error(), "converting") {
			log.Printf("Possible type conversion error")
		}
		return car, err
	}

	car.Engine = engine
	log.Printf("ScanCar successful, car: %+v", car)
	return car, nil
}

func ScanEngine(row scannable) (models.Engine, error) {
	engine := models.Engine{}

	err := row.Scan(
		&engine.ID,
		&engine.Description,
		&engine.Displacement,
		&engine.NoOfCylinders,
		&engine.CarRange,
		&engine.HorsePower,
		&engine.Torque,
		&engine.EngineType,
		&engine.EmissionClass,
		&engine.CreatedAt,
		&engine.UpdatedAt,
	)
	return engine, err
}
