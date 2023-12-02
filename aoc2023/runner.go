package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dswelbor/adventofcode/aoc2023/day_one"
)

func main() {
	// Grab cli args and parse
	dayPtr := flag.Int("day", 1, "problem day number")
	filepathPtr := flag.String("file", "data/day_one_part_one_ex.txt", "relative filtepath to input")
	partPtr := flag.Int("part", 1, "problem part: ex. -part=1 or -part=2")
	flag.Parse()

	// print cli args

	fmt.Println("day:", *dayPtr)
	fmt.Println("file:", *filepathPtr)
	fmt.Println("part: ", *partPtr)
	inputPtr := readInputFile(filepathPtr)

	// Execute day (and part) by passed cli args
	switch *dayPtr {
	case 1:
		day_one.SolveDayOne(inputPtr, *partPtr)
	default:
		fmt.Println("Day: " + strconv.Itoa(*dayPtr) + " not implemented")
	}
}

// Pass in a pointer to file path, read the file by line, and return slice of strings
func readInputFile(filepathPtr *string) *[]string {
	fileStrings := make([]string, 0)

	f, err := os.Open(*filepathPtr)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		// fmt.Println(scanner.Text())
		fileStrings = append(fileStrings, scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &fileStrings
}
