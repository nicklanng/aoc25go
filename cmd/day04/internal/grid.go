package internal

import "strings"

type Point struct {
	X int
	Y int
}

type Grid struct {
	data []bool

	Width  int
	Height int
}

func NewGrid(gridStr string) Grid {
	lines := strings.Split(gridStr, "\n")
	maxY := len(lines)
	maxX := len(lines[0])
	data := make([]bool, maxY*maxX)
	for y, line := range lines {
		for x, char := range line {
			data[y*maxX+x] = char == '@'
		}
	}
	return Grid{data: data, Width: maxX, Height: maxY}
}

func (g Grid) HasRoll(x, y int) bool {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return false
	}
	return g.data[y*g.Width+x]
}

func (g Grid) ClearRoll(x, y int) {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return
	}
	g.data[y*g.Width+x] = false
}

func (g Grid) Neighbors(x, y int) []Point {
	neighbors := make([]Point, 0, 8)
	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			if dx == 0 && dy == 0 {
				continue
			}
			neighbors = append(neighbors, Point{X: x + dx, Y: y + dy})
		}
	}
	return neighbors
}
