package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/BennettB123/advent-of-code-2025/days"
)

var NumDays int = 1

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2 arguments expected. Usage:\n" +
			"  ./advent-of-code-2025 <day#> <inputFile>")
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Unable to parse day to int. Expected an integer but got: '%s'\n", os.Args[1])
		os.Exit(1)
	}

	fileContents, err := os.ReadFile(os.Args[2])
	if err != nil {
		fmt.Printf("Unable to read input file. Error: %s\n", err)
		os.Exit(1)
	}

	switch day {
	case 1:
		days.Day1(fileContents)
	case 2:
		days.Day2(fileContents)
	case 3:
		days.Day3(fileContents)
	case 4:
		days.Day4(fileContents)
	case 5:
		days.Day5(fileContents)
	case 6:
		days.Day6(fileContents)
	case 7:
		days.Day7(fileContents)
	case 8:
		days.Day8(fileContents)
	case 9:
		days.Day9(fileContents)
	default:
		fmt.Printf("Invalid day. Only %d days are complete", NumDays)
		os.Exit(1)
	}
}
