package day_five

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"

	"github.com/dswelbor/adventofcode/aoc2023/utility"
)

type Translator struct {
	TranslatorType string
	fromRanges     *[][]int
	toRanges       *[][]int
}

func (t *Translator) Translate(fromId int) int {
	// TODO: Implement me
	// iterate through from or src ranges
	toId := fromId // if id not in range, default fromId maps to toId
	for i, fromRange := range *t.fromRanges {
		incMin := fromRange[0]
		excMax := fromRange[1]
		// check if id is in range
		if fromId >= incMin && fromId < excMax {
			// from id is in range! calculate offset to find toId
			offset := fromId - incMin
			toRanges := *t.toRanges
			toRange := toRanges[i]
			toStartId := toRange[0]
			toId = toStartId + offset
			return toId
		}
	}
	return toId
}

func (t *Translator) AddRange(fromStart int, toStart int, rangeLength int) {
	// note: these ranges are [idStart, idEnd). That is the first element is an
	// inclusive min and the second element is an exclusive max
	// for instance: fromStart = 50, rangeLength=2 will include: [50, 51] however, this
	// will be represented by the fromRange=[50, 52]

	// edge case - ranges haven't been initialized yet
	if t.fromRanges == nil || t.toRanges == nil {
		// let's init these ranges
		initFromRanges := make([][]int, 0)
		initToRanges := make([][]int, 0)
		t.fromRanges = &initFromRanges
		t.toRanges = &initToRanges
	}

	// Create from/to ranges
	fromRange := make([]int, 2)
	fromRange[0] = fromStart               // this is an inclusive min
	fromRange[1] = fromStart + rangeLength // this is an exclusive max
	toRange := make([]int, 2)
	toRange[0] = toStart
	toRange[1] = toStart + rangeLength

	// update translator with edges
	fromRanges := *t.fromRanges
	toRanges := *t.toRanges
	fromRanges = append(fromRanges, fromRange)
	toRanges = append(toRanges, toRange)
	t.fromRanges = &fromRanges
	t.toRanges = &toRanges
}

// High level entry Point for Day 5 solution
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
	minLocId := naiveLowestLocation(input)
	fmt.Println("Lowest Location Id for initial seeds: ", minLocId)
}

// Entry point for day 5 part 2 solution
func solvePartTwo(input *[]string) {
	fmt.Println("--- Solving Day Five - Part Two! ---")
	// Let's time how long solution takes to complete
	start := time.Now()
	minLocId := rangedLowestLocation(input)
	elapsed := time.Since(start)
	fmt.Println("\nComplete! Lowest Location Id for seeds ranges: ", minLocId)
	fmt.Printf("Time elapsed: %s\n", elapsed)
}

func naiveLowestLocation(input *[]string) int {
	// get seed ids - we assume this is the first line of input
	numReg := regexp.MustCompile("\\d+")
	seedIdStrings := numReg.FindAllString((*input)[0], -1)
	seedIds := utility.ListAtoi(&seedIdStrings)

	// get ordered list of translators
	translators := initTranslators(input)

	// map seed ids to location ids and find lowest location id
	seedLocMap := translateSeeds(seedIds, translators)
	minLocId := math.MaxInt
	for _, locId := range *seedLocMap {
		if locId < minLocId {
			minLocId = locId
		}
	}
	// lowest location id found
	return minLocId
}

func rangedLowestLocation(input *[]string) int {
	// get seed ids - we assume this is the first line of input
	numReg := regexp.MustCompile("\\d+")
	seedIdStrings := numReg.FindAllString((*input)[0], -1)
	seedIds := utility.ListAtoi(&seedIdStrings)
	seedIdRanges := parseSeedIdRanges(seedIds)

	// get ordered list of translators
	translators := initTranslators(input)

	// map seed ids to location ids and find lowest location id
	minLocId := minLocFromSeedRanges(seedIdRanges, translators)
	// lowest location id found
	return minLocId
}

// Parses out a collection of seed ranges from list of numbers. Assumes every
// even-indexed element is an inclusive range min and each odd-indexed element is
// a positive offset. The sum of these two numbers gives an exclusive max for the range
// in the format: [incMin, incMin + offset)
func parseSeedIdRanges(seedNumbersPtr *[]int) *[][]int {
	seedNumbers := *seedNumbersPtr
	// iterate numers 2 at a time - build ranges
	seedNumCount := len(seedNumbers)
	seedRanges := make([][]int, seedNumCount/2)
	for i := 0; i < seedNumCount; i += 2 {
		// calc min and max for seed id range
		incMin := seedNumbers[i]
		offset := seedNumbers[i+1]
		excMax := incMin + offset
		// create seed range and add
		seedRange := []int{incMin, excMax}
		seedRanges[i/2] = seedRange
	}
	fmt.Println("[INFO] Finished parsing ", seedNumCount/2, " seedId ranges")
	// We've got our list of seed ranges
	return &seedRanges
}

// Turns an id range with an inclusive min and exclusive max [minId, maxId) into a list
// of seed ids
func parseSeedIdsFromRanges(seedRange *[]int) *[]int {
	// parse inclusive min and exclusive max from range
	incMin := (*seedRange)[0]
	excMax := (*seedRange)[1]
	length := excMax - incMin
	// blow out id list
	seedIds := make([]int, length)
	for offset := 0; offset < length; offset++ {
		seedId := incMin + offset
		seedIds[offset] = seedId
	}
	// we've got our seed ids
	return &seedIds
}

// creates an ordered list of translators to go from seed id to location id
func initTranslators(input *[]string) *[]Translator {
	// init variables to persist parsed data across multiple lines of input
	var translator Translator
	var translatorType string
	numReg := regexp.MustCompile("\\d+")
	mapReg := regexp.MustCompile("\\w+\\-to\\-\\w+")

	// let's iterate through - build ordered list of translator objects (behaviors)
	translators := make([]Translator, 0)
	for i, inputStr := range *input {
		// special case - seed id list, ignore i=0
		if i == 0 {
			continue
		}

		// Parse numbers and from-to-dest mapping
		translatorMatch := mapReg.FindString(inputStr)
		numberMatches := numReg.FindAllString(inputStr, -1)
		if len(translatorMatch) > 0 {
			// Were's on a fromType-to-destType mapping line
			translatorType = translatorMatch // set the type - persist some iterations
			translator = Translator{TranslatorType: translatorType}
		} else if len(numberMatches) > 0 {
			// parse input string with '<<dest>> <<source>> <<length>>' data
			destStartId, destErr := strconv.Atoi(numberMatches[0])
			srcStartId, srcErr := strconv.Atoi(numberMatches[1])
			rangeLength, rngErr := strconv.Atoi(numberMatches[2])
			if destErr != nil || srcErr != nil || rngErr != nil {
				fmt.Println("[ERROR] Couldn't parse numbers from: ", numberMatches)
			}
			// Add range to translator
			translator.AddRange(srcStartId, destStartId, rangeLength)
		} else if len(translator.TranslatorType) > 0 || i == len(*input)-1 {
			// translator exists and no regex matches
			// no more dest src range to addto translator - add translator to list
			translators = append(translators, translator)
		}
		// edge case - input doesn't have an empty line at the end
		if i == len(*input)-1 {
			translators = append(translators, translator)
		}
	}
	fmt.Println("[INFO] Finished building ", len(translators), " translators from input")
	// Done building ordered list of translators (translate behaviors)
	return &translators
}

// Helper function that iteratively maps seedId to locationId
func translateSeeds(seedIds *[]int, translators *[]Translator) *map[int]int {
	// iterate through the list of seed ids and map seedId: locationId
	seedLocMap := make(map[int]int)
	for _, seedId := range *seedIds {
		locId := translateSeed(seedId, translators)
		seedLocMap[seedId] = locId
	}
	return &seedLocMap
}

// Helper function that takes a seed id and a list of ordered translators, and
// translates seedId to locationId.
func translateSeed(seedId int, translators *[]Translator) int {
	// iterate through translators to get location id
	locId := seedId
	id := seedId
	for _, translator := range *translators {
		locId = translator.Translate(id)
		id = locId
	}
	return locId
}

// Helper function that iteratively maps seedId to locationId
func minLocFromSeedRanges(seedIdRangesPtr *[][]int, translators *[]Translator) int {
	// iterate through the list of seed ids and map seedId: locationId
	// seedLocMap := make(map[int]int)
	minLocId := math.MaxInt
	for chunkCount, seedIdRange := range *seedIdRangesPtr {
		seedIds := parseSeedIdsFromRanges(&seedIdRange)
		// partSeedLocMap := translateSeeds(seedIds, translators)
		locId := minLocFromSeeds(seedIds, translators)
		// Check and store lowest location id value
		if locId < minLocId {
			minLocId = locId
		}
		fmt.Println("[INFO] Finished chunk: ", chunkCount, " / ", len(*seedIdRangesPtr))
	}
	return minLocId
}

func minLocFromSeeds(seedIds *[]int, translators *[]Translator) int {
	// iterate through the list of seed ids and map seedId: locationId
	minLocId := math.MaxInt
	for _, seedId := range *seedIds {
		locId := translateSeed(seedId, translators)
		// we got the locId - let's store the min value
		if locId < minLocId {
			minLocId = locId
		}
	}

	return minLocId

}
