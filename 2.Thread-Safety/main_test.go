package main

import "testing"

func TestIncreaseNumber(t *testing.T) {
	tests := []struct {
		name        string
		numRoutines int
		expected    int
	}{
		{
			name:        "1000 routines",
			numRoutines: 1000,
			expected:    1000,
		},
		{
			name:        "100 routines",
			numRoutines: 100,
			expected:    100,
		},
		{
			name:        "single routine",
			numRoutines: 1,
			expected:    1,
		},
		{
			name:        "zero routines",
			numRoutines: 0,
			expected:    0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sc := SafeCounter{}

			result := IncreaseNumber(&sc, test.numRoutines)

			if result != test.expected {
				t.Errorf("Expected %d , but got %d", test.expected, result)
			}
		})
	}
}
