package helper

import "github.com/ZaharBorisenko/Management-System-Car/models"

type scannable interface {
	Scan(dest ...any) error
}

func ScanCar(row scannable) (models.Car, error) {
	car := models.Car{}

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

		&car.Engine.ID,
		&car.Engine.Displacement,
		&car.Engine.NoOfCyclinders,
		&car.Engine.CarRange,
		&car.Engine.HorsePower,
		&car.Engine.Torque,
		&car.Engine.EngineType,
		&car.Engine.EmissionClass,
		&car.Engine.CreatedAt,
		&car.Engine.UpdatedAt,
	)
	return car, err
}
