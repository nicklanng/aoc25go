package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Operator int

const (
	OpAdd Operator = iota
	OpMul
)

func main() {
	maths, err := os.ReadFile("input/day6")
	if err != nil {
		log.Fatal(err)
	}

	answer1 := puzzle1(string(maths))
	fmt.Println("Puzzle 1: ", answer1)

	answer2 := puzzle2(string(maths))
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
	operatorVals := make([]Operator, len(operators))
	results := make([]int, len(operators))

	// parse the operators
	for i, operator := range operators {
		switch operator {
		case "+":
			operatorVals[i] = OpAdd
			results[i] = int(OpAdd)
		case "*":
			operatorVals[i] = OpMul
			results[i] = int(OpMul)
		}
	}

	// process the lines from bottom to top
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		values := strings.Fields(line)

		for j, value := range values {
			valueInt, _ := strconv.Atoi(value)
			switch operatorVals[j] {
			case OpAdd:
				results[j] += valueInt
			case OpMul:
				results[j] *= valueInt
			}
		}
	}

	answer := 0
	for _, result := range results {
		answer += result
	}
	return answer
}

func puzzle2(maths string) int {
	lines := strings.Split(maths, "\n")
	numberLineCount := len(lines) - 1

	var operators []Operator
	var operatorIndices []int
	var results []int

	// parse the operators and their indices from the last line
	for i, char := range lines[len(lines)-1] {
		switch char {
		case '+':
			operators = append(operators, OpAdd)
			operatorIndices = append(operatorIndices, i)
			results = append(results, int(OpAdd))
		case '*':
			operators = append(operators, OpMul)
			operatorIndices = append(operatorIndices, i)
			results = append(results, int(OpMul))
		}
	}

	// add the end of the line to the operator indicies
	operatorIndices = append(operatorIndices, len(lines[len(lines)-1])+1)

	// process each block of maths
	for block := 0; block < len(operatorIndices)-1; block++ {
		// get the start and end indices of the current block
		operatorIndex := operatorIndices[block]
		operatorIndexEnd := operatorIndices[block+1] - 1

		// determine the operation for this block
		var operation func(a, b int) int
		switch operators[block] {
		case OpAdd:
			operation = func(a, b int) int { return a + b }
		case OpMul:
			operation = func(a, b int) int { return a * b }
		}

		// process the numbers in this block
		for col := operatorIndex; col < operatorIndexEnd; col++ {

			// get the numbers in this column
			numberStr := ""
			for row := range numberLineCount {
				numberStr += string(lines[row][col])
			}

			// convert the number string to an integer and apply the operation
			number, _ := strconv.Atoi(strings.TrimSpace(numberStr))
			results[block] = operation(results[block], number)
		}
	}

	answer := 0
	for _, result := range results {
		answer += result
	}
	return answer
}
