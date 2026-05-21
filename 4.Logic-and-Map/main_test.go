package main

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Example case",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Case with negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			target:   -8,
			expected: []int{2, 4},
		},
		{
			name:     "Case with duplicates",
			nums:     []int{1, 3, 3, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "Case with empty slice",
			nums:     []int{},
			target:   0,
			expected: []int{},
		},
		{
			name:     "Case with Target not found",
			nums:     []int{3, 2, 7},
			target:   6,
			expected: []int{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := TwoSum(test.nums, test.target)
			if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", test.expected) {
				t.Errorf("TwoSum() = %v, want %v", result, test.expected)
			}
		})
	}
}
