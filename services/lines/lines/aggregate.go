package lines

import (
	"github.com/anchamber/lab/libs/types"
	"github.com/google/uuid"
)

type Line struct {
	id   types.LineID
	name types.LineName

	changes []Event
	version uint64
}

func NewFromEvents(events []Event) *Line {
	Line := &Line{}
	for _, event := range events {
		Line.On(event, false)
	}
	return Line
}

func New(name types.LineName) *Line {
	Line := &Line{}
	Line.raise(LineCreated{
		ID:   types.LineID(uuid.New()),
		Name: name,
	})
	return Line
}

func (s *Line) raise(event Event) {
	s.changes = append(s.changes, event)
	s.On(event, true)
}

func (s *Line) On(event Event, newEvent bool) {
	switch e := event.(type) {
	case *LineCreated:
		s.name = e.Name
	}

	if !newEvent {
		s.version++
	}
}

// ID returns the lines's ID
func (s *Line) ID() types.LineID {
	return s.id
}

// Name returns the lines's name
func (s *Line) Name() types.LineName {
	return s.name
}

// Events returns the lines's events
func (s *Line) Events() []Event {
	return s.changes
}

// Version returns the lines's version
func (s *Line) Version() uint64 {
	return s.version
}
