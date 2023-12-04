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
	// Map "symbols"
	symbolPattern := "[\\@\\*\\&\\%\\#\\/\\+\\=\\$\\!\\^\\(\\)\\-\\_]"
	symbols := mapSymbols(input, symbolPattern)
	// init validator - aware of symbol map
	// validator := StdPartValidator{symbols: symbols}
	// Get list of Parts
	partPattern := "\\d+"
	parts := listParts(input, partPattern, nil)

	// build reverse "gear" map - and build gears list
	revGearParts := buildReverseGearMap(parts, symbols)
	gears := listValidGears(revGearParts)

	// grab gear ratios and sum ratios
	gearRatios := listGearRatios(gears)
	gearRatioSum := utility.SumNumbers(gearRatios)

	// Print sum of gear ratios
	fmt.Println("Sum of Gear Ratios: ", gearRatioSum)
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

// Goes through a builds a
func buildReverseGearMap(parts *[]Part, symbols *SymbolCollection) *[][][]Part {
	// init reverse map of symbols that match criteria for gear "*" symbol and adjacent parts
	revGearPartsPtr := initReverseGearMap(symbols)
	revGearParts := *revGearPartsPtr
	// iterate through parts list
	for _, part := range *parts {
		adjCoords := part.adjCoords
		// Iterate through adjacent coords
		for _, coord := range *adjCoords {
			// try to fetch adjacent symbol
			row := coord[0]
			col := coord[1]
			symbol := symbols.Symbol(row, col)
			// check if Part is adj to gear "*" symbol
			if symbol == "*" {
				// current part is adjacent to gear "*" symbol
				// add to list of parts for that row, col
				coordParts := revGearParts[row][col]
				coordParts = append(coordParts, part)
				revGearParts[row][col] = coordParts
				//revGearParts[row][col]
			}
		}

	}

	// we have a (row, col): []Part reverse map
	return revGearPartsPtr
}

// utility function that initializes a (row, col): part list collection. This can be
// used to essentially reverse map symbols to a collection of adjacent parts
func initReverseGearMap(symbols *SymbolCollection) *[][][]Part {
	rowCount := len((*symbols.symbolCoords))
	colCount := 0
	if rowCount > 0 {
		colCount = len((*symbols.symbolCoords)[0])
	}
	revGearParts := make([][][]Part, rowCount) // init list of rows
	for row := 0; row < rowCount; row++ {
		newRow := make([][]Part, colCount)
		revGearParts[row] = newRow
		// we assume that each element is by default an empty []Part
	}

	return &revGearParts
}

func listValidGears(revGearParts *[][][]Part) *[]GearPart {
	// init list of "gears"
	gears := make([]GearPart, 0)

	// iterate each row
	for rowIndex, row := range *revGearParts {
		// iterate each column in row
		for colIndex, partList := range row {
			// check count of adjacent parts
			if len(partList) == 2 {
				// current symbol has exactly 2 adjacent parts - is a gear
				parts := make([]Part, 2)
				for i, part := range partList {
					parts[i] = part
				}
				gear := GearPart{
					parts:  &parts,
					row:    rowIndex,
					col:    colIndex,
					symbol: "*",
				}
				gears = append(gears, gear)
			}
		}
	}

	// finished grabbing gears
	return &gears
}

func listGearRatios(gears *[]GearPart) *[]int {
	// init list of gear rations
	gearRatios := make([]int, 0)

	// iterate through gears, add ratio() value to list
	for _, gear := range *gears {
		ratio := gear.Ratio()
		gearRatios = append(gearRatios, ratio)
	}

	return &gearRatios
}
