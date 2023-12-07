package day_seven

import (
	"fmt"
	"strconv"
)

type RaceRecord struct {
	distance  int
	totalTime int
}

// High level entry Point for Day 7 solution
func SolveDaySeven(input *[]string, part int) {

	if part == 1 {
		solvePartOne(input)
	} else if part == 2 {
		solvePartTwo(input)
	} else {
		fmt.Println("Part: " + strconv.Itoa(part) + "Not supported")
	}

}

// Entry point for day 7 part 1 solution
func solvePartOne(input *[]string) {
	fmt.Println("--- Solving Day Seven - Part One! ---")
}

// Entry point for day 7 part 2 solution
func solvePartTwo(input *[]string) {
	fmt.Println("--- Solving Day Seven - Part Two! ---")
}
