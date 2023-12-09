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
	Cards        string
	Bid          int
	RankBehavior *RankingBehavior
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
		theseCards := []rune(h.Cards)
		otherCards := []rune(otherHand.Cards)
		// cardStrengths := mapCardRuneStrength()
		// var rankBehavior RankingBehavior
		rankBehavior := *h.RankBehavior
		cardStrengths := rankBehavior.MapStrengths()
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
	// cardRunes := []rune(h.Cards)
	// cardFreq := make(map[rune]int)
	// for _, c := range cardRunes {
	// 	cardFreq[c] += 1
	// }
	// Debugging
	if h.Cards == "KTJJT" {
		fmt.Println("Let's pause here")
	}
	rankBehavior := *h.RankBehavior
	handType := rankBehavior.Type(h.Cards)

	return handType

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
	// init Ranking Behavior (std with no jokers)
	var stdRankingBehaviorPtr RankingBehavior
	stdRankingBehaviorPtr = &StdRankingBehavior{
		name:             "Standard Ranking Behavior",
		highLowOrderding: "AKQJT98765432",
	}

	// Grab a collection of card hands
	cardHands := parseCardHands(input, &stdRankingBehaviorPtr)

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

	// Notes: d7 pt2
	// 254456791 is too high
	// init Ranking Behavior (jokers are wildcards)
	var wildRankingBehaviorPtr RankingBehavior
	wildRankingBehaviorPtr = &WildCardRankingBehavior{
		name:             "Joker is Wild Ranking Behavior",
		highLowOrderding: "AKQT98765432J",
	}

	// Grab a collection of card hands
	cardHands := parseCardHands(input, &wildRankingBehaviorPtr)

	// Now let's order in rank (rank=index+1) by CardHand strength
	sort.Slice(*cardHands, func(i, j int) bool {
		return (*cardHands)[i].Less(&(*cardHands)[j])
	})

	// we have a sorted list of card hands - let's sum total winnings
	totalWinnings := calcTotalWinnings(cardHands)
	fmt.Println("Total Winnings: ", totalWinnings)
}

// Takes a list of lines from input and creates a collection of CardHands
func parseCardHands(input *[]string, rankBehavior *RankingBehavior) *[]CardHand {
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
				Cards:        cards,
				Bid:          bid,
				RankBehavior: rankBehavior,
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
		bid := cardHand.Bid
		// the winning for each hand is bid * rank
		winnings := bid * rank
		totalWinnings += winnings
	}
	return totalWinnings
}
