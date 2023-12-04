package day_four

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/dswelbor/adventofcode/aoc2023/utility"
)

type MatchVisitor struct {
	power int
	base  int
}

func (v *MatchVisitor) Visit(match bool) int {
	points := math.Pow(float64(v.base), float64(v.power))
	if match {
		v.power += 1
	}
	return int(points)
}

func SolveDayFour(input *[]string, part int) {

	if part == 1 {
		solvePartOne(input)
	} else if part == 2 {
		solvePartTwo(input)
	} else {
		fmt.Println("Part: " + strconv.Itoa(part) + "Not supported")
	}

}

func solvePartOne(input *[]string) {
	fmt.Println("--- Solving Day Four - Part One! ---")
	// grab points based on win matches
	points := listCardPoints(input)
	// calc sum from collected winning points
	pointSum := utility.SumNumbers(points)
	fmt.Println("Sum of win points: ", pointSum)
}

func solvePartTwo(input *[]string) {
	fmt.Println("--- Solving Day Four - Part Two! ---")
}

func listCardPoints(input *[]string) *[]int {
	// init points list and visitor for tracking points across cards
	pointsList := make([]int, 0)
	for _, cardInputStr := range *input {
		// split Card # from numbers on ":"
		allNumbers := strings.Split(cardInputStr, ":")
		// split winning #'s from scratched numbers on "|"
		numbers := strings.Split(allNumbers[1], "|")

		// Grab a list of winning numbers and map it
		numReg := regexp.MustCompile("\\d+")
		winNumStrings := numReg.FindAllString(numbers[0], -1)
		winMap := make(map[int]bool)
		for _, numStr := range winNumStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("[ERROR] Error parsing: \"", numStr, "\" error: ", err)
			} else {
				// no error - normal case
				winMap[num] = true
			}
		}

		// List scratched off numbers
		numStrings := numReg.FindAllString(numbers[1], -1)

		// visit each scratched off number, and if matched, square current points
		points := 0
		visitor := MatchVisitor{
			base:  2,
			power: 0,
		}
		for _, numStr := range numStrings {
			// init match flag
			match := false

			// parse number string into int
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("[ERROR] Error parsing: \"", numStr, "\" error: ", err)
			} else {
				// no error - normal case
				match = winMap[num]
				// TODO: Refactor in the visit behavior
				if match {
					points = visitor.Visit(match)
				}
			}
		}
		pointsList = append(pointsList, points)

	}
	return &pointsList
}
