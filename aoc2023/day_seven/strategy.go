package day_seven

import "fmt"

type RankingBehavior interface {
	Type(cards string) int
	MapStrengths() *map[rune]int
}

type StdRankingBehavior struct {
	name             string
	highLowOrderding string `default:"AKQJT98765432"`
}

func (r *StdRankingBehavior) Type(cards string) int {
	// map card frequency in hand
	cardFreq := mapCardFreq(cards)

	// eval hand types
	if r.isHighCard(cardFreq) {
		return HIGH_CARD
	}
	if r.isOnePair(cardFreq) {
		return ONE_PAIR
	}
	if r.isTwoPair(cardFreq) {
		return TWO_PAIR
	}
	if r.isThreeOfAKind(cardFreq) {
		return THREE_OF_A_KIND
	}
	if r.isFullHouse(cardFreq) {
		return FULL_HOUSE
	}
	if r.isFourOfAKind(cardFreq) {
		return FOUR_OF_A_KIND
	}
	if r.isFiveOfAKind(cardFreq) {
		return FIVE_OF_A_KIND
	}
	// didn't find a match - print error to console
	fmt.Println("[ERROR] ", r.name, " Couldn't identify type: ", cards)
	return -1
}

// Takes a label (card) ordering string from high to low and return a map for comparisons
func (r *StdRankingBehavior) MapStrengths() *map[rune]int {
	// build the map - and return it
	ordering := r.highLowOrderding
	strengthMap := mapCardRuneStrength(ordering)

	return strengthMap
}

// Simple eval function to determine hand type: high card. This is where all cards'
// runes are distinct: 23456
func (r *StdRankingBehavior) isHighCard(freqMap *map[rune]int) bool {
	return len(*freqMap) == 5
}

// Determines hand type: one pair. This is where two cards share one rune, and the
// other three cards have a different rune from the pair and each other: A23A4
func (r *StdRankingBehavior) isOnePair(freqMap *map[rune]int) bool {
	return len(*freqMap) == 4
}

// Determines hand type: two pair. This is where two cards share one label(rune), two
// other cards share a second label, and the remaining card has a third label: 23432
func (r *StdRankingBehavior) isTwoPair(freqMap *map[rune]int) bool {
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
func (r *StdRankingBehavior) isThreeOfAKind(freqMap *map[rune]int) bool {
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
func (r *StdRankingBehavior) isFullHouse(freqMap *map[rune]int) bool {
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
func (r *StdRankingBehavior) isFourOfAKind(freqMap *map[rune]int) bool {
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
func (r *StdRankingBehavior) isFiveOfAKind(freqMap *map[rune]int) bool {
	return len(*freqMap) == 1
}

// Builds a map of card label (rune) strength. This is so that we can evaluate
// A > K = true and T > J = false and T > 9 = true
func mapCardRuneStrength(highLowCards string) *map[rune]int {
	// Set card labels (runes) from highest to lowest
	// highLowCards := "AKQJT98765432"
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

// takes a "hand" of cards (aka labels) as a string and maps card frequency
func mapCardFreq(cards string) *map[rune]int {
	// map card frequency in hand
	cardRunes := []rune(cards)
	cardFreq := make(map[rune]int)
	for _, c := range cardRunes {
		cardFreq[c] += 1
	}
	return &cardFreq
}
