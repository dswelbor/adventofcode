package day_six

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/dswelbor/adventofcode/aoc2023/utility"
)

type RaceRecord struct {
	distance  int
	totalTime int
}

// Calculates winning moves (or milliseconds to accelerate) to break current RaceRecord
func (r *RaceRecord) Moves() *[]int {
	// calc bounds for min and max times
	time1 := (float64((-1)*(r.totalTime)) + math.Sqrt(float64((r.totalTime*r.totalTime)-(4*r.distance)))) / (-2) // prolly min
	time2 := (float64((-1)*(r.totalTime)) - math.Sqrt(float64((r.totalTime*r.totalTime)-(4*r.distance)))) / (-2) // prolly max
	min := math.Min(time1, time2)
	max := math.Max(time1, time2)
	// Edge case - min and max evaluated to exactly whole numbers. To beat the record,
	// these are exclusive bounds - increment/decrement min/max
	if min == math.Trunc(min) {
		min++
	}
	if max == math.Trunc(max) {
		max--
	}

	// round min up and round max down to stay within record setting/winning range
	minTime := int(math.Ceil(min))
	maxTime := int(math.Floor(max))

	// Add int times to list as moves
	timesCount := maxTime - minTime + 1
	winAccelTimes := make([]int, timesCount)
	for i := 0; i < timesCount; i++ {
		time := minTime + i
		winAccelTimes[i] = time
	}

	return &winAccelTimes

}

// High level entry Point for Day 4 solution
func SolveDaySix(input *[]string, part int) {

	if part == 1 {
		solvePartOne(input)
	} else if part == 2 {
		solvePartTwo(input)
	} else {
		fmt.Println("Part: " + strconv.Itoa(part) + "Not supported")
	}

}

// Entry point for day 6 part 1 solution
func solvePartOne(input *[]string) {
	fmt.Println("--- Solving Day Six - Part One! ---")

	// Fetch Records
	raceRecords := parseRaceRecords(input)
	//fmt.Println("[DEBUG] List of Record objects", raceRecords)
	// Fetch winning move counts
	moveCounts := listMoveCounts(raceRecords)
	// fmt.Println("[DEBUG] These are the expected possible winning move counts: ", moveCounts)
	// Calculate margin for error - multiple all elements of moveCounts
	errMargin := utility.MultipleNumbers(moveCounts)
	fmt.Println("Margin for error (product of winning move counts): ", errMargin)

}

// Entry point for day 6 part 2 solution
func solvePartTwo(input *[]string) {
	fmt.Println("--- Solving Day Six - Part Two! ---")
}

func listMoveCounts(raceRecords *[]RaceRecord) *[]int {
	// Iterate through races and add move() counts to list
	raceCount := len(*raceRecords)
	moveCounts := make([]int, raceCount)
	for i := 0; i < raceCount; i++ {
		// Get moves
		raceRecord := (*raceRecords)[i]
		moves := raceRecord.Moves()
		// Count moves
		moveCount := len(*moves)
		moveCounts[i] = moveCount
	}
	return &moveCounts
}

// Utility functions iterates through input and builds a list of RaceRecords
func parseRaceRecords(input *[]string) *[]RaceRecord {
	// parse out the times and distance into parallel arrays
	numReg := regexp.MustCompile("\\d+")
	timeStrings := make([]string, 0)
	distStrings := make([]string, 0)
	for _, inputStr := range *input {
		if strings.HasPrefix(inputStr, "Time:") {
			// Time input
			timeStrings = numReg.FindAllString(inputStr, -1)
		} else if strings.HasPrefix(inputStr, "Distance:") {
			// Distance input
			distStrings = numReg.FindAllString(inputStr, -1)

		} else {
			fmt.Println("[DEBUG] Didn't recognize input: ", inputStr)
		}
	}
	// Translate parallel arrays into list of RaceRecords
	raceCount := len(timeStrings)
	raceRecords := make([]RaceRecord, raceCount)
	for i := 0; i < raceCount; i++ {
		// Parse number strings into int
		time, tErr := strconv.Atoi(timeStrings[i])
		dist, dErr := strconv.Atoi(distStrings[i])
		if tErr != nil || dErr != nil {
			fmt.Println("[ERROR] Encountered error when parsing time: ", timeStrings[i],
				" or distance: ", timeStrings[i])
		}
		// build RaceRecord struct and add to list
		raceRecord := RaceRecord{totalTime: time, distance: dist}
		raceRecords[i] = raceRecord
	}
	// ship it!
	return &raceRecords
}
