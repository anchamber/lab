package lines_test

import (
	"github.com/anchamber/lab/libs/types"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		lineName types.LineName
	}{
		{
			name:     "create valid lines",
			lineName: types.LineName(""),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
		})
	}
}
