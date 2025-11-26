package models

import "github.com/google/uuid"

type Engine struct {
	ID             uuid.UUID `json:"id"`
	Displacement   int64     `json:"displacement"`   // in cc (e.g. 1998)
	NoOfCyclinders int64     `json:"noOfCyclinders"` // typical: 2, 3, 4, 6, 8, 10, 12, 16
	CarRange       int64     `json:"carRange"`       // km, for EVs or hybrids
	HorsePower     int64     `json:"horse_power"`    // hp
	Torque         int64     `json:"torque"`         // Nm
	EngineType     string    `json:"engine_type"`    // allowed: "i4", "i6", "v6", "v8", "v10", "v12", "electric", "rotary"
	EmissionClass  string    `json:"emission_class"` // allowed: "Euro 3", "Euro 4", "Euro 5", "Euro 6", "Euro 6d"
}

type EngineRequestDTO struct {
	Displacement  int64  `json:"displacement" validate:"required,gte=0,lte=10000"` // 0 для электродвигателя
	NoOfCylinders int64  `json:"noOfCyclinders" validate:"required,oneof=0 2 3 4 6 8 10 12 16"`
	CarRange      int64  `json:"carRange" validate:"gte=0,lte=2000"` // km, для EV/гибридов
	HorsePower    int64  `json:"horse_power" validate:"required,gte=1,lte=2000"`
	Torque        int64  `json:"torque" validate:"required,gte=1,lte=5000"`
	EngineType    string `json:"engine_type" validate:"required,oneof=i4 i6 v6 v8 v10 v12 electric rotary"`
	EmissionClass string `json:"emission_class" validate:"required,oneof='Euro 3' 'Euro 4' 'Euro 5' 'Euro 6' 'Euro 6d'"`
}
