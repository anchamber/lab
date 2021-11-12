package types

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

func (t *SystemType) Validate() error {
	return nil
}

func (i SystemID) Validate() error {
	return nil
}

func (n SystemName) Validate() error {
	if len(n) == 0 {
		return errors.New("name cannot be empty")
	}
	return nil
}

func (c SystemCleaningIntervalInDays) Validate() error {
	if c < 1 {
		return errors.New("cleaning interval needs to be >= 1")
	}
	return nil
}

func (l SystemLastCleaned) Validate() error {
	return nil
}

type SystemType string
type SystemID uuid.UUID
type SystemName string
type SystemCleaningIntervalInDays int16
type SystemLastCleaned time.Time

const (
	Unknown     SystemType = "unknown"
	Techniplast SystemType = "Techniplast"
	Glass       SystemType = "Glass"
)
