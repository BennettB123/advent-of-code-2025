package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day1_main(input []byte) {
	starOne(input)
	starTwo(input)
}

func starOne(input []byte) {
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

func starTwo(input []byte) {
	fmt.Printf("Star 2 answer: todo \n")
}
