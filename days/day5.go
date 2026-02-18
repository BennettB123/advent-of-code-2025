package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type idRange64 struct {
	start int64
	end   int64
}

func Day5(input []byte) {
	idRanges := make([]idRange64, 0)
	ids := make([]int64, 0)

	parsingRanges := true
	for line := range strings.Lines(string(input)) {
		line = strings.TrimSpace(line)
		if line == "" {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			split := strings.Split(line, "-")

			r := idRange64{}

			start, err := strconv.ParseInt(split[0], 10, 64)
			if err != nil {
				fmt.Printf("failed to parse start of range. Expected int64. Error: '%s' \n", err)
				os.Exit(1)
			}
			r.start = start

			end, err := strconv.ParseInt(split[1], 10, 64)
			if err != nil {
				fmt.Printf("failed to parse end of range. Expected int64. Error: '%s' \n", err)
				os.Exit(1)
			}
			r.end = end

			idRanges = append(idRanges, r)
		} else {
			id, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				fmt.Printf("failed to parse id. Expected int64 but got '%s' \n", err)
				os.Exit(1)
			}

			ids = append(ids, id)
		}
	}

	day5_star1(idRanges, ids)
	day5_star2(idRanges)
}

func day5_star1(idRanges []idRange64, idsToCheck []int64) {
	numFresh := 0

	for _, id := range idsToCheck {
		for _, idRange := range idRanges {
			if id >= idRange.start && id <= idRange.end {
				numFresh++
				break
			}
		}
	}

	fmt.Printf("Star 1 answer: %d \n", numFresh)
}

func day5_star2(idRanges []idRange64) {
	// Just for fun, let's see how slow brute force is with keeping track of all fresh IDs...
	// turns out it's impossibly slow...
	// A single idRange didn't complete in several minutes and it ate all my RAM.

	// freshIds := make(map[int64]bool, 0)
	// for rangeIdx, idRange := range idRanges {
	// 	fmt.Printf("on range %d \n", rangeIdx)
	// 	for i := idRange.start; i <= idRange.end; i++ {
	// 		if freshIds[i] != true {
	// 			numFresh++
	// 			freshIds[i] = true
	// 		}
	// 	}
	// }

	// Lets remove all overlaps from the set of ranges
	didReduce := true
	for didReduce {
		idRanges, didReduce = reduceRanges(idRanges)
	}

	numFresh := int64(0)
	for _, idRange := range idRanges {
		numFresh += (idRange.end - idRange.start) + 1
	}

	fmt.Printf("Star 2 answer: %d \n", numFresh)
}

// Does a single pass to reduce overlaps in the ranges.
// Does NOT remove all overlaps with a single call.
// Returns the new ranges and whether or not any reductions occurred.
func reduceRanges(idRanges []idRange64) ([]idRange64, bool) {
	newRanges := make([]idRange64, 0)
	didReduce := false

	for _, oldRange := range idRanges {
		doesOverlap := false

		for i, newRange := range newRanges {
			if oldRange.start <= newRange.end && newRange.start <= oldRange.end {
				doesOverlap = true
				newRanges[i].start = min(oldRange.start, newRange.start)
				newRanges[i].end = max(oldRange.end, newRange.end)
				didReduce = true
			}
		}

		if !doesOverlap {
			newRanges = append(newRanges, oldRange)
		}
	}

	return newRanges, didReduce
}
