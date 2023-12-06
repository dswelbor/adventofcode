package day_four

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
Interface for implementations of the Builder Design pattern to create a GameCardDeck
collection object.
*/
type DeckBuilder interface {
	BuildCard(string)
	GetCollection() *GameCardDeck
}

/*
Concrete data structure for the implementation of the DeckBuilder interface. Supports
building GameCardDeck with different WinBehavior objects - such as PointWinBehavior or
CardCopyWinBehavior,
*/
type DeckBuilderConcrete struct {
	winBehaviorType string
	deck            *GameCardDeck
}

// Takes an input string, creates WinBehavior, and creates GameCard with that behavior.
// iteratively adds the new object to the GameCardDeck that is being built
func (b *DeckBuilderConcrete) BuildCard(cardInputStr string) {
	// Handle initializing the deck
	if b.deck == nil {
		// initialize GameCardDeck an empty list of GameCards
		deckCards := make([]GameCard, 0)
		deck := &GameCardDeck{
			cards: &deckCards,
		}
		// set deck in builder
		b.deck = deck
	}

	// split Card # from numbers on ":"
	allNumbers := strings.Split(cardInputStr, ":")
	// Grab id
	numReg := regexp.MustCompile("\\d+")
	idStr := numReg.FindString(allNumbers[0])
	gameCardId, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("[ERROR] Problem encountered parsing gamecard id \"", idStr, "\"")
		fmt.Println("[ERROR] Error details: ", err)
	}

	// Grab scratched off and winning numbers
	// split winning #'s from scratched numbers on "|"
	numbers := strings.Split(allNumbers[1], "|")

	// Grab a list of winning numbers and map it
	winNumStrings := numReg.FindAllString(numbers[0], -1)
	winMap := mapWinningNumbers(&winNumStrings)
	// Grab scratched off numbers
	scratchedNumStrings := numReg.FindAllString(numbers[1], -1)

	// create WinBehavior for GameCard
	var winBehavior *WinBehavior
	// pointsBehavior = &PointsWinBehavior{base: 2}
	winBehavior = b.createWinBehavior(gameCardId)

	// Create GameCard and add to GameCardDeck being built
	gameCard := &GameCard{
		cardId:      gameCardId,
		winMap:      winMap,
		numStrings:  &scratchedNumStrings,
		winBehavior: winBehavior,
	}
	gameCards := *b.deck.cards
	gameCards = append(gameCards, *gameCard)
	b.deck.cards = &gameCards
}

// Returns the built GameCardDeck collection
func (b *DeckBuilderConcrete) GetCollection() *GameCardDeck {
	return b.deck
}

// Helper method that creates the WinBehavior. WinBehavior type is determined from the
// builders winBehaviorType attribute. Returns nil if type is unsupported
func (b *DeckBuilderConcrete) createWinBehavior(id int) *WinBehavior {
	var winBehavior WinBehavior

	// Pick WinBehavior from winBehaviorType attribute
	switch b.winBehaviorType {
	case "points":
		// this is a win behavior that scores based on a geometric sequence with first
		// term a=1 and common ratio (base) r=2
		winBehavior = &PointsWinBehavior{
			base: 2,
		}
	case "cards":
		// this win behavior calculates score as the sum of the number of cards won from
		// a single card and recursively the sum of the number of cards won from those
		// won cards
		winBehavior = &CardCopyWinBehavior{
			gameCardId: id,
			cardDeck:   b.deck,
		}
	default:
		fmt.Println("[ERROR] WinBehavior type: ", b.winBehaviorType, " not supported")
	}
	return &winBehavior
}

// Function takes a list of winning number strings and returns a simple string: true map
// This allows O(n) lookups to see if a scratched off number matches a winning number.
func mapWinningNumbers(winNumStrings *[]string) *map[int]bool {
	// iterate through list of number strings - map winning numbers to keys
	winMap := make(map[int]bool)
	for _, numStr := range *winNumStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("[ERROR] Error parsing: \"", numStr, "\" error: ", err)
		} else {
			// no error - normal case
			winMap[num] = true
		}
	}

	return &winMap
}
