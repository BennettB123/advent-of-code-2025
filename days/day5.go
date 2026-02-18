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
	day5_star2(idRanges, ids)
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

func day5_star2(idRanges []idRange64, idsToCheck []int64) {
	fmt.Printf("Star 2 answer: %d \n", 0)
}
