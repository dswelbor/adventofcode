package day_three

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/dswelbor/adventofcode/aoc2023/utility"
)

func SolveDayThree(input *[]string, part int) {

	if part == 1 {
		solvePartOne(input)
	} else if part == 2 {
		solvePartTwo(input)
	} else {
		fmt.Println("Part: " + strconv.Itoa(part) + "Not supported")
	}

}

func solvePartOne(input *[]string) {
	fmt.Println("--- Solving Day Three - Part One! ---")
	// doTheStuff(input)

	// testList := make([]string, 10)
	// testList[5] = "testing"
	// testResult := testList[2]
	// fmt.Println("[DEBUG] slice element not set test: ", testResult)

	// Map "symbols"
	symbolPattern := "[\\@\\*\\&\\%\\#\\/\\+\\=\\$\\!\\^\\(\\)\\-\\_]"
	symbols := mapSymbols(input, symbolPattern)
	// init validator - aware of symbol map
	validator := StdPartValidator{symbols: symbols}
	// Get list of Parts
	partPattern := "\\d+"
	parts := listParts(input, partPattern, &validator)

	// Grab part numbers and sum
	partNumbers := listPartNumbers(parts)
	partsSum := utility.SumNumbers(partNumbers)
	fmt.Println("Valid Part #'s Sum: ", partsSum)

	// fmt.Println(symbols)
	// fmt.Println(parts)
}

func solvePartTwo(input *[]string) {
	fmt.Println("--- Solving Day Three - Part Two! ---")
	// TODO: Implement me

}

func findAdjacentCoords(rowIndex int, colIndices []int) *[][]int {
	// get left coords and right
	// adjCoords := make([]string, 0)
	adjCoords := make([][]int, 0)
	// leftCoordStr := rowColumnString(rowIndex, colIndices[0]-1)
	//rightCoordStr := rowColumnString(rowIndex, colIndices[1])
	leftCoord := []int{rowIndex, colIndices[0] - 1}
	rightCoord := []int{rowIndex, colIndices[1]}

	// add l/r coords to list
	adjCoords = append(adjCoords, leftCoord)
	adjCoords = append(adjCoords, rightCoord)

	// Iterate through col indices - get top and bottom coords
	//var topCoordStr string
	//var botCoordStr string
	// var topCoord []int
	// var botCoord []int
	for col := colIndices[0] - 1; col <= colIndices[1]; col++ {
		// calc top and bottom coords
		// topCoordStr = rowColumnString(rowIndex-1, col)
		// botCoordStr = rowColumnString(rowIndex+1, col)
		topCoord := []int{rowIndex - 1, col}
		botCoord := []int{rowIndex + 1, col}

		// Add coords to list
		adjCoords = append(adjCoords, topCoord)
		adjCoords = append(adjCoords, botCoord)
	}
	return &adjCoords
}

func listParts(input *[]string, regPattern string, validator *StdPartValidator) *[]Part {
	// init regex and part collection
	reg := regexp.MustCompile(regPattern)
	parts := make([]Part, 0)

	// Iterate through input rows and grab part numbers with regex
	for row, inputStr := range *input {
		partIndices := reg.FindAllStringIndex(inputStr, -1) // grab part # reg matches
		// iterate through matches, create Part objects, and add to collection
		for _, indices := range partIndices {
			adjCoords := findAdjacentCoords(row, indices)
			partNum, err := strconv.Atoi(inputStr[indices[0]:indices[1]])
			if err != nil {
				panic(err)
			}
			part := Part{
				number:    partNum,
				adjCoords: adjCoords,
				validator: validator,
			}
			parts = append(parts, part)
		}
	}
	return &parts
}

func mapSymbols(input *[]string, regPattern string) *SymbolCollection {
	// init symbol collection
	rowCount := len(*input)
	colCount := len((*input)[0])
	symbolCollection := CreateSymbolCollection(rowCount, colCount)
	symbolCoords := *symbolCollection.symbolCoords
	//fmt.Println(symbolCollection)

	// init regexp obj
	reg := regexp.MustCompile(regPattern)

	// Iterate through rows to map symbol coors
	for row, inputStr := range *input {
		symbolIndices := reg.FindAllStringIndex(inputStr, -1)
		for _, indices := range symbolIndices {
			// fmt.Print(indices)
			// symbolStr := inputStr
			col := indices[0]
			symbol := inputStr[indices[0]:indices[1]]
			// symbolCollection.symbolCoords[row][col] = symbol
			symbolCoords[row][col] = symbol
		}
		fmt.Println(row)
		fmt.Println(symbolIndices)

	}
	return symbolCollection
}

// takes a part list and symbol map, filters valid parts based on adjacent coord checks,
// and returns
func listPartNumbers(parts *[]Part) *[]int {
	// get valid parts
	validParts := listValidParts(parts)
	// build list of part numbers
	partNumbers := make([]int, len(*validParts))
	for i, part := range *validParts {
		partNumbers[i] = part.number
	}

	return &partNumbers
}

// return a filtered list of "valid" parts
func listValidParts(parts *[]Part) *[]Part {
	// init collection of valid parts
	validParts := make([]Part, 0)
	for _, part := range *parts {
		if part.Valid() {
			validParts = append(validParts, part)
		}
	}

	return &validParts
}

// simple function to transform row and col to "row,col" string
func rowColumnString(row int, col int) string {
	coordStr := strconv.Itoa(row) + "," + strconv.Itoa(col)
	return coordStr
}

func doTheStuff(input *[]string) {
	// inputStr := "467..114....1.."
	// partReg := regexp.MustCompile("\\d+")
	// partMatches := partReg.FindAllStringIndex(inputStr, -1)
	// fmt.Println(partMatches)
	fmt.Println("Count of Rows: ", len(*input))
	fmt.Println("Count of Columns: ", len((*input)[0]))

	// Grab "symbol chars"
	runeMap := make(map[rune]bool)
	for _, inputStr := range *input {
		// Iterate through each row in input
		for _, r := range inputStr {
			runeMap[r] = true
		}
	}
	fmt.Println(runeMap)

	// print runes
	for r := range runeMap {
		if r != '.' {
			fmt.Printf("%c", r)
		}
	}
}
