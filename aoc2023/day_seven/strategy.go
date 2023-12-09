package day_seven

import (
	"fmt"
)

/*
Ranking Behavior Interface
This is provided to abstract ranking behavior for CardHand structs
*/
type RankingBehavior interface {
	Name() string
	Type(string) int
	MapStrengths() *map[rune]int

	// note - these are private and allow us to implement dry code for a switch in Type()
	isHighCard(*map[rune]int) bool
	isOnePair(*map[rune]int) bool
	isTwoPair(*map[rune]int) bool
	isThreeOfAKind(*map[rune]int) bool
	isFullHouse(*map[rune]int) bool
	isFourOfAKind(*map[rune]int) bool
	isFiveOfAKind(*map[rune]int) bool
}

/*
StdRankingBehavior implements RankingBehavior
This is a standard ranking behavior that assumes 'J' labels (cards) are jacks and
doesn't support wildcards
*/
type StdRankingBehavior struct {
	name             string
	highLowOrderding string `default:"AKQJT98765432"`
}

// Return the Name of RankingBehavior
func (r *StdRankingBehavior) Name() string {
	return r.name
}

// Determines the enum hand type. This is used for ranking purposes
func (r *StdRankingBehavior) Type(cards string) int {
	// map card frequency in hand
	// cardFreq := mapCardFreq(cards)
	var rb RankingBehavior
	rb = r
	handType := determineType(cards, &rb)

	return handType
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
	return pairCount == 1 && trioCount == 1
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

/*
WildCardRankingBehavior implements RankingBehavior
This is an alternate ranking behavior that assumes 'J' labels (cards) are jokers and
supports wildcards. To balance this, jokers are assigned strength < '2' cards.
*/
type WildCardRankingBehavior struct {
	name             string
	highLowOrderding string `default:"AKQT98765432J"`
	// TODO: Add joker as a rune to make which card is wild config and not hardcoded
	// setting wild='' would result in the same behavior as StdRankingBehavior!
}

// Return the Name of RankingBehavior
func (r *WildCardRankingBehavior) Name() string {
	return r.name
}

// Determines the enum hand type. This is used for ranking purposes
func (r *WildCardRankingBehavior) Type(cards string) int {
	// delegate logic to helper function for dry code
	var rb RankingBehavior
	rb = r
	handType := determineType(cards, &rb)

	return handType
}

// Takes a label (card) ordering string from high to low and return a map for comparisons
func (r *WildCardRankingBehavior) MapStrengths() *map[rune]int {
	// TODO: This can* be refactored behind a helper function - but given loc, low value
	// build the map - and return it
	ordering := r.highLowOrderding
	strengthMap := mapCardRuneStrength(ordering)

	return strengthMap
}

// Simple eval function to determine hand type: high card. This is where all cards'
// runes are distinct: 23456. Wildcards are considered with raw freqency
func (r *WildCardRankingBehavior) isHighCard(freqMap *map[rune]int) bool {
	// Let's check for any wildcards - if wildcard found, at worst it's a pair
	wildCardFound := false
	for card, _ := range *freqMap {
		wildCardFound = wildCardFound || card == 'J'
	}
	return len(*freqMap) == 5 && !wildCardFound
}

// Determines hand type: one pair. This is where two cards share one rune, and the
// other three cards have a different rune from the pair and each other: A23A4
// Wildcards are also considered. Ex A23J4 is ALSO a pair
func (r *WildCardRankingBehavior) isOnePair(freqMap *map[rune]int) bool {
	wildCardCount := 0
	for card, freq := range *freqMap {
		if card == 'J' {
			wildCardCount = freq
		}
	}
	// 4 keys and no wildcard or 5 keys and 1 wildcard
	return (len(*freqMap) == 4 && wildCardCount == 0) || (len(*freqMap) == 5 && wildCardCount == 1)
}

// Determines hand type: two pair. This is where two cards share one label(rune), two
// other cards share a second label, and the remaining card has a third label: 23432
func (r *WildCardRankingBehavior) isTwoPair(freqMap *map[rune]int) bool {
	pairCount := 0
	wildCardCount := 0
	for card, freq := range *freqMap {
		if freq == 2 {
			pairCount++
		}
		if card == 'J' {
			wildCardCount = freq
		}
	}
	// 3 keys, and 2 pairs and no wild cards
	// note: if any wildcards are present and 1 pair is also present: 3 of a kind
	return len(*freqMap) == 3 && pairCount == 2 && wildCardCount == 0
}

// Determines whether a hand type: three of a kind. If exactly 3 cards have the same rune,
// and the other cards (2) are both different, return true
func (r *WildCardRankingBehavior) isThreeOfAKind(freqMap *map[rune]int) bool {
	trioCount := 0
	wildCardCount := 0
	for card, freq := range *freqMap {
		if freq == 3 {
			trioCount++
		}
		if card == 'J' {
			wildCardCount = freq
		}
	}
	// 4 keys and 1 wildcard (the pair becomes a 3 of a kind)
	// OR 3 keys and trio == 1 and no wildcards
	condMet := (len(*freqMap) == 4 && wildCardCount == 1) ||
		(len(*freqMap) == 3 && trioCount == 1 && wildCardCount == 0)
	return condMet
}

// Determines hand type: full house. This is where three cards have the same
// label(rune), and the remaining two cards share a different label: 23332
func (r *WildCardRankingBehavior) isFullHouse(freqMap *map[rune]int) bool {
	pairCount := 0
	trioCount := 0
	wildCardCount := 0
	for card, freq := range *freqMap {
		if freq == 3 {
			trioCount++
		} else if freq == 2 {
			pairCount++
		}
		if card == 'J' {
			wildCardCount = freq
		}
	}
	// 1 pair and 1 trio and no wildcardOR 2 pair and 1 wildcard
	// Note: 1 pair and 2 wildcards will evaluate up to 4 of a kind and
	// 1 pair or trio + 2 or 3 wildcards will turn into five of a kind
	return (pairCount == 1 && trioCount == 1 && wildCardCount == 0) || (pairCount == 2 && wildCardCount == 1)
}

// Determines hand type: four of a kind. This is where four cards have the same
// label(rune) and one card has a different label: AA8AA
func (r *WildCardRankingBehavior) isFourOfAKind(freqMap *map[rune]int) bool {
	quadCount := 0
	trioCount := 0
	pairCount := 0
	wildCardCount := 0
	for card, freq := range *freqMap {
		// check multiples
		switch freq {
		case 4:
			quadCount++
		case 3:
			trioCount++
		case 2:
			pairCount++
		}
		// check wildcards
		if card == 'J' {
			wildCardCount = freq
		}
		/*
			if freq == 4 {
				quadCount++
			} else if freq == 3 {
				trioCount++
			} else if freq == 2 {
				pairCount++
			}
			if card == 'J' {
				wildCardCount++
			}
		*/
	}
	// quad and 0 wildcards OR trio and 1 wildcard OR 1 pair (+ the pair of wildcards) and 2 wildcards will evaluate up to 4 of a kind
	condMet := (quadCount == 1 && wildCardCount == 0) || (trioCount == 1 && wildCardCount == 1) || (pairCount == 2 && wildCardCount == 2)
	return condMet
}

// Determines hand type: five of a kind. This is where all five cards have the same
// label(rune): AAAAA
func (r *WildCardRankingBehavior) isFiveOfAKind(freqMap *map[rune]int) bool {
	// look for any wildcards
	wildCardCount := 0
	for card, freq := range *freqMap {
		if card == 'J' {
			wildCardCount = freq
		}
	}
	// 1 key OR 2 keys and > 1 wildcards
	// Note: any more than 2 keys will downgrade to another type
	return (len(*freqMap) == 1) || (len(*freqMap) == 2 || wildCardCount > 0)
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

func determineType(cards string, rankBehavior *RankingBehavior) int {
	// map card frequency in hand
	cardFreq := mapCardFreq(cards)

	// eval hand types
	// note: would be interesting if this could be refactored with reflection - overkill?
	rb := *rankBehavior
	if rb.isHighCard(cardFreq) {
		return HIGH_CARD
	}
	if rb.isOnePair(cardFreq) {
		return ONE_PAIR
	}
	if rb.isTwoPair(cardFreq) {
		return TWO_PAIR
	}
	if rb.isThreeOfAKind(cardFreq) {
		return THREE_OF_A_KIND
	}
	if rb.isFullHouse(cardFreq) {
		return FULL_HOUSE
	}
	if rb.isFourOfAKind(cardFreq) {
		return FOUR_OF_A_KIND
	}
	if rb.isFiveOfAKind(cardFreq) {
		return FIVE_OF_A_KIND
	}
	// didn't find a match - print error to console
	fmt.Println("[ERROR] ", rb.Name(), " Couldn't identify type: ", cards)
	return -1
}
