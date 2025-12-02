package handler

import (
	"database/sql"
	"errors"
	helpers "github.com/ZaharBorisenko/Management-System-Car/internal/handler/helpers/ID"
	"github.com/ZaharBorisenko/Management-System-Car/internal/models"
	"log/slog"
	"net/http"

	libJSON "github.com/ZaharBorisenko/Management-System-Car/internal/handler/helpers/JSON"
	"github.com/ZaharBorisenko/Management-System-Car/internal/service"
)

type CarHandler struct {
	service service.CarServiceInterface
	logger  *slog.Logger
}

func NewCarHandler(service service.CarServiceInterface, logger *slog.Logger) *CarHandler {
	if logger == nil {
		logger = slog.Default()
	}
	return &CarHandler{
		service: service,
		logger:  logger,
	}
}

func (h *CarHandler) GetCarById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")

	if err := helpers.CheckID(id); err != nil {
		libJSON.WriteError(w, http.StatusBadRequest, err.Error())
	}

	car, err := h.service.GetCarById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			libJSON.WriteError(w, http.StatusNotFound, "Car not found")
		} else {
			libJSON.WriteError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	libJSON.WriteJSON(w, http.StatusOK, &car)
}

func (h *CarHandler) GetCarByBrand(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	brand := r.PathValue("brand")

	cars, err := h.service.GetCarByBrand(ctx, brand)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			libJSON.WriteError(w, http.StatusNotFound, "Brand not found")
		} else {
			libJSON.WriteError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	libJSON.WriteJSON(w, http.StatusOK, &cars)
}

func (h *CarHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	carReq := models.CarRequestDTO{}
	if err := libJSON.ReadJSON(r, &carReq); err != nil {
		libJSON.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	createdCar, err := h.service.CreateCar(ctx, &carReq)
	if err != nil {
		libJSON.WriteError(w, http.StatusBadRequest, err.Error())
	}

	libJSON.WriteJSON(w, http.StatusCreated, createdCar)
}
