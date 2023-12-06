package day_five

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

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
	minLocId := naiveLowestLocation(input)
	fmt.Println("Lowest Location Id for initial seeds: ", minLocId)
}

// Entry point for day 5 part 2 solution
func solvePartTwo(input *[]string) {
	fmt.Println("--- Solving Day Five - Part Two! ---")
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
