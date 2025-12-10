package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle1(t *testing.T) {
	input := `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`
	answer := puzzle1([]byte(input))
	assert.Equal(t, 7, answer)
}

func TestPuzzle2(t *testing.T) {
	input := `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`
	answer := puzzle2([]byte(input))
	assert.Equal(t, 33, answer)
}

func TestParseLights(t *testing.T) {
	testCases := []struct {
		input    string
		expected uint16
	}{
		{input: "[.##.]", expected: uint16(0b0110)},
		{input: "[...#.]", expected: uint16(0b01000)},
		{input: "[.###.#]", expected: uint16(0b101110)},
	}
	for _, testCase := range testCases {
		lights := parseLights(testCase.input)
		assert.Equal(t, testCase.expected, lights)
	}
}

func TestParseButtons(t *testing.T) {
	testCases := []struct {
		input    []string
		expected []uint16
	}{
		{
			input:    []string{"(0,2,3,4)", "(2,3)", "(0,4)", "(0,1,2)", "(1,2,3,4)"},
			expected: []uint16{0b11101, 0b1100, 0b10001, 0b111, 0b11110},
		},
	}
	for _, testCase := range testCases {
		buttons := parseButtons(testCase.input)
		assert.Equal(t, testCase.expected, buttons)
	}
}
