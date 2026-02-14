package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day1(input []byte) {
	day1_star1(input)
	day1_star2(input)
}

func day1_star1(input []byte) {
	lines := strings.Lines(string(input))

	numTimesZero := 0
	dialPos := 50

	// fmt.Printf("Dial starting at %d \n", dialPos)

	for line := range lines {
		line = strings.TrimSpace(line)

		isLeft := line[0] == 'L'
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("invalid input. Expected integer but got '%s'\n", line[1:])
			os.Exit(1)
		}

		if isLeft {
			dialPos -= value
			for dialPos < 0 {
				dialPos += 100
			}
		} else {
			dialPos += value
			for dialPos > 99 {
				dialPos -= 100
			}
		}

		// fmt.Printf("Rotating %s = %d \n", line, dialPos)

		if dialPos == 0 {
			numTimesZero++
		}
	}

	fmt.Printf("Star 1 answer: %d \n", numTimesZero)
}

func day1_star2(input []byte) {
	lines := strings.Lines(string(input))

	numTimesZero := 0
	dialPos := 50

	fmt.Printf("Dial starting at %d \n", dialPos)

	for line := range lines {
		line = strings.TrimSpace(line)

		isLeft := line[0] == 'L'
		clicks, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("invalid input. Expected integer but got '%s'\n", line[1:])
			os.Exit(1)
		}

		// ugly brute force because I'm tired of trying to fix edge cases
		for range clicks {
			if isLeft {
				dialPos -= 1
			} else {
				dialPos += 1
			}

			if dialPos < 0 {
				dialPos += 100
			} else if dialPos > 99 {
				dialPos -= 100
			}

			if dialPos == 0 {
				numTimesZero++
			}
		}

	}

	fmt.Printf("Star 2 answer: %d \n", numTimesZero)
}
