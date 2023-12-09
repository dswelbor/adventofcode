package day_eight

import (
	"fmt"
	"regexp"
	"strconv"
)

type DesertMap struct {
	steps *[]int               // 0 is left, 1 is right
	edges *map[string][]string // node edges
}

// High level entry Point for Day 8 solution
func SolveDayEight(input *[]string, part int) {

	if part == 1 {
		solvePartOne(input)
	} else if part == 2 {
		solvePartTwo(input)
	} else {
		fmt.Println("Part: " + strconv.Itoa(part) + "Not supported")
	}

}

// Entry point for day 8 part 1 solution
func solvePartOne(input *[]string) {
	fmt.Println("--- Solving Day Eight - Part One! ---")

	// Let's grab int indexed steps and nodeId -> nodeId mapped edges
	desertMap := parseDesertMap(input)
	fmt.Println("[DEBUG] desertMap steps length: ", len(*desertMap.steps))
	fmt.Println("[DEBUG] desertMap edges length: ", len(*desertMap.edges))
}

// Entry point for day 8 part 2 solution
func solvePartTwo(input *[]string) {
	fmt.Println("--- Solving Day Eight - Part Two! ---")
}

func parseDesertMap(input *[]string) *DesertMap {
	// initialize edges for desert map
	edgesMap := make(map[string][]string)
	// init int index encoded steps for desert map
	stepsLen := len((*input)[0])
	steps := make([]int, stepsLen)
	// init desertMap
	desertMap := DesertMap{
		edges: &edgesMap,
		steps: &steps,
	}
	// Initialize regex for parsing
	nodeReg := regexp.MustCompile("\\w+")
	// iterate through input lines. i=0 is the L/R steps order. i > 1 are map "edges"
	for i, inputStr := range *input {
		if i == 0 {
			// First line of input is the L/R step list - build list of ints where L=0, and R=1
			stepRunes := []rune(inputStr)
			// steps := make([]int, len(stepRunes))
			lrMap := map[rune]int{'L': 0, 'R': 1}
			for j, step := range stepRunes {
				// Get L=0 or R=1 for step direction - ie L/R edge index
				stepDir := lrMap[step]
				// steps = append(steps, stepDir)
				steps[j] = stepDir
			}
			// We've got the L=0, R=1 step list - re-assign to desert map
			// Commented out - shouldn't be needed
			// desertMap.steps = &steps  // this might* not be needed - but
		} else if i > 1 {
			// i=1 is an empty line, we want to process mapped edges for lines i>=2
			nodeIds := nodeReg.FindAllString(inputStr, -1)

			// regex match returns triple: [nodeId, leftNodeId, rightNodeId]
			// these are nodeId -> [leftNodeId, rightNodeId] edges
			nodeId := nodeIds[0]
			leftNodeId := nodeIds[1]
			rightNodeId := nodeIds[2]
			edgeNodeIds := []string{leftNodeId, rightNodeId} // instantiate L/R edges
			// update node/edge map
			edges := *desertMap.edges
			edges[nodeId] = edgeNodeIds
		}
	}
	// DesertMap struct build with direction indexed steps and mapped nodeId edges
	return &desertMap
}

func navigateDesertMap(steps *[]rune) int {
	// TODO: Implement me
	return 0
}
