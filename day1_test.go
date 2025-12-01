package main

import (
	"strings"
	"testing"
)

func TestPuzzle1(t *testing.T) {
	t.Run("Left 50 steps", func(t *testing.T) {
		input := "L50"
		commands := strings.Split(input, "\n")
		count := puzzle1(commands)
		if count != 1 {
			t.Errorf("Expected 1, got %d", count)
		}
	})

	t.Run("Right 50 steps", func(t *testing.T) {
		input := "R50"
		commands := strings.Split(input, "\n")
		count := puzzle1(commands)
		if count != 1 {
			t.Errorf("Expected 1, got %d", count)
		}
	})

	t.Run("Left 50 steps, then right 50 steps", func(t *testing.T) {
		input := "L50\nR50"
		commands := strings.Split(input, "\n")
		count := puzzle1(commands)
		if count != 1 {
			t.Errorf("Expected 1, got %d", count)
		}
	})

	t.Run("Right 49 steps", func(t *testing.T) {
		input := "R49"
		commands := strings.Split(input, "\n")
		count := puzzle1(commands)
		if count != 0 {
			t.Errorf("Expected 0, got %d", count)
		}
	})

	t.Run("Left 51 steps, Right 1 step", func(t *testing.T) {
		input := "L51\nR1"
		commands := strings.Split(input, "\n")
		count := puzzle1(commands)
		if count != 1 {
			t.Errorf("Expected 1, got %d", count)
		}
	})

	t.Run("Left 250 steps", func(t *testing.T) {
		input := "L250"
		commands := strings.Split(input, "\n")
		count := puzzle1(commands)
		if count != 1 {
			t.Errorf("Expected 1, got %d", count)
		}
	})

	t.Run("Right 250 steps", func(t *testing.T) {
		input := "R250"
		commands := strings.Split(input, "\n")
		count := puzzle1(commands)
		if count != 1 {
			t.Errorf("Expected 1, got %d", count)
		}
	})
}

func TestPuzzle2(t *testing.T) {
	t.Run("Left 50 steps", func(t *testing.T) {
		input := "L50"
		commands := strings.Split(input, "\n")
		count := puzzle2(commands)
		if count != 1 {
			t.Errorf("Expected 1, got %d", count)
		}
	})

	t.Run("Left 150 steps", func(t *testing.T) {
		input := "L150"
		commands := strings.Split(input, "\n")
		count := puzzle2(commands)
		if count != 2 {
			t.Errorf("Expected 2, got %d", count)
		}
	})

	t.Run("Left 1000 steps", func(t *testing.T) {
		input := "L1000"
		commands := strings.Split(input, "\n")
		count := puzzle2(commands)
		if count != 10 {
			t.Errorf("Expected 10, got %d", count)
		}
	})

	t.Run("Right 1000 steps", func(t *testing.T) {
		input := "R1000"
		commands := strings.Split(input, "\n")
		count := puzzle2(commands)
		if count != 10 {
			t.Errorf("Expected 10, got %d", count)
		}
	})

	t.Run("Left 1000 steps, Right 1000 steps", func(t *testing.T) {
		input := "L1000\nR1000"
		commands := strings.Split(input, "\n")
		count := puzzle2(commands)
		if count != 20 {
			t.Errorf("Expected 20, got %d", count)
		}
	})
}
