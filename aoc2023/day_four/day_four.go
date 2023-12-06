package day_four

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type MatchVisitor struct {
	power  int
	base   int
	points int
}

func (v *MatchVisitor) Visit(match bool) int {
	points := math.Pow(float64(v.base), float64(v.power))
	if match {
		v.power += 1
		v.points = int(points)
	}
	return int(points)
}

func (v *MatchVisitor) Points() int {
	return v.points
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

	// Build builder
	var deckBuilder DeckBuilder
	deckBuilder = &DeckBuilderConcrete{winBehaviorType: "points"}
	// Call Construct (usually abstracted in a Director interface)
	deck := ConstructGameCardDeck(deckBuilder, input)

	// Iterate through cards in collection and get points
	points := 0
	for _, gameCard := range *deck.cards {
		points += gameCard.Win()
	}
	fmt.Println("Sum of win points: ", points)
}

func solvePartTwo(input *[]string) {
	fmt.Println("--- Solving Day Four - Part Two! ---")

	// Build builder
	var deckBuilder DeckBuilder
	deckBuilder = &DeckBuilderConcrete{winBehaviorType: "cards"}
	// Call Construct (usually abstracted in a Director interface)
	deck := ConstructGameCardDeck(deckBuilder, input)

	// Iterate through cards in collection and get points
	points := 0
	for _, gameCard := range *deck.cards {
		points += gameCard.Win()
	}
	fmt.Println("Sum of win points: ", points)
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
		// points := 0
		/*
			visitor := MatchVisitor{
				base:  2,
				power: 0,
			}
		*/
		var pointsBehavior WinBehavior
		pointsBehavior = &PointsWinBehavior{base: 2}
		matchCount := 0

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
					// visitor.Visit(match)
					matchCount += 1
				}
			}
		}
		// points = visitor.Points()
		points := pointsBehavior.Win(matchCount)
		pointsList = append(pointsList, points)

	}
	return &pointsList
}

// Utility function acts as the Director with a Construct. In more complex
// constructions, this should be abstracted as a class (or an interface in Go)
// However, since this level of complexity is unneeded, in go fashion we are implemnting
// this component of the Design pattern as a utility function
func ConstructGameCardDeck(builder DeckBuilder, input *[]string) *GameCardDeck {
	// Iterate through input and build a card
	// builder := *builderPtr
	for _, inputStr := range *input {
		builder.BuildCard(inputStr)
	}
	// Finished calling BuildCard() from input - get build GameCardDeck
	return builder.GetCollection()
}
