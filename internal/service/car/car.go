package car

import (
	"context"
	"log/slog"

	"github.com/ZaharBorisenko/Management-System-Car/internal/models/dto"
	"github.com/ZaharBorisenko/Management-System-Car/internal/store"
	"github.com/go-playground/validator/v10"
)

type Service struct {
	store     store.CarStoreInterface
	validator *validator.Validate
	logger    *slog.Logger
}

func NewCarService(store store.CarStoreInterface, logger *slog.Logger) *Service {
	return &Service{
		store:     store,
		validator: validator.New(),
		logger:    logger,
	}
}

func (s *Service) GetAllCar(ctx context.Context) ([]dto.CarResponse, error) {
	return s.store.GetAllCar(ctx)
}

func (s *Service) GetCarById(ctx context.Context, id string) (dto.CarResponse, error) {
	return s.store.GetCarById(ctx, id)
}

func (s *Service) GetCarByVinCode(ctx context.Context, vin string) (dto.CarResponse, error) {
	return s.store.GetCarByVinCode(ctx, vin)
}

func (s *Service) GetCarByBrand(ctx context.Context, brand string) ([]dto.CarResponse, error) {
	return s.store.GetCarByBrand(ctx, brand)
}

func (s *Service) CreateCar(ctx context.Context, req *dto.CarCreateRequest) (dto.CarResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return dto.CarResponse{}, err
	}
	return s.store.CreateCar(ctx, req)
}

func (s *Service) UpdateCar(ctx context.Context, req *dto.CarUpdateRequest, id string) error {
	if err := s.validator.Struct(req); err != nil {
		return err
	}
	return s.store.UpdateCar(ctx, req, id)
}

func (s *Service) DeleteCar(ctx context.Context, id string) error {
	return s.store.DeleteCar(ctx, id)
}
