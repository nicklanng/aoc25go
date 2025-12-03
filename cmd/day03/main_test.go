package main

import "testing"

func TestScanRow(t *testing.T) {
	tests := []struct {
		input      string
		batteryLen int
		expected   int
	}{
		{"987654321111111", 2, 98},
		{"811111111111119", 2, 89},
		{"234234234234278", 2, 78},
		{"818181911112111", 2, 92},
		{"987654321111111", 12, 987654321111},
		{"811111111111119", 12, 811111111119},
		{"234234234234278", 12, 434234234278},
		{"818181911112111", 12, 888911112111},
	}

	for _, test := range tests {
		result := scanRow(test.input, test.batteryLen)
		if result != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, result)
		}
	}
}
