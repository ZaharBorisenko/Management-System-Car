package myErr

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrInvalidInput   = errors.New("invalid input data")
	ErrConflict       = errors.New("resource already exists")
	ErrDuplicateVIN   = errors.New("car with this VIN already exists")
	ErrEngineNotFound = errors.New("engine not found")
	ErrInternal       = errors.New("internal server error")
)
