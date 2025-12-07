package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	dropper, err := os.ReadFile("input/day7")
	if err != nil {
		log.Fatal(err)
	}

	answer1 := puzzle1(dropper)
	fmt.Println("Puzzle 1: ", answer1)

	answer2 := puzzle2(dropper)
	fmt.Println("Puzzle 2: ", answer2)
}

func puzzle1(maths []byte) int {
	var answer int

	levels := bytes.Split(maths, []byte("\n"))

	// init beams array
	beams := make([]bool, len(levels[0]))
	nextGenBeams := make([]bool, len(levels[0]))

	// find starting beam location
	for i, char := range levels[0] {
		if char == 'S' {
			beams[i] = true
			break
		}
	}

	for i := 1; i < len(levels); i++ {
		level := levels[i]
		for j, beam := range beams {
			// if theres no beam on the previous line in this position, skip
			if !beam {
				continue
			}

			// if we've hit a splitter, add beams to the left and right
			if level[j] == '^' {
				nextGenBeams[j-1] = true
				nextGenBeams[j+1] = true
				answer++
			} else {
				// otherwise, just add a beam in the same position
				nextGenBeams[j] = true
			}
		}
		beams = nextGenBeams
		nextGenBeams = make([]bool, len(levels[0]))
	}

	return answer
}

func puzzle2(maths []byte) int {
	levels := bytes.Split(maths, []byte("\n"))

	// remove all lines that are just dots
	filteredLevels := make([][]byte, 0, len(levels))
	for _, level := range levels {
		for _, char := range level {
			if char == '.' {
				continue
			}
			filteredLevels = append(filteredLevels, level)
			break
		}
	}
	levels = filteredLevels

	// init beams array
	beams := make([]int, len(levels[0]))
	nextGenBeams := make([]int, len(levels[0]))

	// find starting beam location
	for i, char := range levels[0] {
		if char == 'S' {
			beams[i] = 1
			break
		}
	}

	// process each level
	for i := 1; i < len(levels); i++ {
		level := levels[i]
		for j := 0; j < len(level); j++ {
			// if we've hit a splitter, add superpositioned beams to the left and right
			if level[j] == '^' {
				nextGenBeams[j-1] += beams[j]
				nextGenBeams[j+1] += beams[j]
			} else {
				// otherwise, just add the superpositioned beams in the same position
				nextGenBeams[j] += beams[j]
			}
		}
		beams = nextGenBeams
		nextGenBeams = make([]int, len(levels[0]))
	}

	var answer int
	for _, beam := range beams {
		answer += beam
	}
	return answer
}
