package main

import (
	"aoc/2022/aoc2022/readinput"
	"fmt"
	"strconv"
	"strings"
)

func two(data []string) {
	vals := calcCycleVals(data)

	var subtract int
	for cycle := 1; cycle <= 240; cycle++ {
		if cycle%40 == 1 {
			fmt.Print("\n")
			subtract = cycle - 1
		}

		sprite_start := cycle - subtract - 2
		sprite_end := cycle - subtract

		if sprite_start <= vals[cycle] && vals[cycle] <= sprite_end {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
}

func one(data []string) int {
	cycleVals := calcCycleVals(data)

	var out int
	for j := 20; j < len(cycleVals); j += 40 {
		out += j * cycleVals[j]
	}

	return out
}

func calcCycleVals(data []string) map[int]int {
	cycleVals := make(map[int]int)
	cycle, registerVal := 1, 1
	cycleVals[1] = 1

	for _, e := range data {
		instruction := strings.Split(e, " ")

		if instruction[0] != "noop" {
			addVal, _ := strconv.Atoi(instruction[1])
			registerVal += addVal
			cycle += 2
			cycleVals[cycle] = registerVal

		} else {
			cycle++
		}

		// Fill missing values
		for i := 1; i <= cycle; i++ {
			if cycleVals[i] == 0 {
				cycleVals[i] = cycleVals[i-1]
			}
		}
	}

	return cycleVals
}

func main() {
	data := readinput.ReadData("input.txt")
	// data := readinput.ReadData("input_test.txt")
	fmt.Println("First result:", one(data))
	fmt.Print("Second result:")
	two(data)
}
