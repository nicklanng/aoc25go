package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Min int
	Max int
}

func main() {
	input, err := os.ReadFile("input/day2")
	if err != nil {
		log.Fatal(err)
	}

	ranges := strings.Split(string(input), ",")

	count := puzzle1(ranges)
	fmt.Println("Puzzle 1: ", count)

	count = puzzle2(ranges)
	fmt.Println("Puzzle 2: ", count)
}

func puzzle1(rangesStr []string) int {
	sum := 0

	ranges := make([]Range, len(rangesStr))
	for i, r := range rangesStr {
		ranges[i] = parseRange(r)
	}
	for _, r := range ranges {
		for i := r.Min; i <= r.Max; i++ {
			if isSequenceRepeatedOnce(i) {
				sum += i
			}
		}
	}

	return sum
}

func puzzle2(rangesStr []string) int {
	sum := 0

	ranges := make([]Range, len(rangesStr))
	for i, r := range rangesStr {
		ranges[i] = parseRange(r)
	}
	for _, r := range ranges {
		for i := r.Min; i <= r.Max; i++ {
			if isSequenceRepeatedAtLeastOnce(i) {
				sum += i
			}
		}
	}

	return sum
}

// parseRange parses a range string into a Range struct
func parseRange(r string) Range {
	parts := strings.Split(r, "-")
	min, _ := strconv.Atoi(parts[0])
	max, _ := strconv.Atoi(parts[1])

	// defensive, not sure if needed
	if min > max {
		return Range{Min: max, Max: min}
	}

	return Range{Min: min, Max: max}
}

// isSequenceRepeatedOnce checks if a number is a repeated sequence once
func isSequenceRepeatedOnce(r int) bool {
	// parse int to string
	str := fmt.Sprintf("%d", r)

	// split the string into two halves and check if they are equal
	mid := len(str) / 2
	return str[:mid] == str[mid:]
}

// isSequenceRepeatedAtLeastOnce checks if a number is a repeated sequence at least once
func isSequenceRepeatedAtLeastOnce(r int) bool {
	// parse int to string
	str := fmt.Sprintf("%d", r)

	// factor the length of the string
	factors := factor(len(str))

	// drop the first factor, coz 1 is just the whole string
	factors = factors[1:]

	// check if any factor is a repeated sequence
	for _, f := range factors {
		chunks := splitStringIntoChunks(str, f)
		if allEqual(chunks) {
			return true
		}
	}

	return false
}

// factor returns all factors of a number
func factor(n int) []int {
	factors := []int{}
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

// splitStringIntoChunks splits a string into n chunks
func splitStringIntoChunks(s string, n int) []string {
	chunks := make([]string, n)
	for i := range n {
		chunks[i] = s[i*len(s)/n : (i+1)*len(s)/n]
	}
	return chunks
}

// allEqual checks if all elements in a slice are equal
func allEqual[T comparable](s []T) bool {
	first := s[0]
	for _, v := range s {
		if v != first {
			return false
		}
	}
	return true
}
