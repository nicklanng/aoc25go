package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle1(t *testing.T) {
	// 	input := `3-5
	// 10-14
	// 16-20
	// 12-18

	// 1
	// 5
	// 8
	// 11
	// 17
	// 32`

	input, err := os.ReadFile("../../input/day5")
	if err != nil {
		log.Fatal(err)
	}

	answer := puzzle1(string(input))
	assert.Equal(t, 848, answer)
}
