package day_four

import (
	"math"
)

/*
A common WinBehavior interface for behaviors (or algorithms) when a GameCard
has a win(s). This follows the strategy design pattern.
*/
type WinBehavior interface {
	Win(matchCount int) int
}

/*
WinBehavior concretion. This object implements a Geometric sequence algorithm
*/
type PointsWinBehavior struct {
	// power  int
	base   int
	points int
}

/*
Implements WinBehavior Win() interface. This is a geometic sequence like {1, 2, 4, 8,...}
with first term a=1 and common ratio r=base. General forumala: (base)^(matches - 1).
There is a special case where match count < 1 (this is outside the domain of the
sequence indices but a valid matchCount parameter), which should return 0.
*/
func (w *PointsWinBehavior) Win(matchCount int) int {
	// edge case where match count = 0
	if matchCount < 1 {
		return 0
	}
	// geometric sequence for points = (a)(r)^(n-1) where start term a is 1
	points := math.Pow(float64(w.base), float64(matchCount-1))
	w.points = int(points)
	return int(points)
}

/*
WinBehavior concretion. This data structure supports scoring based on number of cards
win. This implementation relies on awareness of a GameCardDeck collection (to select
cards relative to current card id)
*/
type CardCopyWinBehavior struct {
	gameCardId int
	cardDeck   *GameCardDeck
}

/*
Calculates score as num of cards won for this card, and recursively, the number of cards
won from those won card copies. Cards won are selected in an offset range from current
game card id.
*/
func (w *CardCopyWinBehavior) Win(matchCount int) int {
	// Grab won cards from deck
	wonCards := make([]GameCard, 0)
	for offset := 1; offset <= matchCount; offset++ {
		// fetch won card and add it slice of cards
		offsetId := w.gameCardId + offset
		wonCard := w.cardDeck.Get(offsetId)
		wonCards = append(wonCards, *wonCard)
	}
	curScore := len(wonCards) // we won 0 or more cards for this card - include in score

	// Iterate through won cards and recursively get scores for won cards
	recursiveScore := 0
	for _, wonCard := range wonCards {
		recursiveScore += wonCard.Win()
	}

	// Add score for this card and recusively aggregated score
	score := curScore + recursiveScore
	return score
}
