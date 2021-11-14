package types_test

import (
	"github.com/anchamber/lab/libs/types"
	"testing"
)

func TestName_Validate(t *testing.T) {
	tests := []struct {
		name       string
		systemName types.SystemName
		valid      bool
	}{
		{
			name:       "validate valid name",
			systemName: types.SystemName("test"),
			valid:      true,
		},
		{
			name:       "validate empty name",
			systemName: types.SystemName(""),
			valid:      false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := tc.systemName.Validate()
			if err != nil {
				if tc.valid {
					t.Errorf("validation should not have failed, err: %s", err)
				}
			} else {
				if !tc.valid {
					t.Errorf("validation should have failed, name: %s", tc.systemName)
				}
			}
		})
	}
}
