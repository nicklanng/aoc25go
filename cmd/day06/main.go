package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Operator int

func main() {
	maths, err := os.ReadFile("input/day6")
	if err != nil {
		log.Fatal(err)
	}

	answer1 := puzzle1(string(maths))
	fmt.Println("Puzzle 1: ", answer1)

	answer2 := puzzle2(maths)
	fmt.Println("Puzzle 2: ", answer2)
}

func puzzle1(maths string) int {
	// reverse the lines so we can process them from the bottom up
	lines := strings.Split(maths, "\n")
	slices.Reverse(lines)

	// the first line is the operators
	operatorLine := lines[0]
	operators := strings.Fields(operatorLine)

	// find the operators for each column and init their result space
	operatorVals := make([]byte, len(operators))
	results := make([]int, len(operators))

	// parse the operators
	for i, operator := range operators {
		switch v := operator; v {
		case "+":
			operatorVals[i] = v[0]
			results[i] = 0
		case "*":
			operatorVals[i] = v[0]
			results[i] = 1
		}
	}

	// process the lines from bottom to top
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		values := strings.Fields(line)

		for j, value := range values {
			valueInt, _ := strconv.Atoi(value)
			operation := opToFunc(operatorVals[j])
			results[j] = operation(results[j], valueInt)
		}
	}

	var answer int
	for _, result := range results {
		answer += result
	}
	return answer
}

func puzzle2(maths []byte) int {
	// split the maths into lines and transpose them
	lines := bytes.Split(maths, []byte("\n"))
	transposed := transpose(lines)

	// result for each block
	results := make([]int, 1)

	// per block vars
	var (
		block     int
		firstLine bool = true
		operation func(a, b int) int
	)

	// process each line
	for _, line := range transposed {
		lineStr := string(line)

		// if the line is empty, start a new block
		if strings.TrimSpace(lineStr) == "" {
			block++
			results = append(results, 0)
			firstLine = true
			continue
		}

		// if this is the first line of a block, determine the operation and init the result
		if firstLine {
			operator := lineStr[len(lineStr)-1]
			operation = opToFunc(operator)
			if operator == '*' {
				results[block] = 1
			}
			firstLine = false
		}

		// get the number from the line and apply the operation
		numberStr := strings.TrimSpace(lineStr[:len(lineStr)-1])
		number, _ := strconv.Atoi(numberStr)
		results[block] = operation(results[block], number)
	}

	// sum the results
	var answer int
	for _, result := range results {
		answer += result
	}
	return answer
}

func transpose[T any](slice [][]T) [][]T {
	xl := len(slice[0])
	yl := len(slice)

	// init slices
	result := make([][]T, xl)
	for i := range result {
		result[i] = make([]T, yl)
	}

	// transpose the slice
	for i := range xl {
		for j := range yl {
			result[i][j] = slice[j][i]
		}
	}

	return result
}

func opToFunc(op byte) func(a, b int) int {
	switch op {
	case '+':
		return func(a, b int) int { return a + b }
	case '*':
		return func(a, b int) int { return a * b }
	}
	return func(a, b int) int { return 0 }
}
