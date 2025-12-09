package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle1(t *testing.T) {
	input := []byte(`7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`)
	answer := puzzle1(input)
	assert.Equal(t, 50, answer)
}

func TestPuzzle2(t *testing.T) {
	input := []byte(`7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`)
	answer := puzzle2(input)
	assert.Equal(t, 24, answer)
}
