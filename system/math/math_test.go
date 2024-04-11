// math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
	// Test cases
	cases := []struct {
		a, b, expected int
	}{
		{1, 2, 3},       // Case 1: 1 + 2 = 3
		{0, 0, 0},       // Case 2: 0 + 0 = 0
		{-1, 1, 0},      // Case 3: -1 + 1 = 0
		{10, -5, 5},     // Case 4: 10 + (-5) = 5
		{-10, -10, -20}, // Case 5: -10 + (-10) = -20
	}

	// Iterate through test cases
	for _, tc := range cases {
		result := Add(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
		}
	}
}
