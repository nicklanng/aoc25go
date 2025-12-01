package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input/day1")
	if err != nil {
		log.Fatal(err)
	}

	commands := strings.Split(string(input), "\n")

	count := puzzle1(commands)
	fmt.Println("Puzzle 1: ", count)

	count = puzzle2(commands)
	fmt.Println("Puzzle 2: ", count)
}

func puzzle1(commands []string) int {
	count := 0
	pos := 50

	for _, command := range commands {
		dir, stepsStr := command[0], command[1:]
		steps, _ := strconv.Atoi(stepsStr)

		switch dir {
		case 'L':
			pos -= steps
			for {
				if pos < 0 {
					pos += 100
				} else {
					break
				}
			}
		case 'R':
			pos += steps
			for {
				if pos > 99 {
					pos -= 100
				} else {
					break
				}
			}
		}

		if pos == 0 {
			count++
		}
	}

	return count
}

func puzzle2(commands []string) int {
	count := 0
	pos := 50

	for _, command := range commands {
		dir, stepsStr := command[0], command[1:]
		steps, _ := strconv.Atoi(stepsStr)

		switch dir {
		case 'L':
			for range steps {
				pos -= 1
				if pos == 0 {
					count++
				}
				if pos == -1 {
					pos = 99
				}
			}
		case 'R':
			for range steps {
				pos += 1
				if pos == 100 {
					pos = 0
					count++
				}
			}
		}
	}

	return count
}
