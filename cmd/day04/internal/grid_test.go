package internal

import (
	"slices"
	"testing"
)

func TestGrid_HasRoll(t *testing.T) {
	grid := NewGrid(`.@........
..@.......
..........
..........
..........`)

	if grid.HasRoll(0, 0) {
		t.Errorf("Expected no roll at 0, 0")
	}
	if grid.HasRoll(0, 1) {
		t.Errorf("Expected no roll at 0, 1")
	}
	if !grid.HasRoll(1, 0) {
		t.Errorf("Expected roll at 1, 0")
	}
	if grid.HasRoll(1, 1) {
		t.Errorf("Expected no roll at 1, 1")
	}
	if !grid.HasRoll(2, 1) {
		t.Errorf("Expected roll at 2, 1")
	}
	if grid.HasRoll(2, 2) {
		t.Errorf("Expected no roll at 2, 2")
	}
}

func TestGrid_Neighbors(t *testing.T) {
	grid := NewGrid(`.@........
..@.......
..........
..........
..........`)

	neighbors := slices.Collect(grid.Neighbors(1, 1))

	if len(neighbors) != 8 {
		t.Errorf("Expected 8 neighbors, got %d", len(neighbors))
	}
	if !slices.Contains(neighbors, Point{0, 0}) {
		t.Errorf("Expected neighbor at 0, 0")
	}
	if !slices.Contains(neighbors, Point{0, 1}) {
		t.Errorf("Expected neighbor at 0, 1")
	}
	if !slices.Contains(neighbors, Point{0, 2}) {
		t.Errorf("Expected neighbor at 0, 2")
	}
	if !slices.Contains(neighbors, Point{1, 0}) {
		t.Errorf("Expected neighbor at 1, 0")
	}
	if !slices.Contains(neighbors, Point{1, 2}) {
		t.Errorf("Expected neighbor at 1, 2")
	}
	if !slices.Contains(neighbors, Point{2, 0}) {
		t.Errorf("Expected neighbor at 2, 0")
	}
	if !slices.Contains(neighbors, Point{2, 1}) {
		t.Errorf("Expected neighbor at 2, 1")
	}
	if !slices.Contains(neighbors, Point{2, 2}) {
		t.Errorf("Expected neighbor at 2, 2")
	}
}
