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
General forumala: (base)^(matches - 1). There is a special case where match count < 1,
which should return 0.
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
