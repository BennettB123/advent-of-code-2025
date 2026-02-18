package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day6(input []byte) {
	day6_star1(input)
	day6_star2(input)
}

func day6_star1(input []byte) {
	numbers := make([][]int, 0)
	operands := make([]string, 0)

	parsingOperands := false

	for line := range strings.Lines(string(input)) {
		// check if we're parsing numbers or operands
		firstChar := strings.TrimSpace(line)[0]
		if firstChar == '*' || firstChar == '+' {
			parsingOperands = true
		}

		if !parsingOperands {
			newNumbers := make([]int, 0)
			split := strings.Split(line, " ")
			for _, numberStr := range split {
				numberStr = strings.TrimSpace(numberStr)
				if numberStr == "" {
					continue
				}

				number, err := strconv.Atoi(numberStr)
				if err != nil {
					fmt.Printf("Unable convert string '%s' to int. Error %s", numberStr, err)
					os.Exit(1)
				}

				newNumbers = append(newNumbers, number)
			}

			numbers = append(numbers, newNumbers)
		} else {
			split := strings.Split(line, " ")
			for _, operand := range split {
				operand = strings.TrimSpace(operand)
				if operand == "" {
					continue
				}

				operands = append(operands, operand)
			}
		}
	}

	// ensure all number lists and the operand list are the same numColumns
	numColumns := len(numbers[0])
	for _, numList := range numbers {
		if len(numList) != numColumns {
			fmt.Println("Number lists are not all the same size!")
			os.Exit(1)
		}
	}
	if len(operands) != numColumns {
		fmt.Println("Operand list is not the same size as the number lists!")
		os.Exit(1)
	}

	accumulator := 0

	for col := range numColumns {
		if operands[col] == "*" {
			result := 1
			for row := range numbers {
				result *= numbers[row][col]
			}
			accumulator += result
		} else if operands[col] == "+" {
			result := 0
			for row := range numbers {
				result += numbers[row][col]
			}
			accumulator += result
		} else {
			fmt.Printf("Invalid operand found: '%s' \n", operands[col])
			os.Exit(1)
		}
	}

	fmt.Printf("Star 1 answer: %d \n", accumulator)
}

func day6_star2(input []byte) {
	fmt.Printf("Star 2 answer: %d \n", 0)
}
