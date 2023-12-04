package day_three

// import "fmt"

type Part struct {
	number    int
	adjCoords *[][]int
	validator *StdPartValidator
}

func (p Part) Valid() bool {
	return p.validator.Valid(&p)
}

type SymbolCollection struct {
	symbolCoords *[][]string
}

// Constructor creates Symbol Collection with known row and col sizes
func CreateSymbolCollection(rowCount int, colCount int) *SymbolCollection {
	// Create collection of rows
	symbolCoords := make([][]string, rowCount)
	// Iteratively create each row in collection of rows
	for row := 0; row < rowCount; row++ {
		newRow := make([]string, colCount)
		symbolCoords[row] = newRow
	}
	// init collection
	symbolCollection := SymbolCollection{symbolCoords: &symbolCoords}

	return &symbolCollection
}

type StdPartValidator struct {
	symbols *SymbolCollection
}

func (v StdPartValidator) Valid(part *Part) bool {
	// iterate through part adjc coords
	symbolMap := *v.symbols.symbolCoords

	// set row and col max indices
	rowMax := len(symbolMap)
	colMax := 0
	if rowMax > 0 {
		colMax = len(symbolMap[0])
	}

	for _, coord := range *part.adjCoords {
		// handle index out of bounds errors
		if coord[0] >= 0 && coord[0] < rowMax && coord[1] >= 0 && coord[1] < colMax {
			symbol := symbolMap[coord[0]][coord[1]]
			if len(symbol) > 0 {
				// a symbol was found in the adj coords list
				return true
			}
		}
	}

	// symbol was matched in part adj coords
	return false
}
