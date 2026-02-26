package days

import (
	"cmp"
	"fmt"
	"maps"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day8(input []byte) {
	lines := slices.Collect(strings.Lines(string(input)))
	for i := range len(lines) {
		lines[i] = strings.TrimSpace(lines[i])
	}

	solve(lines)
}

type point struct {
	x int
	y int
	z int
}

type pairOfPoints struct {
	p1 point
	p2 point
}

type pairDistance struct {
	points   pairOfPoints
	distance float64
}

func (c point) String() string {
	return fmt.Sprintf("%d,%d,%d", c.x, c.y, c.z)
}

func getDistance(p1, p2 point) float64 {
	x2 := math.Pow(float64(p1.x)-float64(p2.x), 2)
	y2 := math.Pow(float64(p1.y)-float64(p2.y), 2)
	z2 := math.Pow(float64(p1.z)-float64(p2.z), 2)
	return math.Sqrt(x2 + y2 + z2)
}

func solve(lines []string) {
	points := make([]point, 0)
	circuits := make([]map[point]bool, 0)

	for _, line := range lines {
		split := strings.Split(line, ",")
		if len(split) != 3 {
			fmt.Printf("Expected string to split into 3, but didn't. String: '%s' \n", line)
			os.Exit(1)
		}

		x, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Printf("Error converting string '%s' to int. Error: %s \n", split[0], err)
		}

		y, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Printf("Error converting string '%s' to int. Error: %s \n", split[1], err)
		}

		z, err := strconv.Atoi(split[2])
		if err != nil {
			fmt.Printf("Error converting string '%s' to int. Error: %s \n", split[2], err)
		}

		points = append(points, point{x, y, z})

		circuit := map[point]bool{{x, y, z}: true}
		circuits = append(circuits, circuit)
	}

	distances := make([]pairDistance, 0)
	visited := make(map[pairOfPoints]bool)

	for _, p1 := range points {
		for _, p2 := range points {
			if p1 == p2 {
				continue
			}

			if !visited[pairOfPoints{p1, p2}] && !visited[pairOfPoints{p2, p1}] {
				pair := pairOfPoints{p1, p2}
				distances = append(distances, pairDistance{pair, getDistance(p1, p2)})
				visited[pair] = true
			}
		}
	}

	slices.SortFunc(distances, func(pd1, pd2 pairDistance) int {
		return cmp.Compare(pd1.distance, pd2.distance)
	})

	for i, distance := range distances {
		if i >= len(points) {
			break
		}

		c1Idx := -1
		c2Idx := -1
		for j, circuit := range circuits {
			if circuit[distance.points.p1] {
				c1Idx = j
			}

			if circuit[distance.points.p2] {
				c2Idx = j
			}
		}

		if c1Idx == c2Idx {
			continue
		}

		maps.Copy(circuits[c1Idx], circuits[c2Idx])
		circuits = slices.Delete(circuits, c2Idx, c2Idx+1)
	}

	slices.SortFunc(circuits, func(c1, c2 map[point]bool) int {
		return cmp.Compare(len(c2), len(c1))
	})

	fmt.Printf("Star 1 answer: %d \n", len(circuits[0])*len(circuits[1])*len(circuits[2]))

	////////// BEGIN STAR 2 //////////
	part2Answer := -1

	// same circuit building algo as before,
	//   just start at where we left off and stop when there's 1 circuit
	for i := len(points); i < len(distances); i++ {
		c1Idx := -1
		c2Idx := -1
		for j, circuit := range circuits {
			if circuit[distances[i].points.p1] {
				c1Idx = j
			}

			if circuit[distances[i].points.p2] {
				c2Idx = j
			}
		}

		if c1Idx == c2Idx {
			continue
		}

		maps.Copy(circuits[c1Idx], circuits[c2Idx])
		circuits = slices.Delete(circuits, c2Idx, c2Idx+1)

		if len(circuits) == 1 {
			part2Answer = distances[i].points.p1.x * distances[i].points.p2.x
		}
	}

	fmt.Printf("Star 2 answer: %d \n", part2Answer)
}
