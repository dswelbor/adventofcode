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

// Calculates winning moves (or milliseconds to accelerate) to break current
// RaceRecord. This relies on:
//
//	Tt=Tc+Tm, d<=v*Tm, v=Tc*1
//
// and solving for Tc where:
//
//	Tt is total time, Tc is time charging, Tm is time moving
//	d is distance, and v is velocity
//
// With those equations known, we get 0 <= (-1)*(Tc)^2 + Tt*Tc + (-1)*d
// and we can solve for Tc with Tt and d known for a given race
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

	// Add int Tc times to list as moves
	timesCount := maxTime - minTime + 1
	winAccelTimes := make([]int, timesCount)
	for i := 0; i < timesCount; i++ {
		time := minTime + i
		winAccelTimes[i] = time
	}

	return &winAccelTimes

}

// High level entry Point for Day 6 solution
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
	raceRecords := parseRaceRecords(input, false)
	// Fetch winning move counts
	moveCounts := listMoveCounts(raceRecords)
	// Calculate margin for error - multiple all elements of moveCounts
	errMargin := utility.MultipleNumbers(moveCounts)
	fmt.Println("Margin for error (product of winning move counts): ", errMargin)

}

// Entry point for day 6 part 2 solution
func solvePartTwo(input *[]string) {
	fmt.Println("--- Solving Day Six - Part Two! ---")

	// Fetch Records - fix "kerning" by replacing spaces in input
	raceRecords := parseRaceRecords(input, true)
	// Fetch winning move counts
	moveCounts := listMoveCounts(raceRecords)
	// Calculate margin for error - multiple all elements of moveCounts
	errMargin := utility.MultipleNumbers(moveCounts)
	fmt.Println("Margin for error (product of winning move counts) with fixed kerning: ", errMargin)
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
func parseRaceRecords(input *[]string, replaceSpaces bool) *[]RaceRecord {
	// parse out the times and distance into parallel arrays
	numReg := regexp.MustCompile("\\d+")
	timeStrings := make([]string, 0)
	distStrings := make([]string, 0)
	for _, inputStr := range *input {
		// Edge case - fix kerning by replacing all whitespace
		if replaceSpaces {
			spaceReg := regexp.MustCompile("\\s+")
			inputStr = spaceReg.ReplaceAllString(inputStr, "")
		}
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
