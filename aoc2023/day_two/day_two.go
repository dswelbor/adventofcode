package day_two

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dswelbor/adventofcode/aoc2023/utility"
)

type rgbValidator struct {
	// gameRoundPtr *map[string]string
	blueMax, greenMax, redMax int
}

// Utility method to set the validator's internal gameRound
// func (v rgbValidator) setGameRound(infoPtr *map[string]string) {
// 	v.gameRoundPtr = infoPtr
// }

func (v rgbValidator) Validate(infoPtr *map[string]string) bool {
	var valid bool
	// dereference info map
	info := *infoPtr

	// Edge Case - Check keys
	validKeys := []string{"red", "green", "blue"}
	foundInvalid := false
	for key := range info {
		foundInvalid = foundInvalid || !utility.ListContainsString(&validKeys, key)
	}
	valid = !foundInvalid

	// parse string to int - errors fail validation
	redCount, rErr := strconv.Atoi(info["red"])
	greenCount, gErr := strconv.Atoi(info["green"])
	blueCount, bErr := strconv.Atoi(info["blue"])
	if rErr != nil || gErr != nil || bErr != nil {
		// one of the string -> int conversions failed - invalid
		valid = false
	}

	valid = valid && redCount <= v.redMax && greenCount <= v.greenMax && blueCount <= v.blueMax

	return valid
}

type rgbDiceRound struct {
	red, green, blue int
}

// convenience function to return dice round info map
func (r rgbDiceRound) Info() *map[string]string {
	infoMap := map[string]string{
		"red":   strconv.Itoa(r.red),
		"green": strconv.Itoa(r.green),
		"blue":  strconv.Itoa(r.blue),
	}
	return &infoMap

}

type rgbDiceGame struct {
	gameId    string
	validator *utility.Validator
	gamePower *utility.PowerBehavior
	games     *[]utility.GameRound
}

func (g rgbDiceGame) Id() string {
	return g.gameId
}

// return id and game info as string=>string map
func (g rgbDiceGame) Info() *map[string]string {
	infoMap := map[string]string{
		"id":    g.Id(),
		"valid": strconv.FormatBool(g.valid()),
	}
	return &infoMap
}

// Iteratively validate the list of rgbDiceRounds using the flyweight validator
func (g rgbDiceGame) valid() bool {
	valid := true
	for _, round := range *g.games {
		// Grab rgbDiceRound info map
		roundInfoPtr := round.Info()

		// Set game round info and validate
		gameValidator := *g.validator // derefence validator flyweight
		valid = valid && gameValidator.Validate(roundInfoPtr)
	}
	return valid
}

type minValidPowerBehavior struct {
	colors []string
}

func (p minValidPowerBehavior) Power(gameRounds *[]utility.GameRound) int {
	// TODO: Implement me
	return 1
}

/*
Calculates a power for a "game"
*/
func (g rgbDiceGame) power() int {
	// TODO: Implement me
	return 1
}

func SolveDayTwo(input *[]string, part int) {

	if part == 1 {
		solvePartOne(input)
	} else if part == 2 {
		solvePartTwo(input)
	} else {
		fmt.Println("Part: " + strconv.Itoa(part) + "Not supported")
	}

}

func solvePartOne(input *[]string) {
	fmt.Println("--- Solving Day Two - Part One! ---")
	// TODO: Implement me
	for _, inputStr := range *input {
		gameId := parseGameId(inputStr)
		fmt.Println(gameId)
	}

	// Build validator object
	var validator utility.Validator
	validator = rgbValidator{
		redMax:   12,
		greenMax: 13,
		blueMax:  14,
	}

	// Build a list of Games
	gamesPtr := parseGames(input, &validator)

	// Iterate through parsed games and filter gameIds for valid games
	validGameIds := make([]int, 0)
	for _, gamePtr := range *gamesPtr {
		gameId, err := strconv.Atoi(gamePtr.Id())
		if err != nil {
			panic(err)
		}
		gameStatus := *gamePtr.Info()
		valid, err := strconv.ParseBool(gameStatus["valid"])
		if err == nil && valid {
			fmt.Println("[DEBUG]: GameID: " + strconv.Itoa(gameId) + " is valid")
			validGameIds = append(validGameIds, gameId)
		}

	}

	// Calc game id sum
	idSum := utility.SumNumbers(&validGameIds)

	fmt.Println("Valid Game IDs Sum: ", strconv.Itoa(idSum))
	// fmt.Println("[DEBUG]: ", gamesPtr)
}

func solvePartTwo(input *[]string) {
	fmt.Println("--- Solving Day Two - Part Two! ---")
	// TODO: Implement me
}

/*
Iterate through game metadata input strings and build a list of Game objects
that implement the utility.Game interface
*/
func parseGames(input *[]string, validator *utility.Validator) *[]utility.Game {
	// Parse a list of Games from input strings
	games := make([]utility.Game, 0)
	for _, inputStr := range *input {
		// grab id, rounds, and validator - init Game
		gameId := parseGameId(inputStr)
		gameRoundsPtr := parseRounds(inputStr)
		game := rgbDiceGame{
			gameId:    gameId,
			games:     gameRoundsPtr,
			validator: validator,
		}
		games = append(games, game)
	}

	return &games
}

func parseGameId(inputStr string) string {
	// Parse out Game \d*: from input str
	fmt.Println("[DEBUG]: ", inputStr)
	idReg := regexp.MustCompile("Game \\d*:")
	gameIdStr := idReg.FindString(inputStr)
	fmt.Println("[DEBUG]: ", gameIdStr)

	// Parse digit from Game \d*: string
	digitReg := regexp.MustCompile("\\d+")
	idStr := digitReg.FindString(gameIdStr)
	fmt.Println("[DEBUG]: ", idStr)

	return idStr // ex. '14' from 'Game 14:...' string
}

func parseRounds(inputStr string) *[]utility.GameRound {
	roundStrings := strings.Split(inputStr, "; ")

	// parse rgbDiceRound objs from color and count from each round
	// rounds := make([]rgbDiceRound, 0)
	rounds := make([]utility.GameRound, 0)
	for _, roundStr := range roundStrings {
		// get rgbDiceRound from game round string - add to list
		roundPtr := parseRound(roundStr)
		rounds = append(rounds, *roundPtr)

	}

	// list of rgbDiceRounds pased - return
	return &rounds
}

func parseRound(roundStr string) *utility.GameRound {
	// Grab a list of rounds in a given game
	colorCountPattern := "\\d+ (green|red|blue)"
	colorCountReg := regexp.MustCompile(colorCountPattern)
	colorCountStrings := colorCountReg.FindAllString(roundStr, -1)

	// iterate trough color count string - build rgbDiceRound obj
	// var round utility.GameRound
	//round := new(rgbDiceRound)
	roundColorMap := map[string]int{"red": 0, "green": 0, "blue": 0}
	for _, colorCountStr := range colorCountStrings {
		// grab qty and color details
		qtyColorTuple := strings.Split(colorCountStr, " ")
		qty, _ := strconv.Atoi(qtyColorTuple[0])
		color := qtyColorTuple[1]
		switch color {
		case "red":
			roundColorMap["red"] = qty
		case "green":
			roundColorMap["green"] = qty
		case "blue":
			roundColorMap["blue"] = qty
		default:
			panic("Error parsing color and qty")
		}

	}

	// Create new GameRound object from roundColorMap
	var round utility.GameRound
	round = rgbDiceRound{
		red:   roundColorMap["red"],
		green: roundColorMap["green"],
		blue:  roundColorMap["blue"],
	}

	// rgbDiceRound populated - return it
	return &round
}
