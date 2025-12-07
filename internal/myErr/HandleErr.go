package myErr

import (
	"errors"
	libJSON "github.com/ZaharBorisenko/Management-System-Car/internal/handler/helpers/JSON"
	"net/http"
)

func HandleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrNotFound):
		libJSON.WriteError(w, http.StatusNotFound, "resource not found")

	case errors.Is(err, ErrDuplicateVIN), errors.Is(err, ErrConflict):
		libJSON.WriteError(w, http.StatusConflict, "car with this VIN already exists")

	case errors.Is(err, ErrEngineNotFound):
		libJSON.WriteError(w, http.StatusBadRequest, "specified engine does not exist")

	case errors.Is(err, ErrInvalidInput):
		libJSON.WriteError(w, http.StatusBadRequest, "invalid or malformed input data")

	default:
		libJSON.WriteError(w, http.StatusInternalServerError, "internal server error")
	}
}
