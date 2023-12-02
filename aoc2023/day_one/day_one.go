package day_one

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
)

func SolveDayOne(input *[]string, part int) {

	if part == 1 {
		solvePartOne(input)
	} else if part == 2 {
		solvePartTwo(input)
	} else {
		fmt.Println("Part: " + strconv.Itoa(part) + "Not supported")
	}

}

func solvePartOne(input *[]string) {
	fmt.Println("--- Solving Day One - Part One! ---")
	// Get all the calibration partOneNumbers
	partOneRegex := "\\d"
	partOneNumbers := calibrationNumbers(input, partOneRegex, nil)
	// fmt.Println(*numbers)

	// Sum all the calibration numbers
	partOneSum := sumNumbers(partOneNumbers)

	// Print the sum
	resultStr := "Calibration Number Sum: " + strconv.Itoa(partOneSum)
	fmt.Println(resultStr)
}

func solvePartTwo(input *[]string) {
	fmt.Println("--- Solving Day One - Part Two! ---")
	// Create map of int and string representations of numbers
	digitsMapPtr := createWordIntMap()

	// Get all the calibration partOneNumbers
	// partTwoRegex := "((?=(\\d))|(?=(one))|(?=(two))|(?=(three))|(?=(four))|(?=(five))|(?=(six))|(?=(seven))|(?=(eight))|(?=(nine)))"
	partTwoRegex := "(\\d|one|two|three|four|five|six|seven|eight|nine)"
	partTwoNumbers := calibrationNumbers(input, partTwoRegex, digitsMapPtr)

	// Sum all the calibration numbers
	partTwoSum := sumNumbers(partTwoNumbers)

	// Print the sum
	resultStr := "Calibration Number Sum: " + strconv.Itoa(partTwoSum)
	fmt.Println(resultStr)
}

// takes a pointer to a slice of strings, parses callibration numbers, and returns a pointer to a list of callibration numbers
func calibrationNumbers(input *[]string, regexPattern string, digitsMap *map[string]string) *[]int {
	// iterate through each line and add parsed calibrationNumber to slice
	numbers := make([]int, 0)

	for _, strElement := range *input {
		num := calibrationNumber(strElement, regexPattern, digitsMap)
		numbers = append(numbers, num)
	}

	// return slice of parsed calibration numbers
	return &numbers

}

// Parse a calibration number from an input string line
func calibrationNumber(inputStr string, regexPattern string, digitsMap *map[string]string) int {
	// Grab all overlapping matches
	// Note: regexp.FindAllString(string) does not support overlapping matches
	// This is important since abconeightxyz should return matches: ["one", "eight"] with a shared 'e'
	matchesPtr := findOverlappingStrings(inputStr, regexPattern)
	// handle number "words" and not just digits
	if digitsMap != nil {
		matchesPtr = parseNumbers(matchesPtr, digitsMap)
	}

	// dereference matches
	matches := *matchesPtr
	// grab first and last numbers
	matchesLen := len(matches)
	firstNum := matches[0]
	lastNum := matches[matchesLen-1]

	// combine the calibration number elements
	combinedNum, err := strconv.Atoi(firstNum + lastNum)
	if err != nil {
		panic(err)
	}

	// return the combined calibration number
	return combinedNum
}

func regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}

/*
*

	Implemented a find all overlapping strings function since
	regexp.FindAllString(string) returns all non overlapping matches. regexp2
	provides a right to left search, as well as lookahead capabilities.
	However, this function allows us to find overlapping matches using stdlib
	regexp functionality
*/
func findOverlappingStrings(inputStr string, regexPattern string) *[]string {
	// testing regexp lib behavior
	// testStr := "eightbpsqrkzhqbhjlrxmzsixvvmgtrseventwo7oneightjbx"
	// testRegex := "(\\d|one|two|three|four|five|six|seven|eight|nine)"

	// build regexp obj
	reg := regexp.MustCompile(regexPattern)

	var matches []string
	matchIndex := reg.FindStringIndex(inputStr)
	for len(matchIndex) != 0 {
		i, j := matchIndex[0], matchIndex[1]
		match := inputStr[i:j]
		matches = append(matches, match)

		// trim the testStr to get overlapping matches
		inputStr = inputStr[i+1 : len(inputStr)]
		matchIndex = reg.FindStringIndex(inputStr)
	}
	return &matches
}

func reverseString(inputString string) string {
	// Convert string to runes
	strRunes := []rune(inputString)
	// iterate through each "rune" in string, and reverse
	var reverseString strings.Builder
	strLen := len(inputString)
	for i := (strLen - 1); i >= 0; i-- {
		c := strRunes[i]
		reverseString.WriteRune(c)
	}

	// for _, c := range inputString {
	// 	reverseString.WriteRune(c)
	// }
	reversedString := reverseString.String()
	return reversedString
}

// Add up all the calibration numbers and return their sum
func sumNumbers(numbers *[]int) int {
	sum := 0
	for _, num := range *numbers {
		sum += num
	}
	return sum
}

func parseNumbers(numbers *[]string, digitsMapPtr *map[string]string) *[]string {
	// Fix
	// Derefence digitsMap
	digitsMap := *digitsMapPtr

	parsedNumbers := make([]string, 0)
	for _, num := range *numbers {
		parsedNum := digitsMap[num]
		parsedNumbers = append(parsedNumbers, parsedNum)
	}
	return &parsedNumbers
}

func createWordIntMap() *map[string]string {
	// initialize number word to digit map
	digitsMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	// add numerical digit to digit mapp
	for i := 1; i < 10; i++ {
		num := strconv.Itoa(i)
		digitsMap[num] = num
	}
	return &digitsMap
}
