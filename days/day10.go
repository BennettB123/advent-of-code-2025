package days

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type machine struct {
	desiredLights []bool
	buttons       [][]int
	// joltage []int
}

func Day10(input []byte) {
	lines := slices.Collect(strings.Lines(string(input)))
	machines := make([]machine, 0)

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// parse lights
		lightStr := line[1:strings.Index(line, "]")]
		lights := make([]bool, len(lightStr))
		for i, char := range lightStr {
			on := char == '#'
			lights[i] = on
		}

		// parse buttons
		buttonStr := line[strings.Index(line, "(") : strings.Index(line, "{")-1]
		splitButtons := strings.Split(buttonStr, " ")
		buttons := make([][]int, len(splitButtons))
		for i, bStr := range splitButtons {
			buttonInts := strings.Split(bStr[1:len(bStr)-1], ",")
			innerList := make([]int, len(buttonInts))
			for j, buttonInt := range buttonInts {
				val, err := strconv.Atoi(buttonInt)
				if err != nil {
					fmt.Printf("Failed to convert string '%s' to int. Error: %s", buttonInt, err)
					os.Exit(1)
				}
				innerList[j] = val
			}
			buttons[i] = innerList
		}

		machine := machine{lights, buttons}

		machines = append(machines, machine)
	}

	day10_star1(machines)
	day10_star2(machines)
}

// pressButton simulates a button press, toggling each light from the
// list of indexes in the button parameter. Returns the resulting light configuration
func pressButton(currLights []bool, button []int) []bool {
	newLights := make([]bool, len(currLights))
	copy(newLights, currLights)

	for _, lightIdx := range button {
		newLights[lightIdx] = !newLights[lightIdx]
	}

	return newLights
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// This approach takes a few seconds and consumes ~2GB (queue gets large)
// ////////////////////////////////////////////////////////////////////////////////////////////////////
func day10_star1(machines []machine) {
	answer := 0

	// button sequences to try
	queue := make([][]int, 0)

	for _, machine := range machines {
		clear(queue)

		for i := range machine.buttons {
			queue = append(queue, []int{i})
		}

		buttonPresses := 0

		// BFS to find quickest path to desired lights
		for len(queue) > 0 {
			thisSequence := queue[0]
			queue = queue[1:]

			currLights := make([]bool, len(machine.desiredLights))
			for _, buttonIdx := range thisSequence {
				currLights = pressButton(currLights, machine.buttons[buttonIdx])
			}

			if slices.Equal(currLights, machine.desiredLights) {
				buttonPresses = len(thisSequence)
				break
			}

			for i := range machine.buttons {
				queue = append(queue, slices.Concat(thisSequence, []int{i}))
			}
		}

		fmt.Println(len(queue))
		answer += buttonPresses
	}

	fmt.Printf("Star 1 answer: %d \n", answer)
}

func day10_star2(machines []machine) {
	fmt.Printf("Star 2 answer: %d \n", 0)
}
