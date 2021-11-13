package lines

import "github.com/anchamber/lab/libs/types"

// https://victoramartinez.com/posts/event-sourcing-in-go/

// Event is an interface to mark lines.
// It contains an un-exported method to prevent outside structs to count as Events for this package.
type Event interface {
	isEvent()
}

func (e LineCreated) isEvent() {}

type LineCreated struct {
	ID   types.LineID   `json:"id"`
	Name types.LineName `json:"name"`
}

type LineActivated struct {
	ID types.LineID `json:"id"`
}

type LineDeactivated struct {
	ID types.LineID `json:"id"`
}
