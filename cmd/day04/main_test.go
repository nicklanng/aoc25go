package main

import "testing"

func TestPuzzle1(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	answer := puzzle1(input)
	if answer != 13 {
		t.Errorf("Expected 13, got %d", answer)
	}
}

func TestPuzzle2(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	answer := puzzle2(input)
	if answer != 43 {
		t.Errorf("Expected 43, got %d", answer)
	}
}
