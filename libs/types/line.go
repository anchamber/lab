package types

import (
	"errors"
	"github.com/pborman/uuid"
)

type LineName string

func (t LineName) Validate() error {
	if len(t) == 0 {
		return errors.New("name cannot be empty")
	}
	return nil
}

type LineID uuid.UUID

func (i LineID) Validate() error {
	return nil
}
