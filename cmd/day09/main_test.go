package main

import (
	"os"
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

func BenchmarkPuzzle2(b *testing.B) {
	input, err := os.ReadFile("../../input/day9")
	if err != nil {
		b.Fatal(err)
	}
	for b.Loop() {
		_ = puzzle2(input)
	}
}
