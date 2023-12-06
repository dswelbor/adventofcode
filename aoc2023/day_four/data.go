package day_four

import (
	"fmt"
	"strconv"
)

/*
Data structure to store the values for a given game card including:
- gane card id
- game card winning numbers (that the numbers the card has have to match)
- game card numbers (these have to match winning numbers to get a score)
- a WinBehavior object - abstracts how scoring is calculated
*/
type GameCard struct {
	cardId      int
	winMap      *map[int]bool
	numStrings  *[]string
	winBehavior *WinBehavior
}

// simple accessor function that returns the GameCard id
func (c *GameCard) Id() int {
	return c.cardId
}

// Determines the score for a given GameCard. This leverages the WinBehavior
// the GameCard object has and relies on a count of how many numbers the GameCard has
// that match the GameCard's winning numbers
func (c *GameCard) Win() int {
	matchCount := c.matchCount()
	// use WinBehavior to calculate score
	winBehavior := *c.winBehavior
	score := winBehavior.Win(matchCount)

	return score
}

// Helper method to caclulate the count for GameCard numbers that match winning numbers
func (c *GameCard) matchCount() int {
	// init match could and dereference winning nunber map
	matchCount := 0
	winMap := *c.winMap

	// iterate through GameCard numbers - increment count on matches to winning numbers
	for _, numStr := range *c.numStrings {
		// init match flag
		match := false
		// parse number string into int
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("[ERROR] Error parsing: \"", numStr, "\" error: ", err)
		} else {
			// no error - normal case
			match = winMap[num]
			if match {
				// match found! let's increment the match counter
				matchCount += 1
			}
		}
	}

	return matchCount
}

/*
Collection of GameCards
*/
type GameCardDeck struct {
	cards *[]GameCard
}

// Fetches the GameCard from the GameCardDeck collection. Returns nil if not found
func (d *GameCardDeck) Get(id int) *GameCard {
	gameCards := *d.cards
	// handle bad input
	if id < 0 || id > len(gameCards) {
		return nil
	}
	// fetch and return
	gameCard := gameCards[id-1]
	return &gameCard
}
