package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/nicklanng/aoc25go/cmd/day09/internal"
)

func main() {
	input, err := os.ReadFile("input/day9")
	if err != nil {
		log.Fatal(err)
	}

	answer1 := puzzle1(input)
	fmt.Println("Puzzle 1: ", answer1)

	answer2 := puzzle2(input)
	fmt.Println("Puzzle 2: ", answer2)
}

func puzzle1(input []byte) int {
	lines := strings.Split(string(input), "\n")

	points := make([]internal.Point, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points[i] = internal.Point{X: x, Y: y}
	}

	var maxArea float64
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			width := math.Abs(float64(points[i].X-points[j].X)) + 1
			height := math.Abs(float64(points[i].Y-points[j].Y)) + 1
			area := width * height
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return int(maxArea)
}

func puzzle2(input []byte) int {
	return 0
}
