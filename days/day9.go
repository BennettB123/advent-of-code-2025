package days

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day9(input []byte) {
	lines := slices.Collect(strings.Lines(string(input)))
	for i := range len(lines) {
		lines[i] = strings.TrimSpace(lines[i])
	}

	day9_star1(lines)
	day9_star2(lines)
}

func day9_star1(lines []string) {
	redTiles := make([]coordinate, 0)
	for _, line := range lines {
		split := strings.Split(line, ",")
		if len(split) != 2 {
			fmt.Printf("Expected split to contain 2 elements, but it contained %d", len(split))
			os.Exit(1)
		}

		x, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Printf("Unable to convert '%s' to int. Error: %v", split[0], err)
			os.Exit(1)
		}

		y, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Printf("Unable to convert '%s' to int. Error: %v", split[1], err)
			os.Exit(1)
		}

		redTiles = append(redTiles, coordinate{x, y})
	}

	maxArea := int64(0)
	for i := range redTiles {
		for j := range redTiles {
			if i == j {
				continue
			}

			xDist := int64(redTiles[i].x-redTiles[j].x) + 1
			yDist := int64(redTiles[i].y-redTiles[j].y) + 1
			area := int64(math.Abs(float64(xDist * yDist)))

			if area > maxArea {
				maxArea = area
			}
		}
	}
	fmt.Printf("Star 1 answer: %d \n", maxArea)
}

func day9_star2(lines []string) {
	fmt.Printf("Star 2 answer: %d \n", 0)
}
