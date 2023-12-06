package utility

import (
	"fmt"
	"strconv"
)

type Validator interface {
	Validate(*map[string]string) bool
}

type Game interface {
	Id() string
	Info() *map[string]string
}

type GameRound interface {
	Info() *map[string]string
}

type PowerBehavior interface {
	Power(*[]GameRound) int
}

/*
Utility function to implement contains for a list of strings.
*/
func ListContainsString(items *[]string, searchTerm string) bool {
	// iterate through item in list
	for _, item := range *items {
		if item == searchTerm {
			return true
		}
	}
	return false
}

func ListAtoi(numStrings *[]string) *[]int {
	numbers := make([]int, len(*numStrings)) // init list of ints

	for i, numStr := range *numStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("[ERROR] Problem parsing number string: ", numStr,
				" into an int")
			panic(err) // TODO: Implement a better way to handle this
		}
		// add number to list
		numbers[i] = num
	}
	return &numbers
}
