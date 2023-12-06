package day_four

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
func SolveDayFour(input *[]string, part int) {

	if part == 1 {
		solvePartOne(input)
	} else if part == 2 {
		solvePartTwo(input)
	} else {
		fmt.Println("Part: " + strconv.Itoa(part) + "Not supported")
	}

}

// Entry point for day 4 part 1 solution
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

// Entry point for day 4 part 2 solution
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
