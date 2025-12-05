package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nicklanng/aoc25go/cmd/day05/internal"
)

func main() {
	input, err := os.ReadFile("input/day5")
	if err != nil {
		log.Fatal(err)
	}

	answer1 := puzzle1(string(input))
	fmt.Println("Puzzle 1: ", answer1)

	answer2 := puzzle2(string(input))
	fmt.Println("Puzzle 2: ", answer2)
}

func puzzle1(database string) int {
	var freshIds int

	db, err := internal.ParseDatabase(database)
	if err != nil {
		panic(err)
	}

	for _, id := range db.Ids {
		if db.FindId(id) {
			freshIds++
			continue
		}
	}

	return freshIds
}

func puzzle2(database string) int {
	db, err := internal.ParseDatabase(database)
	if err != nil {
		panic(err)
	}

	return db.CountAvailableIds()
}
