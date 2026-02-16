package days

import (
	"fmt"
	"math"
	"strings"
)

func Day3(input []byte) {
	day3_star1(input)
	day3_star2(input)
}

func day3_star1(input []byte) {
	// 2 pass strategy
	//   1st pass finds the largest number (can't be the final digit)
	//   2nd pass finds the largest number that follows the first
	accumulator := 0

	for line := range strings.Lines(string(input)) {
		line = strings.TrimSpace(line)

		// pass 1
		largestNum := 0
		largestNumIdx := 0

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
	// 12 pass strategy
	//   1st pass finds the largest number (can't be within 12 digits from the end)
	//   2nd pass finds the largest number that follows the first
	//   3nd pass finds the largest number that follows the second
	//   ...
	var accumulator int64 = 0
	numDigitsNeeded := 12

	for line := range strings.Lines(string(input)) {
		line = strings.TrimSpace(line)

		currIdx := -1
		var joltage int64 = 0

		for i := range numDigitsNeeded {
			largestNum := 0

			for idx, digit := range line {
				lastIdx := len(line) - numDigitsNeeded + i
				if idx <= currIdx || idx > lastIdx {
					continue
				}

				num := int(digit - '0')
				if num > largestNum {
					largestNum = num
					currIdx = idx
				}
			}

			joltage += int64(largestNum * int(math.Pow(10, float64(numDigitsNeeded-i-1))))
		}

		accumulator += joltage
	}

	fmt.Printf("Star 2 answer: %d \n", accumulator)
}
