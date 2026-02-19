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
	day7_star1(input)
	day7_star2(input)
}

func day7_star1(input []byte) {
	// x index of all current beams
	beams := make(map[int]struct{}, 0)
	numSplits := 0

	lines := slices.Collect(strings.Lines(string(input)))
	beams[strings.Index(lines[0], string(Start))] = struct{}{}

	for _, line := range lines[1:] {
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

func day7_star2(input []byte) {
	fmt.Printf("Star 2 answer: %d \n", 0)
}
