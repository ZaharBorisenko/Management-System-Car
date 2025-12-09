package handler

import (
	helpers "github.com/ZaharBorisenko/Management-System-Car/internal/handler/helpers/ID"
	"github.com/ZaharBorisenko/Management-System-Car/internal/models"
	"github.com/ZaharBorisenko/Management-System-Car/internal/myErr"
	"github.com/go-playground/validator/v10"
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

func (h *CarHandler) GetAllCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cars, err := h.service.GetAllCar(ctx)
	if err != nil {
		myErr.HandleError(w, err)
		return
	}

	libJSON.WriteJSON(w, http.StatusOK, cars)
}

func (h *CarHandler) GetCarById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")

	if err := helpers.CheckID(id); err != nil {
		libJSON.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	car, err := h.service.GetCarById(ctx, id)
	if err != nil {
		myErr.HandleError(w, err)
		return
	}

	libJSON.WriteJSON(w, http.StatusOK, car)
}

func (h *CarHandler) GetCarByVinCode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vinCode := r.PathValue("vin")

	car, err := h.service.GetCarByVinCode(ctx, vinCode)
	if err != nil {
		myErr.HandleError(w, err)
		return
	}

	libJSON.WriteJSON(w, http.StatusOK, car)
}

func (h *CarHandler) GetCarByBrand(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	brand := r.PathValue("brand")

	cars, err := h.service.GetCarByBrand(ctx, brand)
	if err != nil {
		myErr.HandleError(w, err)
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
		if _, ok := err.(validator.ValidationErrors); ok {
			libJSON.WriteError(w, http.StatusBadRequest, err.Error())
		}

		myErr.HandleError(w, err)
		return
	}

	libJSON.WriteJSON(w, http.StatusCreated, createdCar)
}

func (h *CarHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.PathValue("id")
	if err := helpers.CheckID(id); err != nil {
		libJSON.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	carReq := models.CarUpdateDTO{}
	if err := libJSON.ReadJSON(r, &carReq); err != nil {
		libJSON.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	err := h.service.UpdateCar(ctx, &carReq, id)
	if err != nil {
		myErr.HandleError(w, err)
		return
	}

	libJSON.WriteJSON(w, http.StatusOK, map[string]string{"status": "car updated: " + id})
}

func (h *CarHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.PathValue("id")
	if err := helpers.CheckID(id); err != nil {
		libJSON.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.DeleteCar(ctx, id)
	if err != nil {
		myErr.HandleError(w, err)
		return
	}

	libJSON.WriteJSON(w, http.StatusOK, map[string]string{"status": "car deleted:" + id})
}
