package engine

import (
	"context"
	"github.com/ZaharBorisenko/Management-System-Car/internal/models/entity"

	"github.com/ZaharBorisenko/Management-System-Car/internal/models/dto"
	"github.com/ZaharBorisenko/Management-System-Car/internal/models/mapper"
	"github.com/ZaharBorisenko/Management-System-Car/internal/store"
	"github.com/go-playground/validator/v10"
)

type Service struct {
	store store.EngineStoreInterface
	v     *validator.Validate
}

func NewEngineService(store store.EngineStoreInterface) *Service {
	return &Service{
		store: store,
		v:     validator.New(),
	}
}

func (s *Service) GetAllEngine(ctx context.Context) ([]dto.EngineResponse, error) {
	es, err := s.store.GetAllEngine(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]dto.EngineResponse, 0, len(es))
	for _, e := range es {
		res = append(res, mapper.ToEngineResponse(e))
	}
	return res, nil
}

func (s *Service) GetEngineById(ctx context.Context, id string) (dto.EngineResponse, error) {
	e, err := s.store.GetEngineById(ctx, id)
	if err != nil {
		return dto.EngineResponse{}, err
	}
	return mapper.ToEngineResponse(e), nil
}

func (s *Service) CreateEngine(ctx context.Context, req *dto.EngineCreateRequest) (dto.EngineResponse, error) {
	if err := s.v.Struct(req); err != nil {
		return dto.EngineResponse{}, err
	}

	e := entity.Engine{
		Description:   req.Description,
		Displacement:  req.Displacement,
		NoOfCylinders: req.NoOfCylinders,
		CarRange:      req.CarRange,
		HorsePower:    req.HorsePower,
		Torque:        req.Torque,
		EngineType:    req.EngineType,
		EmissionClass: req.EmissionClass,
	}

	created, err := s.store.CreateEngine(ctx, e)
	if err != nil {
		return dto.EngineResponse{}, err
	}

	return mapper.ToEngineResponse(created), nil
}

func (s *Service) UpdateEngine(ctx context.Context, req *dto.EngineUpdateRequest, id string) error {
	if err := s.v.Struct(req); err != nil {
		return err
	}
	return s.store.UpdateEngine(ctx, req, id)
}

func (s *Service) DeleteEngine(ctx context.Context, id string) error {
	return s.store.DeleteEngine(ctx, id)
}
