package internal

import (
	"iter"
	"strings"
)

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

	height := len(lines)
	width := len(lines[0])

	data := make([]bool, height*width)
	for y, line := range lines {
		for x, char := range line {
			data[y*width+x] = char == '@'
		}
	}

	return Grid{data: data, Width: width, Height: height}
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

func (g Grid) Neighbors(x, y int) iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}
				if !yield(Point{X: x + dx, Y: y + dy}) {
					return
				}
			}
		}
	}
}
