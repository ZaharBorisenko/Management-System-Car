package helpers

import (
	"errors"
	"github.com/google/uuid"
)

func CheckID(id string) error {
	if id == "" {
		return errors.New("ID is empty")
	}

	if _, err := uuid.Parse(id); err != nil {
		return errors.New("invalid ID format")
	}

	return nil
}
