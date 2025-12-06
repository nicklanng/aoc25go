package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle1(t *testing.T) {
	input := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
	answer := puzzle1(input)
	assert.Equal(t, 4277556, answer)
}

func TestPuzzle2(t *testing.T) {
	input := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
	answer := puzzle2([]byte(input))
	assert.Equal(t, 3263827, answer)
}
