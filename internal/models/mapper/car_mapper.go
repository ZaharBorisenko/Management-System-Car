package mapper

import (
	"github.com/ZaharBorisenko/Management-System-Car/internal/models/dto"
	"github.com/ZaharBorisenko/Management-System-Car/internal/models/entity"
)

type CarAggregate struct {
	Car    entity.Car
	Brand  entity.Brand
	Engine entity.Engine
}

func ToCarResponse(a CarAggregate) dto.CarResponse {
	return dto.CarResponse{
		ID:           a.Car.ID,
		Description:  a.Car.Description,
		Year:         a.Car.Year,
		Model:        a.Car.Model,
		FuelType:     a.Car.FuelType,
		Price:        a.Car.Price,
		VIN:          a.Car.VIN,
		Mileage:      a.Car.Mileage,
		Transmission: a.Car.Transmission,
		Color:        a.Car.Color,
		BodyType:     a.Car.BodyType,
		CreatedAt:    a.Car.CreatedAt,
		Brand: dto.BrandShort{
			ID:   a.Brand.ID,
			Name: a.Brand.Name,
		},
		Engine: dto.EngineShort{
			ID:         a.Engine.ID,
			HorsePower: a.Engine.HorsePower,
			EngineType: a.Engine.EngineType,
		},
	}
}
