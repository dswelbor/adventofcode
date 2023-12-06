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
WinBehavior concretion. This behavior implements a Geometric sequence algorithm
*/
type PointsWinBehavior struct {
	// power  int
	base   int
	points int
}

func (w *PointsWinBehavior) Win(matchCount int) int {
	points := math.Pow(float64(w.base), float64(matchCount-1))
	// w.power += 1
	w.points = int(points)
	return int(points)
}
