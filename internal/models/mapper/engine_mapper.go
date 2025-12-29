package mapper

import (
	"github.com/ZaharBorisenko/Management-System-Car/internal/models/dto"
	"github.com/ZaharBorisenko/Management-System-Car/internal/models/entity"
)

func ToEngineResponse(e entity.Engine) dto.EngineResponse {
	return dto.EngineResponse{
		ID:            e.ID,
		Description:   e.Description,
		Displacement:  e.Displacement,
		NoOfCylinders: e.NoOfCylinders,
		CarRange:      e.CarRange,
		HorsePower:    e.HorsePower,
		Torque:        e.Torque,
		EngineType:    e.EngineType,
		EmissionClass: e.EmissionClass,
		CreatedAt:     e.CreatedAt,
	}
}
