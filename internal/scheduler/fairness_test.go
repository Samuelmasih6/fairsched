package scheduler

import (
	"math"
	"testing"
)

func TestJainFairnessIndex(t *testing.T) {
	tests := []struct {
		name     string
		counts   map[string]int
		expected float64
	}{
		{
			name:     "equal distribution",
			counts:   map[string]int{"A": 5, "B": 5},
			expected: 1.0,
		},
		{
			name:     "unequal distribution",
			counts:   map[string]int{"A": 8, "B": 2},
			expected: 100.0 / 136.0,
		},
		{
			name:     "no tenants",
			counts:   map[string]int{},
			expected: 1.0,
		},
		{
			name:     "all zero",
			counts:   map[string]int{"A": 0, "B": 0},
			expected: 1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JainFairnessIndex(tt.counts)

			if math.Abs(got-tt.expected) > 1e-9 {
				t.Fatalf("got %.6f, want %.6f", got, tt.expected)
			}
		})
	}
}
