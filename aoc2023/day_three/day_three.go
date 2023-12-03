package day_three

import (
	"fmt"
	"strconv"
)

func SolveDayThree(input *[]string, part int) {

	if part == 1 {
		solvePartOne(input)
	} else if part == 2 {
		solvePartTwo(input)
	} else {
		fmt.Println("Part: " + strconv.Itoa(part) + "Not supported")
	}

}

func solvePartOne(input *[]string) {
	fmt.Println("--- Solving Day Three - Part One! ---")
	// TODO: Implement me
}

func solvePartTwo(input *[]string) {
	fmt.Println("--- Solving Day Three - Part Two! ---")
	// TODO: Implement me

}
