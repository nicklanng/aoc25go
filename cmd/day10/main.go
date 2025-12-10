package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	Lights   uint16
	Buttons  []uint16
	Joltages [16]uint8
}

type history struct {
	lights   uint16
	length   int
	joltages [16]uint8
}

func (h history) push(button uint16) history {
	for i := range 16 {
		if button&(1<<i) != 0 {
			h.joltages[i]++
		}
	}
	return history{lights: h.lights ^ button, length: h.length + 1, joltages: h.joltages}
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
	openList := make([]history, 1)
	visited := map[uint16]struct{}{0: {}}

	for {
		if len(openList) == 0 {
			break
		}

		// take the first state from the open list
		state := openList[0]
		openList = openList[1:]

		for _, button := range machine.Buttons {
			// push the button
			newState := state.push(button)

			// if we've reached the target lights
			if newState.lights == machine.Lights {
				return newState.length
			}

			// if we're cycling, skip
			if _, ok := visited[newState.lights]; ok {
				continue
			}

			// add the new state to the visited set
			visited[newState.lights] = struct{}{}

			// add the new state to the open list
			openList = append(openList, newState)
		}
	}
	return 0
}

func puzzle2(input []byte) int {
	machines := parseInput(input)
	shortestPaths := make([]int, len(machines))

	for i, machine := range machines {
		shortestPaths[i] = findShortestPathToJoltages(machine)
	}

	var answer int
	for _, path := range shortestPaths {
		answer += path
	}
	return answer
}

func findShortestPathToJoltages(machine Machine) int {
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
