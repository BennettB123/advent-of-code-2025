package days

import (
	"fmt"
	"strings"
)

func Day3(input []byte) {
	day3_star1(input)
	day3_star2(input)
}

func day3_star1(input []byte) {
	// 2 pass strategy
	//   1st pass finds the largest number
	//   2nd pass finds the largest number that follows the first
	accumulator := 0

	for line := range strings.Lines(string(input)) {
		line = strings.TrimSpace(line)

		// pass 1
		largestNum := 0
		largestNumIdx := 0

		// first number cannot be the final digit in the line
		for idx, digit := range line[:len(line)-1] {
			num := int(digit - '0')
			if num > largestNum {
				largestNum = num
				largestNumIdx = idx
			}
		}

		// pass 2
		nextLargestNum := 0
		for _, digit := range line[largestNumIdx+1:] {
			num := int(digit - '0')
			if num > nextLargestNum {
				nextLargestNum = num
			}
		}

		joltage := (largestNum * 10) + nextLargestNum
		accumulator += joltage
	}

	fmt.Printf("Star 1 answer: %d \n", accumulator)
}

func day3_star2(input []byte) {
	fmt.Printf("Star 2 answer: todo \n")
}
