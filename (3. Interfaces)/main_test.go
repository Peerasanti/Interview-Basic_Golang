package main

import "testing"

func TestArea(t *testing.T) {
	tests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{
			name:     "Circle with radius 5",
			shape:    Circle{Radius: 5},
			expected: 78.5,
		},
		{
			name:     "Circle with radius 1",
			shape:    Circle{Radius: 1},
			expected: 3.14,
		},
		{
			name:     "Rectangle with width 5 and height 3",
			shape:    Rectangle{Width: 5, Height: 3},
			expected: 15,
		},
		{
			name:     "Rectangle with width 3 and height 3",
			shape:    Rectangle{Width: 3, Height: 3},
			expected: 9,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.shape.Area()
			if result != test.expected {
				t.Errorf("Expected %f, but got %f", test.expected, result)
			}
		})
	}
}
