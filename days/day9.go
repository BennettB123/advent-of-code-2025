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
	day9_star1(input)
	day9_star2(input)
}

func parseInput(input []byte) []coordinate {
	lines := slices.Collect(strings.Lines(string(input)))
	for i := range len(lines) {
		lines[i] = strings.TrimSpace(lines[i])
	}

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

	return redTiles
}

func day9_star1(input []byte) {
	redTiles := parseInput(input)

	maxArea := int64(0)
	for i := range redTiles {
		for j := range redTiles {
			if i == j {
				continue
			}

			xDist := int64(math.Abs(float64(redTiles[i].x-redTiles[j].x))) + 1
			yDist := int64(math.Abs(float64(redTiles[i].y-redTiles[j].y))) + 1
			area := xDist * yDist

			if area > maxArea {
				maxArea = area
			}
		}
	}
	fmt.Printf("Star 1 answer: %d \n", maxArea)
}

type coordinatePair struct {
	c1 coordinate
	c2 coordinate
}

func day9_star2(input []byte) {
	redTiles := parseInput(input)

	xCoords := make([]int, 0)
	yCoords := make([]int, 0)
	for _, coord := range redTiles {
		if !slices.Contains(xCoords, coord.x) {
			xCoords = append(xCoords, coord.x)
		}
		if !slices.Contains(yCoords, coord.y) {
			yCoords = append(yCoords, coord.y)
		}
	}
	slices.Sort(xCoords)
	slices.Sort(yCoords)

	// compress heights
	compressedWidths := make(map[int]int, 0)
	for i := 0; i < len(xCoords)-1; i++ {
		compressedWidths[(i*2)+1] = xCoords[i+1] - xCoords[i] - 1
	}

	compressedHeights := make(map[int]int, 0)
	for i := 0; i < len(yCoords)-1; i++ {
		compressedHeights[(i*2)+1] = yCoords[i+1] - yCoords[i] - 1
	}

	// create compressed list of coordinate
	compressedRedTiles := make([]coordinate, 0)
	for _, tile := range redTiles {
		cx := slices.Index(xCoords, tile.x) * 2
		cy := slices.Index(yCoords, tile.y) * 2
		compressedRedTiles = append(compressedRedTiles, coordinate{cx, cy})
	}

	// initialize compressed grid to all zeros
	grid := make([][]int, 0)
	gridHeight := len(yCoords)*2 - 1
	gridWidth := len(xCoords)*2 - 1
	for range gridHeight {
		grid = append(grid, make([]int, gridWidth, gridWidth))
	}

	// add red/green tile edges to the grid
	for currIdx := range redTiles {
		nextIdx := currIdx + 1
		if currIdx == len(redTiles)-1 {
			nextIdx = 0
		}
		x1 := redTiles[currIdx].x
		y1 := redTiles[currIdx].y
		x2 := redTiles[nextIdx].x
		y2 := redTiles[nextIdx].y
		cx1 := min(slices.Index(xCoords, x1)*2, slices.Index(xCoords, x2)*2)
		cx2 := max(slices.Index(xCoords, x1)*2, slices.Index(xCoords, x2)*2)
		cy1 := min(slices.Index(yCoords, y1)*2, slices.Index(yCoords, y2)*2)
		cy2 := max(slices.Index(yCoords, y1)*2, slices.Index(yCoords, y2)*2)

		for cx := cx1; cx <= cx2; cx++ {
			for cy := cy1; cy <= cy2; cy++ {
				grid[cy][cx] = 1
			}
		}
	}

	// inverse flood-fill algo to find all tiles outside the polygon
	outside := make(map[coordinate]bool)
	outside[coordinate{-1, -1}] = true
	queue := []coordinate{{-1, -1}}

	for len(queue) > 0 {
		x := queue[0].x
		y := queue[0].y
		queue = queue[1:]

		neighbors := []coordinate{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
		for _, n := range neighbors {
			if n.x < -1 || n.y < -1 || n.x > len(grid[0]) || n.y > len(grid) {
				continue
			} else if n.x >= 0 && n.x < len(grid[0]) && n.y >= 0 && n.y < len(grid) && grid[n.y][n.x] == 1 {
				continue
			} else if outside[coordinate{n.x, n.y}] {
				continue
			}

			outside[coordinate{n.x, n.y}] = true
			queue = append(queue, n)
		}
	}

	// fill tiles within polygon if they we didn't just find them with the inverse flood-fill
	for y := range len(grid) {
		for x := range len(grid[0]) {
			if grid[y][x] == 0 && !outside[coordinate{x, y}] {
				grid[y][x] = 1
			}
		}
	}

	// brute-force find largest area
	maxArea := int64(0)
	for i := range redTiles {
		for j := range redTiles {
			// calculate area in "non-compressed" space
			if i == j {
				continue
			}

			xDist := int64(math.Abs(float64(redTiles[i].x-redTiles[j].x))) + 1
			yDist := int64(math.Abs(float64(redTiles[i].y-redTiles[j].y))) + 1
			area := xDist * yDist

			// just continue early if this can't be the biggest
			if area < maxArea {
				continue
			}

			// use "compressed space" to check if every cell in
			// this area of the grid is red/green tile
			minX := min(compressedRedTiles[i].x, compressedRedTiles[j].x)
			maxX := max(compressedRedTiles[i].x, compressedRedTiles[j].x)
			minY := min(compressedRedTiles[i].y, compressedRedTiles[j].y)
			maxY := max(compressedRedTiles[i].y, compressedRedTiles[j].y)
			valid := true
			for y := minY; y <= maxY; y++ {
				for x := minX; x <= maxX; x++ {
					if grid[y][x] == 0 {
						valid = false
						break
					}
				}
				if !valid {
					break
				}
			}

			if valid {
				maxArea = area
			}
		}
	}

	// fmt.Println()
	// visualize(grid)

	fmt.Printf("Star 2 answer: %d \n", maxArea)
}

func visualize(grid [][]int) {
	height := len(grid)
	width := len(grid[0])

	for y := range height {
		for x := range width {
			if grid[y][x] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}

		fmt.Println()
	}
}
