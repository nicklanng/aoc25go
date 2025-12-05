package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input/day3")
	if err != nil {
		log.Fatal(err)
	}

	selections := strings.Split(string(input), "\n")

	answer1 := puzzle1(selections)
	fmt.Println("Puzzle 1: ", answer1)

	answer2 := puzzle2(selections)
	fmt.Println("Puzzle 2: ", answer2)
}

func puzzle1(selections []string) int {
	var sum int

	for _, selectionRow := range selections {
		sum += scanRow(selectionRow, 2)
	}

	return sum
}

func puzzle2(selections []string) int {
	var sum int

	for _, selectionRow := range selections {
		sum += scanRow(selectionRow, 12)
	}

	return sum
}

// scanRow scans a row of the selection and returns the highest joltage
// by moving a sliding window across the selection
func scanRow(selectionRow string, batteryLen int) int {
	var (
		localSum   int
		startIndex int // where we start checking for the highest joltage
	)

	for i := range batteryLen {

		// which slot of the battery we're finding the highest joltage for
		batteryPos := batteryLen - i - 1

		// the furthest index we can check while saving enough batteries to fill all 12 slots
		maxIndex := len(selectionRow) - batteryPos

		// find the highest joltage in the remaining selection
		maxJoltage, foundIndex, err := findHighestJoltage(selectionRow[startIndex:maxIndex])
		if err != nil {
			panic(err)
		}

		// multiply by factor of 10 to get the correct position
		localSum += maxJoltage * int(math.Pow(10, float64(batteryPos)))
		startIndex += foundIndex + 1
	}
	return localSum
}

// finds the highest character in the string and converts it to an integer
func findHighestJoltage(selection string) (int, int, error) {
	highest := '0'
	highestIndex := 0

	// find the highest joltage in the selection
	for i, v := range selection {
		if v > highest {
			highest = v
			highestIndex = i
		}
		if highest == '9' {
			break
		}
	}

	// convert the highest joltage to an integer
	joltage, err := strconv.Atoi(string(highest))
	if err != nil {
		return 0, 0, err
	}

	return joltage, highestIndex, nil
}
