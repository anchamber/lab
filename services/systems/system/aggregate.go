package system

import (
	"github.com/anchamber/lab/libs/types"
	"github.com/pborman/uuid"
)

type System struct {
	id                     types.SystemID
	name                   types.SystemName
	systemType             types.SystemType
	cleaningIntervalInDays types.SystemCleaningIntervalInDays
	lastCleaned            types.SystemLastCleaned
	responsible            string
	location               string
	active                 bool

	changes []Event
	version uint64
}

func NewFromEvents(events []Event) *System {
	system := &System{}
	for _, event := range events {
		system.On(event, false)
	}
	return system
}

func New(name types.SystemName, systemType types.SystemType, intervalInDays types.SystemCleaningIntervalInDays, lastCleaned types.SystemLastCleaned) *System {
	system := &System{}
	system.raise(SystemCreated{
		ID:                     types.SystemID(uuid.New()),
		Name:                   name,
		SystemType:             systemType,
		CleaningIntervalInDays: intervalInDays,
		LastCleaned:            lastCleaned,
		Responsible:            "",
		Location:               "",
	})
	return system
}

func (s *System) Move(newLocation string) error {
	if s.location == newLocation {
		return nil
	}
	s.raise(SystemMoved{
		ID:          s.id,
		NewLocation: newLocation,
	})
	return nil
}

func (s *System) ChangeResponsibility(newResponsible string) error {
	if s.responsible == newResponsible {
		return nil
	}
	s.raise(SystemResponsibilityChanged{
		ID:             s.id,
		NewResponsible: newResponsible,
	})
	return nil
}

func (s *System) Clean(cleaningDate types.SystemLastCleaned) error {
	if s.lastCleaned == cleaningDate {
		return nil
	}
	s.raise(SystemCleaned{
		ID:   s.id,
		Date: cleaningDate,
	})
	return nil
}

func (s *System) raise(event Event) {
	s.changes = append(s.changes, event)
	s.On(event, true)
}

func (s *System) On(event Event, newEvent bool) {
	switch e := event.(type) {
	case *SystemCreated:
		s.name = e.Name
		s.systemType = e.SystemType
		s.cleaningIntervalInDays = e.CleaningIntervalInDays
		s.lastCleaned = e.LastCleaned
		s.responsible = e.Responsible
		s.location = e.Location
	case *SystemMoved:
		s.location = e.NewLocation
	case *SystemCleaned:
		s.lastCleaned = e.Date
	case *SystemResponsibilityChanged:
		s.responsible = e.NewResponsible
	case *SystemRemoved:
		s.active = false
	}

	if !newEvent {
		s.version++
	}
}

// ID returns the lines's ID
func (s *System) ID() types.SystemID {
	return s.id
}

// Name returns the lines's name
func (s *System) Name() types.SystemName {
	return s.name
}

// SystemType returns the lines's type
func (s *System) SystemType() types.SystemType {
	return s.systemType
}

// CleaningIntervalInDays returns the lines's cleaning interval in days
func (s *System) CleaningIntervalInDays() types.SystemCleaningIntervalInDays {
	return s.cleaningIntervalInDays
}

// LastCleaned returns when the lines was cleaned the last time
func (s *System) LastCleaned() types.SystemLastCleaned {
	return s.lastCleaned
}

// Location returns the lines's location id
func (s *System) Location() string {
	return s.location
}

// Responsible returns the lines's responsible person id
func (s *System) Responsible() string {
	return s.responsible
}

// Events returns the lines's events
func (s *System) Events() []Event {
	return s.changes
}

// Version returns the lines's version
func (s *System) Version() uint64 {
	return s.version
}
