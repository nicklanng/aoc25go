package main

import (
	"bytes"
	"fmt"
	"log"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	Lights   uint16
	Buttons  []uint16
	Joltages [16]uint8
}

func main() {
	input, err := os.ReadFile("input/day10")
	if err != nil {
		log.Fatal(err)
	}

	answer1 := puzzle1(input)
	fmt.Println("Puzzle 1: ", answer1)

	answer2 := puzzle2(input)
	fmt.Println("Puzzle 2: ", answer2)
}

func puzzle1(input []byte) int {
	machines := parseInput(input)
	shortestPaths := make([]int, len(machines))

	for i, machine := range machines {
		shortestPaths[i] = findShortestPathToLights(machine)
	}

	var answer int
	for _, path := range shortestPaths {
		answer += path
	}
	return answer
}

func findShortestPathToLights(machine Machine) int {
	len := len(machine.Buttons)

	var buttonPresses int = 1<<len - 1

	for i := 1; i < 1<<len; i++ {
		var candidate uint16

		for j := range len {
			if i&(1<<j) != 0 {
				candidate ^= machine.Buttons[j]
			}
		}

		if candidate == machine.Lights {
			bitCount := bits.OnesCount(uint(i))
			if bitCount < buttonPresses {
				buttonPresses = bitCount
			}
		}
	}

	return buttonPresses
}

func puzzle2(input []byte) int {
	return 0
}

func parseInput(input []byte) []Machine {
	var machines []Machine

	for line := range strings.SplitSeq(string(input), "\n") {
		fields := strings.Fields(line)
		lights := parseLights(fields[0])
		buttons := parseButtons(fields[1 : len(fields)-1])
		joltages := parseJoltages(fields[len(fields)-1])
		machines = append(machines, Machine{
			Lights:   lights,
			Buttons:  buttons,
			Joltages: joltages,
		})
	}

	return machines
}

func parseLights(input string) uint16 {
	var lights uint16

	for i, char := range input[1 : len(input)-1] {
		if char == '#' {
			lights |= 1 << i
		}
	}
	return lights
}

func parseButtons(input []string) []uint16 {
	var buttons = make([]uint16, len(input))

	for i, button := range input {
		var value uint16
		bits := bytes.SplitSeq([]byte(button[1:len(button)-1]), []byte(","))
		for bit := range bits {
			shift := bit[0] - byte('0')
			value |= 1 << shift
		}
		buttons[i] = value
	}

	return buttons
}

func parseJoltages(input string) [16]uint8 {
	var joltages [16]uint8

	i := 0
	for joltageStr := range strings.SplitSeq(input[1:len(input)-1], ",") {
		joltage, _ := strconv.Atoi(joltageStr)
		joltages[i] = uint8(joltage)
		i++
	}

	return joltages
}
