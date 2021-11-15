package types

import (
	"errors"
	"github.com/pborman/uuid"
	"time"
)

type SystemType string

func (t *SystemType) Validate() error {
	return nil
}

type SystemID uuid.UUID

func (i SystemID) Validate() error {
	return nil
}

type SystemName string

func (n SystemName) Validate() error {
	if len(n) == 0 {
		return errors.New("name cannot be empty")
	}
	return nil
}

type SystemCleaningIntervalInDays int16

func (c SystemCleaningIntervalInDays) Validate() error {
	if c < 1 {
		return errors.New("cleaning interval needs to be >= 1")
	}
	return nil
}

type SystemLastCleaned time.Time

func (l SystemLastCleaned) Validate() error {
	return nil
}

const (
	Unknown     SystemType = "unknown"
	Techniplast SystemType = "Techniplast"
	Glass       SystemType = "Glass"
)
