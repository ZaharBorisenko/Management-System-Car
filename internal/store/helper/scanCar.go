package helper

import (
	"github.com/ZaharBorisenko/Management-System-Car/internal/models/mapper"
)

type scannable interface {
	Scan(dest ...any) error
}

func ScanCar(row scannable) (mapper.CarAggregate, error) {
	var agg mapper.CarAggregate

	err := row.Scan(
		&agg.Car.ID,
		&agg.Car.Description,
		&agg.Car.Year,
		&agg.Car.Model,
		&agg.Car.FuelType,
		&agg.Car.Price,
		&agg.Car.VIN,
		&agg.Car.Mileage,
		&agg.Car.Transmission,
		&agg.Car.Color,
		&agg.Car.BodyType,
		&agg.Car.CreatedAt,
		&agg.Car.UpdatedAt,

		&agg.Brand.ID,
		&agg.Brand.Name,
		&agg.Brand.CreatedAt,
		&agg.Brand.UpdatedAt,

		&agg.Engine.ID,
		&agg.Engine.HorsePower,
		&agg.Engine.EngineType,
		&agg.Engine.CreatedAt,
		&agg.Engine.UpdatedAt,
	)

	return agg, err
}

//func ScanEngine(row scannable) (models.Engine, error) {
//	engine := models.Engine{}
//
//	err := row.Scan(
//		&engine.ID,
//		&engine.Description,
//		&engine.Displacement,
//		&engine.NoOfCylinders,
//		&engine.CarRange,
//		&engine.HorsePower,
//		&engine.Torque,
//		&engine.EngineType,
//		&engine.EmissionClass,
//		&engine.CreatedAt,
//		&engine.UpdatedAt,
//	)
//	return engine, err
//}
