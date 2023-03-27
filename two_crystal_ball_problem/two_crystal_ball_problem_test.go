package twocrystalballproblem

import (
	"testing"
)

func TestTwoCrystalBall(t *testing.T) {
	testCases := []struct {
		name     string
		breaks   []bool
		expected int
	}{
		{
			name:     "Single floor - ball does not break",
			breaks:   []bool{false},
			expected: -1,
		},
		{
			name:     "Single floor - ball breaks",
			breaks:   []bool{true},
			expected: 0,
		},
		{
			name:     "Two floors - ball breaks on first floor",
			breaks:   []bool{true, true},
			expected: 0,
		},
		{
			name:     "Two floors - ball breaks on second floor",
			breaks:   []bool{false, true},
			expected: 1,
		},
		{
			name:     "Multiple floors - ball breaks on last floor",
			breaks:   []bool{false, false, false, false, true},
			expected: 4,
		},
		{
			name:     "Multiple floors - ball breaks on middle floor",
			breaks:   []bool{false, false, true, true, true},
			expected: 2,
		},
		{
			name:     "All floors - ball does not break",
			breaks:   []bool{false, false, false, false, false},
			expected: -1,
		},
		{
			name:     "All floors - ball breaks",
			breaks:   []bool{true, true, true, true, true},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := TwoCrystalBall(tc.breaks)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}