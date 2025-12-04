package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nicklanng/aoc25go/cmd/day04/internal"
)

func main() {
	input, err := os.ReadFile("input/day4")
	if err != nil {
		log.Fatal(err)
	}

	answer1 := puzzle1(string(input))
	fmt.Println("Puzzle 1: ", answer1)

	answer2 := puzzle2(string(input))
	fmt.Println("Puzzle 2: ", answer2)
}

func puzzle1(gridStr string) int {
	answer := 0

	grid := internal.NewGrid(gridStr)

	// scan the grid
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			// if theres no roll here, skip
			if !grid.HasRoll(x, y) {
				continue
			}

			// count the number of neighbors with rolls
			neighborCount := 0
			neighbors := grid.Neighbors(x, y)
			for neighbor := range neighbors {
				if grid.HasRoll(neighbor.X, neighbor.Y) {
					neighborCount++
				}
			}

			// if the number of neighbors with rolls is less than 4, add 1 to the answer
			if neighborCount < 4 {
				answer++
			}
		}
	}

	return answer
}

func puzzle2(gridStr string) int {
	answer := 0

	grid := internal.NewGrid(gridStr)

	// keep removing rolls until no more rolls accessible
	for {
		var removed int

		// scan the grid
		for y := 0; y < grid.Height; y++ {
			for x := 0; x < grid.Width; x++ {
				// if theres no roll here, skip
				if !grid.HasRoll(x, y) {
					continue
				}

				// count the number of neighbors with rolls
				neighborCount := 0
				neighbors := grid.Neighbors(x, y)
				for neighbor := range neighbors {
					if grid.HasRoll(neighbor.X, neighbor.Y) {
						neighborCount++
					}
				}

				// if the number of neighbors with rolls is less than 4, remove the roll and increment the removed count
				if neighborCount < 4 {
					grid.ClearRoll(x, y)
					removed++
				}
			}
		}

		// if no rolls were removed, then we're done
		if removed == 0 {
			break
		}

		// add the number of removed rolls to the answer
		answer += removed
	}

	return answer
}
