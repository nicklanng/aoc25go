package main

import (
	"reflect"
	"testing"
)

func TestIsSequenceRepeatedOnce(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		// Invalid IDs (should return true)
		{11, true},
		{22, true},
		{99, true},
		{1010, true},
		{1188511885, true},
		{222222, true},
		{446446, true},
		{38593859, true},
		// Valid IDs (should return false)
		{1, false},
		{12, false},
		{95, false},
		{100, false},
		{998, false},
		{1011, false},
		{1012, false},
		{1188511880, false},
		{1188511881, false},
		{222220, false},
		{222221, false},
		{1698522, false},
		{1698523, false},
		{1698528, false},
		{446443, false},
		{446444, false},
		{446445, false},
		{38593856, false},
		{38593857, false},
		{38593858, false},
	}

	for _, test := range tests {
		result := isSequenceRepeatedOnce(test.input)
		if result != test.expected {
			t.Errorf("Expected %v for %d, got %v", test.expected, test.input, result)
		}
	}
}

func TestIsSequenceRepeatedAtLeastOnce(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		// Invalid IDs (should return true)
		{11, true},
		{22, true},
		{99, true},
		{111, true},
		{999, true},
		{1010, true},
		{1188511885, true},
		{222222, true},
		{446446, true},
		{38593859, true},
		{565656, true},
		{824824824, true},
		{2121212121, true},
		// Valid IDs (should return false)
		{1, false},
		{12, false},
		{95, false},
		{100, false},
		{998, false},
		{1011, false},
		{1012, false},
		{1188511880, false},
		{1188511881, false},
		{222220, false},
		{222221, false},
		{1698522, false},
		{1698523, false},
		{1698528, false},
		{446443, false},
		{446444, false},
		{446445, false},
		{565653, false},
		{565654, false},
		{565655, false},
		{824824821, false},
		{824824822, false},
		{824824823, false},
		{2121212118, false},
		{2121212119, false},
		{2121212120, false},
	}

	for _, test := range tests {
		result := isSequenceRepeatedAtLeastOnce(test.input)
		if result != test.expected {
			t.Errorf("Expected %v for %d, got %v", test.expected, test.input, result)
		}
	}
}

func TestFactor(t *testing.T) {
	tests := []struct {
		input    int
		expected []int
	}{
		{1, []int{1}},
		{2, []int{1, 2}},
		{3, []int{1, 3}},
		{4, []int{1, 2, 4}},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(factor(test.input), test.expected) {
			t.Errorf("Expected %v, got %v", test.expected, factor(test.input))
		}
	}
}
