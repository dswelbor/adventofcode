package day_seven

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
Enum for various card hand types - such as four of a kind or two pair. These are
declared in weakest to strongest order so that they can be evaluated with comparison
operators. For instance, we would expect FIVE_OF_A_KIND > FOUR_OF_A_KIND.
*/
const (
	HIGH_CARD       = iota // HIGH_CARD = 0
	ONE_PAIR        = iota // ONE_CARD = 1
	TWO_PAIR        = iota
	THREE_OF_A_KIND = iota // THREE_OF_A_KIND = 3
	FULL_HOUSE      = iota
	FOUR_OF_A_KIND  = iota
	FIVE_OF_A_KIND  = iota // FIVE_OF_A_KIND = 6
)

type CardHand struct {
	cards string
	bid   int
}

func (h *CardHand) Less(otherHand *CardHand) bool {
	var isLess bool
	if h.HandType() < otherHand.HandType() {
		// usual rule, this hand has a lower hand type than the other hand
		isLess = true
	} else if h.HandType() > otherHand.HandType() {
		// usual rule, this hand has a high hand type than other hand, and less is false
		isLess = false
	} else {
		// both hands have the same type - need to evaluate card strength from first to
		// last in hand
		// Let's iterate in parallel and compare
		theseCards := []rune(h.cards)
		otherCards := []rune(otherHand.cards)
		cardStrengths := mapCardRuneStrength()
		for i := 0; i < len(theseCards); i++ {
			// compare current card
			thisCard := theseCards[i]
			otherCard := otherCards[i]
			thisCardStrength := (*cardStrengths)[thisCard]
			otherCardStrength := (*cardStrengths)[otherCard]
			if thisCardStrength < otherCardStrength {
				// we found a card that was lower - this hand is less - stop
				isLess = true
				break
			} else if thisCardStrength > otherCardStrength {
				// we found a card that was higher - this hand is not less - stop
				isLess = false
				break
			}
			// if we get here - the 2 cards from this hand and other hand were the same
			// keep iterating
		}

	}
	return isLess
}

func (h *CardHand) HandType() int {
	// map card frequency in hand
	cardRunes := []rune(h.cards)
	cardFreq := make(map[rune]int)
	for _, c := range cardRunes {
		cardFreq[c] += 1
	}

	// eval hand types
	if isHighCard(&cardFreq) {
		return HIGH_CARD
	}
	if isOnePair(&cardFreq) {
		return ONE_PAIR
	}
	if isTwoPair(&cardFreq) {
		return TWO_PAIR
	}
	if isThreeOfAKind(&cardFreq) {
		return THREE_OF_A_KIND
	}
	if isFullHouse(&cardFreq) {
		return FULL_HOUSE
	}
	if isFourOfAKind(&cardFreq) {
		return FOUR_OF_A_KIND
	}
	if isFiveOfAKind(&cardFreq) {
		return FIVE_OF_A_KIND
	}
	// didn't find a match - print error to console
	fmt.Println("[ERROR] Couldn't identify type: ", h.cards)
	return -1
}

// High level entry Point for Day 7 solution
func SolveDaySeven(input *[]string, part int) {

	if part == 1 {
		solvePartOne(input)
	} else if part == 2 {
		solvePartTwo(input)
	} else {
		fmt.Println("Part: " + strconv.Itoa(part) + "Not supported")
	}

}

// Entry point for day 7 part 1 solution
func solvePartOne(input *[]string) {
	fmt.Println("--- Solving Day Seven - Part One! ---")
	// Grab a collection of card hands
	cardHands := parseCardHands(input)

	// Now let's order in rank (rank=index+1) by CardHand strength
	sort.Slice(*cardHands, func(i, j int) bool {
		return (*cardHands)[i].Less(&(*cardHands)[j])
	})

	// we have a sorted list of card hands - let's sum total winnings
	totalWinnings := calcTotalWinnings(cardHands)
	fmt.Println("Total Winnings: ", totalWinnings)

}

// Entry point for day 7 part 2 solution
func solvePartTwo(input *[]string) {
	fmt.Println("--- Solving Day Seven - Part Two! ---")
}

// Takes a list of lines from input and creates a collection of CardHands
func parseCardHands(input *[]string) *[]CardHand {
	// let's iterate through the input and create CardHand structs
	cardHands := make([]CardHand, 0)
	for _, inputStr := range *input {
		// each line has hand and bid info in the form of: <<cards>> <<bid>>
		// ex: 32T3K 765
		// where cards="32T3K" and bid=765
		splitInput := strings.Split(inputStr, " ")
		cards := splitInput[0]
		bid, bidErr := strconv.Atoi(splitInput[1])
		if bidErr != nil {
			fmt.Println("[ERROR] Couldn't parse: \"", splitInput[1], "\"")
		} else {
			// we have hand and bid - lets create CardHand struct
			cardHand := CardHand{
				cards: cards,
				bid:   bid,
			}
			cardHands = append(cardHands, cardHand)
		}
	}
	// Yay! We have a collection of CardHands
	return &cardHands
}

func calcTotalWinnings(sortedCardHands *[]CardHand) int {
	// validate list of CardHands is sorted
	isSorted := sort.SliceIsSorted(*sortedCardHands, func(i, j int) bool {
		return (*sortedCardHands)[i].Less(&(*sortedCardHands)[j])
	})
	if !isSorted {
		// sort the list
		sort.Slice(*sortedCardHands, func(i, j int) bool {
			return (*sortedCardHands)[i].Less(&(*sortedCardHands)[j])
		})

	}

	// Let's iterate through sorted list and sum total winnings
	totalWinnings := 0
	for i, cardHand := range *sortedCardHands {
		rank := i + 1
		bid := cardHand.bid
		// the winning for each hand is bid * rank
		winnings := bid * rank
		totalWinnings += winnings
	}
	return totalWinnings
}

// Simple eval function to determine hand type: high card. This is where all cards'
// runes are distinct: 23456
func isHighCard(freqMap *map[rune]int) bool {
	return len(*freqMap) == 5
}

// Determines hand type: one pair. This is where two cards share one rune, and the
// other three cards have a different rune from the pair and each other: A23A4
func isOnePair(freqMap *map[rune]int) bool {
	return len(*freqMap) == 4
}

// Determines hand type: two pair. This is where two cards share one label(rune), two
// other cards share a second label, and the remaining card has a third label: 23432
func isTwoPair(freqMap *map[rune]int) bool {
	pairCount := 0
	for _, freq := range *freqMap {
		if freq == 2 {
			pairCount++
		}
	}
	return len(*freqMap) == 3 && pairCount == 2
}

// Determines whether a hand type: three of a kind. If exactly 3 cards have the same rune,
// and the other cards (2) are both different, return true
func isThreeOfAKind(freqMap *map[rune]int) bool {
	trioCount := 0
	for _, freq := range *freqMap {
		if freq == 3 {
			trioCount++
		}
	}
	return len(*freqMap) == 3 && trioCount == 1
}

// Determines hand type: full house. This is where three cards have the same
// label(rune), and the remaining two cards share a different label: 23332
func isFullHouse(freqMap *map[rune]int) bool {
	pairCount := 0
	trioCount := 0
	for _, freq := range *freqMap {
		if freq == 3 {
			trioCount++
		}
		if freq == 2 {
			pairCount++
		}
	}
	return len(*freqMap) == 2 && pairCount == 1 && trioCount == 1
}

// Determines hand type: four of a kind. This is where four cards have the same
// label(rune) and one card has a different label: AA8AA
func isFourOfAKind(freqMap *map[rune]int) bool {
	quadCount := 0
	for _, freq := range *freqMap {
		if freq == 4 {
			quadCount++
		}
	}
	return len(*freqMap) == 2 && quadCount == 1
}

// Determines hand type: five of a kind. This is where all five cards have the same
// label(rune): AAAAA
func isFiveOfAKind(freqMap *map[rune]int) bool {
	return len(*freqMap) == 1
}

// Builds a map of card label (rune) strength. This is so that we can evaluate
// A > K = true and T > J = false and T > 9 = true
func mapCardRuneStrength() *map[rune]int {
	// Set card labels (runes) from highest to lowest
	highLowCards := "AKQJT98765432"
	highLowRunes := []rune(highLowCards)
	// iteratively build the map
	strengthMap := make(map[rune]int)
	cardCount := len(highLowRunes)
	for offset := 0; offset < cardCount; offset++ {
		// grab card and calculate strength
		strength := cardCount - offset
		card := highLowRunes[offset]
		// map the card strength
		strengthMap[card] = strength
	}
	return &strengthMap
}
