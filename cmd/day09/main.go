package main

import (
	"bytes"
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
	points := make([]internal.Point, 0, bytes.Count(input, []byte("\n")))
	for line := range strings.SplitSeq(string(input), "\n") {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, internal.Point{X: float64(x), Y: float64(y)})
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
	// parse points
	points := make([]internal.Point, 0, bytes.Count(input, []byte("\n"))+1)
	for line := range bytes.SplitSeq(input, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}

		// Find comma position
		commaIdx := bytes.IndexByte(line, ',')
		if commaIdx == -1 {
			continue
		}

		// Parse x and y directly from bytes
		x, _ := strconv.Atoi(string(line[:commaIdx]))
		y, _ := strconv.Atoi(string(line[commaIdx+1:]))
		points = append(points, internal.Point{X: float64(x), Y: float64(y)})
	}

	// create lines outlining the tiles from the points
	lines := make([]internal.Line, 0, len(points))
	for i := range points {
		lines = append(lines, internal.Line{Start: points[i], End: points[(i+1)%len(points)]})
	}

	var maxArea float64
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			// ignore rectangles that are just a line
			if points[i].X == points[j].X || points[i].Y == points[j].Y {
				continue
			}

			// create the rectangle
			rect := internal.NewRectangle(points[i].X, points[i].Y, points[j].X, points[j].Y)
			area := rect.Area()

			// skip rectangle if its smaller than the current max area
			if area <= maxArea {
				continue
			}

			shrunkRect := rect.Shrink(0.1)
			if isValidRectangle(shrunkRect, lines) {
				maxArea = area
			}

		}
	}

	return int(maxArea)
}

func isValidRectangle(rectangle internal.Rectangle, lines []internal.Line) bool {
	var rectangleLines [4]internal.Line
	rectangleLines[0] = internal.Line{Start: internal.Point{X: rectangle.MinX, Y: rectangle.MinY}, End: internal.Point{X: rectangle.MaxX, Y: rectangle.MinY}}
	rectangleLines[1] = internal.Line{Start: internal.Point{X: rectangle.MaxX, Y: rectangle.MinY}, End: internal.Point{X: rectangle.MaxX, Y: rectangle.MaxY}}
	rectangleLines[2] = internal.Line{Start: internal.Point{X: rectangle.MaxX, Y: rectangle.MaxY}, End: internal.Point{X: rectangle.MinX, Y: rectangle.MaxY}}
	rectangleLines[3] = internal.Line{Start: internal.Point{X: rectangle.MinX, Y: rectangle.MaxY}, End: internal.Point{X: rectangle.MinX, Y: rectangle.MinY}}

	// check if the rectangle is completely contained within the lines
	for _, line := range rectangleLines {
		for _, line2 := range lines {
			if line.Crosses(line2) {
				return false
			}
		}
	}
	return true
}
