package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nicklanng/aoc25go/internal/data"
)

func main() {
	input, err := os.ReadFile("input/day2")
	if err != nil {
		log.Fatal(err)
	}

	ranges := strings.Split(string(input), ",")

	answer1 := puzzle1(ranges)
	fmt.Println("Puzzle 1: ", answer1)

	answer2 := puzzle2(ranges)
	fmt.Println("Puzzle 2: ", answer2)
}

func puzzle1(rangesStr []string) int {
	sum := 0

	ranges := make([]data.Range, len(rangesStr))
	for i, r := range rangesStr {
		var err error
		ranges[i], err = data.ParseRange(r)
		if err != nil {
			panic(err)
		}
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

	ranges := make([]data.Range, len(rangesStr))
	for i, r := range rangesStr {
		var err error
		ranges[i], err = data.ParseRange(r)
		if err != nil {
			panic(err)
		}
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
