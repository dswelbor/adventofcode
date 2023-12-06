package day_five

import (
	"fmt"
	"strconv"
)

type MatchVisitor struct {
	power  int
	base   int
	points int
}

// High level entry Point for Day 4 solution
func SolveDayFive(input *[]string, part int) {

	if part == 1 {
		solvePartOne(input)
	} else if part == 2 {
		solvePartTwo(input)
	} else {
		fmt.Println("Part: " + strconv.Itoa(part) + "Not supported")
	}

}

// Entry point for day 5 part 1 solution
func solvePartOne(input *[]string) {
	fmt.Println("--- Solving Day Five - Part One! ---")
}

// Entry point for day 5 part 2 solution
func solvePartTwo(input *[]string) {
	fmt.Println("--- Solving Day Five - Part Two! ---")
}
