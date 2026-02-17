package days

import (
	"fmt"
	"strings"
)

func Day4(input []byte) {
	day4_star1(input)
	day4_star2(input)
}

const PaperRoll = '@'

func day4_star1(input []byte) {
	grid := make([][]rune, 0)
	for line := range strings.Lines(string(input)) {
		line = strings.TrimSpace(line)
		row := make([]rune, 0)
		for _, element := range line {
			row = append(row, element)
		}

		grid = append(grid, row)
	}

	height := len(grid)
	width := len(grid[0])

	numRolls := 0

	for y := range height {
		for x := range width {
			if r, _ := getAt(grid, x, y); r != PaperRoll {
				continue
			}

			numNeighbors := 0
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if dx == 0 && dy == 0 {
						continue
					}

					if r, _ := getAt(grid, x+dx, y+dy); r == PaperRoll {
						numNeighbors++
					}
				}
			}

			if numNeighbors < 4 {
				numRolls++
			}
		}
	}

	fmt.Printf("Star 2 answer: %d \n", numRolls)
}

func day4_star2(input []byte) {
	fmt.Printf("Star 2 answer: todo \n")
}

func getAt(grid [][]rune, x, y int) (rune, error) {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
		return '0', fmt.Errorf("Out of bounds")
	}

	return grid[y][x], nil
}
