package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2(input []byte) {
	day2_star1(input)
	day2_star2(input)
}

type idRange struct {
	start int
	end   int
}

func day2_star1(input []byte) {
	idRanges := make([]idRange, 0, 256)

	for _, rangeStr := range strings.Split(string(input), ",") {
		rangeStr = strings.TrimSpace(rangeStr)
		split := strings.Split(rangeStr, "-")

		start, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Printf("Expected int but got '%s'", split[0])
			os.Exit(1)
		}

		end, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Printf("Expected int but got '%s'", split[1])
			os.Exit(1)
		}

		idRanges = append(idRanges, idRange{start, end})
	}

	accumulator := 0
	for _, r := range idRanges {
		for i := range (r.end - r.start) + 1 {
			id := r.start + i
			idStr := strconv.Itoa(id)

			// if length isn't even, can't be a repeat
			if len(idStr)%2 != 0 {
				continue
			}

			halfLen := len(idStr) / 2
			firstHalf := idStr[:halfLen]
			lastHalf := idStr[halfLen:]
			if firstHalf == lastHalf {
				accumulator += id
			}
		}
	}

	fmt.Printf("Star 1 answer: %d \n", accumulator)
}

func day2_star2(input []byte) {
	fmt.Printf("Star 2 answer: todo \n")
}
