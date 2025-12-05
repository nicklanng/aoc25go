package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRange(t *testing.T) {
	tests := []struct {
		input    string
		expected Range
	}{
		{"1-10", Range{1, 10}},
		{"234-345", Range{234, 345}},
	}

	for _, test := range tests {
		result, err := ParseRange(test.input)
		assert.NoError(t, err)
		assert.Equal(t, test.expected, result)
	}
}

func TestOverlaps(t *testing.T) {
	tests := []struct {
		input    Range
		other    Range
		expected bool
	}{
		{Range{1, 10}, Range{5, 15}, true},
		{Range{1, 10}, Range{11, 20}, false},
		{Range{5, 15}, Range{1, 10}, true},
	}

	for _, test := range tests {
		result := test.input.Overlaps(test.other)
		assert.Equal(t, test.expected, result)
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		input    Range
		expected int
	}{
		{Range{1, 10}, 10},
		{Range{234, 345}, 112},
	}
	for _, test := range tests {
		result := test.input.Length()
		assert.Equal(t, test.expected, result)
	}
}
