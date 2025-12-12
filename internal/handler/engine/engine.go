package handler

import (
	"database/sql"
	"errors"
	helpers "github.com/ZaharBorisenko/Management-System-Car/internal/handler/helpers/ID"
	libJSON "github.com/ZaharBorisenko/Management-System-Car/internal/handler/helpers/JSON"
	"github.com/ZaharBorisenko/Management-System-Car/internal/models"
	"github.com/ZaharBorisenko/Management-System-Car/internal/myErr"
	"log/slog"
	"net/http"

	"github.com/ZaharBorisenko/Management-System-Car/internal/service"
)

type EngineHandler struct {
	service service.EngineServiceInterface
	logger  *slog.Logger
}

func NewEngineHandler(service service.EngineServiceInterface, logger *slog.Logger) *EngineHandler {
	if logger == nil {
		logger = slog.Default()
	}
	return &EngineHandler{
		service: service,
		logger:  logger,
	}
}

func (h *EngineHandler) GetAllCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	engines, err := h.service.GetAllEngine(ctx)
	if err != nil {
		myErr.HandleError(w, err)
		return
	}

	libJSON.WriteJSON(w, http.StatusOK, engines)
}

func (h *EngineHandler) GetEngineById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")

	if err := helpers.CheckID(id); err != nil {
		libJSON.WriteError(w, http.StatusBadRequest, err.Error())
	}

	engine, err := h.service.GetEngineById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			libJSON.WriteError(w, http.StatusNotFound, "Engine not found")
		} else {
			libJSON.WriteError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	libJSON.WriteJSON(w, http.StatusOK, &engine)
}

func (h *EngineHandler) CreateEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	engineReq := models.EngineRequestDTO{}

	if err := libJSON.ReadJSON(r, &engineReq); err != nil {
		libJSON.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	createdEngine, err := h.service.CreateEngine(ctx, &engineReq)
	if err != nil {
		libJSON.WriteError(w, http.StatusBadRequest, err.Error())
	}

	libJSON.WriteJSON(w, http.StatusCreated, createdEngine)
}
