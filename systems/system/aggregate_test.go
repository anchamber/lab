package system_test

import (
	"github.com/anchamber/lab/systems/system"
	"github.com/anchamber/lab/types"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name             string
		systemName       types.SystemName
		systemType       types.SystemType
		cleaningInterval types.SystemCleaningIntervalInDays
		lastCleaned      types.SystemLastCleaned
	}{
		{
			name:             "create valid system",
			systemName:       types.SystemName(""),
			systemType:       types.Techniplast,
			cleaningInterval: types.SystemCleaningIntervalInDays(25),
			lastCleaned:      types.SystemLastCleaned(time.Now()),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			s := system.New(tc.systemName, tc.systemType, tc.cleaningInterval, tc.lastCleaned)
			if len(s.Events()) != 1 {
				t.Errorf("New should have created exactly one event, created %d event(s)", len(s.Events()))
			}

			e := s.Events()[0]
			_, ok := e.(system.SystemCreated)
			if !ok {
				t.Errorf("New should have created a SystemCreated event, created a %T event", e)
			}
		})
	}
}
