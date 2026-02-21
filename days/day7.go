package days

import (
	"fmt"
	"slices"
	"strings"
)

const (
	Start    = 'S'
	Empty    = '.'
	Splitter = '^'
)

func Day7(input []byte) {
	lines := slices.Collect(strings.Lines(string(input)))
	for i := range len(lines) {
		lines[i] = strings.TrimSpace(lines[i])
	}

	day7_star1(lines)
	day7_star2(lines)
}

func day7_star1(lines []string) {
	beams := make(map[int]struct{}, 0)
	numSplits := 0

	// starting point
	beams[strings.Index(lines[0], string(Start))] = struct{}{}

	for _, line := range lines {
		for beamX := range beams {
			if line[beamX] == Splitter {
				delete(beams, beamX)
				beams[beamX-1] = struct{}{}
				beams[beamX+1] = struct{}{}
				numSplits++
			}
		}
	}

	fmt.Printf("Star 1 answer: %d \n", numSplits)
}

type coordinate struct {
	x int
	y int
}

var coordinateCache map[coordinate]int64 = make(map[coordinate]int64)

func day7_star2(lines []string) {
	startCol := strings.Index(lines[0], string(Start))
	numSplits := getNumTimelines(lines, 0, startCol)

	fmt.Printf("Star 2 answer: %d \n", numSplits)
}

// memoized recursive function to calculate the number of timelines that result from
// a particle at a given location
func getNumTimelines(lines []string, row, col int) int64 {
	currCoord := coordinate{x: col, y: row}
	if timelines, ok := coordinateCache[currCoord]; ok {
		return timelines
	}

	// base case: reached bottom of grid, so that represents 1 complete timeline
	if row >= len(lines)-1 {
		return 1
	}

	// If on an empty spot or starting point, just move down
	if lines[row][col] == Start || lines[row][col] == Empty {
		return getNumTimelines(lines, row+1, col)
	}

	// must be a splitter, so get the number of timelines from both resulting particles
	numSplits := getNumTimelines(lines, row+1, col-1) + getNumTimelines(lines, row+1, col+1)
	coordinateCache[currCoord] = numSplits
	return numSplits
}
