package system

import "github.com/anchamber/lab/types"

// https://victoramartinez.com/posts/event-sourcing-in-go/

// Event is an interface to mark system.
// It contains an un-exported method to prevent outside structs to count as Events for this package.
type Event interface {
	isEvent()
}

func (e SystemCreated) isEvent()               {}
func (e SystemMoved) isEvent()                 {}
func (e SystemResponsibilityChanged) isEvent() {}
func (e SystemCleaned) isEvent()               {}
func (e SystemRemoved) isEvent()               {}

type SystemCreated struct {
	ID                     types.SystemID                     `json:"id"`
	Name                   types.SystemName                   `json:"name"`
	SystemType             types.SystemType                   `json:"systemType"`
	CleaningIntervalInDays types.SystemCleaningIntervalInDays `json:"cleaning_interval"`
	LastCleaned            types.SystemLastCleaned            `json:"last_cleaned"`
	Responsible            string                             `json:"responsible"`
	Location               string                             `json:"location"`
}

type SystemMoved struct {
	ID          types.SystemID `json:"id"`
	NewLocation string         `json:"new_location"`
}

type SystemResponsibilityChanged struct {
	ID             types.SystemID `json:"id"`
	NewResponsible string         `json:"new_responsible"`
}

type SystemCleaned struct {
	ID   types.SystemID          `json:"id"`
	Date types.SystemLastCleaned `json:"date"`
}

type SystemRemoved struct {
	ID types.SystemID `json:"id"`
}
