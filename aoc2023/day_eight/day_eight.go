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

	// Let's count the steps from traversing the mapped graph
	totalSteps := navigateDesertMap(desertMap, "AAA", "ZZZ")
	fmt.Println("[COMPLETE] Reached destination \"ZZZ\"! Steps taken: ", totalSteps)
}

// Entry point for day 8 part 2 solution
func solvePartTwo(input *[]string) {
	fmt.Println("--- Solving Day Eight - Part Two! ---")

	// Let's grab int indexed steps and nodeId -> nodeId mapped edges
	desertMap := parseDesertMap(input)
	fmt.Println("[DEBUG] desertMap steps length: ", len(*desertMap.steps))
	fmt.Println("[DEBUG] desertMap edges length: ", len(*desertMap.edges))

	// NOTE: Part 2 is conceptually a BFS? Maybe?
	totalSteps := navDesertMapParallel(desertMap, "\\w\\wA", "\\w\\wZ")

	// Let's print our results
	fmt.Println("[COMPLETE] Finished traversal with total steps: ", totalSteps)
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

func navigateDesertMap(desertMap *DesertMap, startNodeId string, endNodeId string) int {
	// Grab edges, steps, and init nodeId
	steps := *desertMap.steps
	edges := *desertMap.edges
	nodeId := startNodeId
	// Iterate through steps - emulate a circular linked list with modulo
	stepCount := 0

	for found := false; !found; stepCount++ {
		// test
		// found = true

		// Check if we are done
		if nodeId == endNodeId {
			found = true
		}

		// get mod from step counter - steps loop around
		// grab step direction index
		stepIndex := stepCount % len(steps)
		stepDir := steps[stepIndex]

		// Traverse to next node
		nodeEdges := edges[nodeId]
		nextNodeId := nodeEdges[stepDir]
		fmt.Println("[INFO] Traversing from ", nodeId, " to ", nextNodeId,
			"\t\tstep count: ", stepCount+1)
		// set nodeId to next nodeId
		nodeId = nextNodeId

	}

	// traversed the graph to destination - return counted steps
	// account for stepCount++ always incrementing - even on the last iteration where we are found
	return stepCount - 1
}

func navDesertMapParallel(desertMap *DesertMap, startPattern string, endPattern string) int {
	// build reg objects
	// startPattern = "\\w\\wA"
	// endPattern = "\\w\\wZ"
	startReg := regexp.MustCompile(startPattern)
	endReg := regexp.MustCompile(endPattern)

	// Grab edges, steps, and init nodeId
	// steps := *desertMap.steps
	edges := *desertMap.edges
	// Find list of starting nodeIds and endNodeIds
	startNodeIds := make([]string, 0)
	endNodeIds := make([]string, 0)
	for nodeId, _ := range edges {
		startMatch := startReg.FindString(nodeId)
		endMatch := endReg.FindString(nodeId)
		if len(startMatch) > 0 {
			// found a starting node
			startNodeIds = append(startNodeIds, startMatch)
		} else if len(endMatch) > 0 {
			// found an end node
			endNodeIds = append(endNodeIds, endMatch)
		}
	}

	fmt.Println("[DEBUG] Starting Node IDs: ", startNodeIds)
	fmt.Println("[DEBUG] Ending NodeIDs: ", endNodeIds)

	// init nodeId and steps
	steps := *desertMap.steps
	nodeIds := startNodeIds
	// Iterate through steps - emulate a circular linked list with modulo
	stepCount := 0
	for foundAll := false; !foundAll; stepCount++ {
		// grab step direction index
		stepIndex := stepCount % len(steps)
		stepDir := steps[stepIndex]
		lrMap := []string{"L", "R"}
		stepDirStr := lrMap[stepDir]
		// init list for next traverse iteration
		nextNodeIds := make([]string, len(nodeIds))

		// debugging - store found ids to check length
		foundEndIds := make([]string, 0)

		// Check if we are done
		for i, nodeId := range nodeIds {
			// foundEndIds := make([]int, 0)
			// nextNodeId := nodeEdges[stepDir] // useful for debugging
			found := endReg.FindString(nodeId)
			if len(found) > 0 {
				// fmt.Println("[Debug] Found End Node: ", found, "\tstep direction: ", stepDirStr, "\tstep count: ", stepCount)
				foundEndIds = append(foundEndIds, found)
			}
			foundAll = foundAll && len(found) > 0

			// Traverse to next node
			nodeEdges := edges[nodeId]
			nextNodeId := nodeEdges[stepDir]
			// fmt.Println("[INFO] Traversing from ", nodeId, " to ", nextNodeId,
			// 	"\t\tstep count: ", stepCount+1)
			// set nodeId to next nodeId
			// nodeId = nextNodeId
			// add nextNodeId to nextNodeIds list
			nextNodeIds[i] = nextNodeId
		}

		// Set stopping condition
		if len(foundEndIds) == len(nodeIds) {
			// all nodes in parallel ordered traversal meet conditions
			foundAll = true
		}

		// debugging - how many end nodes are found?
		if len(foundEndIds) > 1 {
			// we have more than one end id
			fmt.Println("[Debug] Found ", len(foundEndIds), " Nodes: ", foundEndIds, "\tstep direction: ", stepDirStr, "\tstep count: ", stepCount)
		}
		// if nodeId == endNodeId {
		// 	found = true
		//}

		/*
			// Traverse to next node
			nodeEdges := edges[nodeId]
			nextNodeId := nodeEdges[stepDir]
			fmt.Println("[INFO] Traversing from ", nodeId, " to ", nextNodeId,
				"\t\tstep count: ", stepCount+1)
			// set nodeId to next nodeId
			nodeId = nextNodeId
		*/
		nodeIds = nextNodeIds
	}
	return stepCount - 1
}
