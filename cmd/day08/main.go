package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/nicklanng/aoc25go/cmd/day08/internal"
)

// point represents a point in 3D space
type point struct {
	x float64
	y float64
	z float64
}

// dist2 returns the squared distance between two points
func (p1 point) dist2(p2 point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	dz := p1.z - p2.z
	return dx*dx + dy*dy + dz*dz
}

// edge represents an edge between two points
type edge struct {
	i1  int
	i2  int
	len float64
}

func main() {
	input, err := os.ReadFile("input/day8")
	if err != nil {
		log.Fatal(err)
	}

	answer1 := puzzle1(input, 1000)
	fmt.Println("Puzzle 1: ", answer1)

	answer2 := puzzle2(input)
	fmt.Println("Puzzle 2: ", answer2)
}

func puzzle1(input []byte, n int) int {
	points := parsePoints(input)
	dsu := internal.NewDisjointSetUnion(len(points))

	// find edges and sort by len
	edges := findEdges(points)
	slices.SortFunc(edges, func(a, b edge) int {
		return int(a.len - b.len)
	})

	// connect 10 edges
	for i := range n {
		dsu.Union(edges[i].i1, edges[i].i2)
	}

	sizes := findCircuitSizes(points, dsu)
	slices.SortFunc(sizes, func(a, b int) int {
		return b - a
	})

	return sizes[0] * sizes[1] * sizes[2]
}

func puzzle2(input []byte) int {
	points := parsePoints(input)

	dsu := internal.NewDisjointSetUnion(len(points))

	// find edges and sort by len
	edges := findEdges(points)
	slices.SortFunc(edges, func(a, b edge) int {
		return int(a.len - b.len)
	})

	// connect edges until there is only one set
	var lastEdge edge
	for i := range edges {
		dsu.Union(edges[i].i1, edges[i].i2)
		if dsu.Sets() == 1 {
			lastEdge = edges[i]
			break
		}
	}

	return int(points[lastEdge.i1].x * points[lastEdge.i2].x)
}

// parsePoints parses the input into a slice of points
func parsePoints(input []byte) []point {
	lines := strings.Split(string(input), "\n")

	points := make([]point, 0, len(lines))
	for _, line := range lines {
		fields := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(fields[0], 64)
		y, _ := strconv.ParseFloat(fields[1], 64)
		z, _ := strconv.ParseFloat(fields[2], 64)
		points = append(points, point{x: x, y: y, z: z})
	}

	return points
}

// findEdges finds all edges between points
func findEdges(points []point) []edge {
	edges := make([]edge, 0, len(points)*(len(points)-1)/2)

	for i := range points {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, edge{
				i1:  i,
				i2:  j,
				len: points[i].dist2(points[j]), // use squared distance for performance
			})
		}
	}

	return edges
}

// findCircuitSizes finds the sizes of the circuits in the graph
func findCircuitSizes(points []point, dsu *internal.DisjointSetUnion) []int {
	// create a map of root point to circuit size to dedupe points in the same set
	circuitSizes := make(map[int]int)
	for i := range points {
		root := dsu.Find(i)
		circuitSizes[root] = dsu.Size(root)
	}

	// turn the map into a slice and sort by size
	sizes := make([]int, 0, len(circuitSizes))
	for _, size := range circuitSizes {
		sizes = append(sizes, size)
	}

	return sizes
}
