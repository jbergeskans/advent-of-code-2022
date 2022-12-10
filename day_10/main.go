package main

import (
	"aoc/2022/aoc2022/readinput"
	"fmt"
	"strconv"
	"strings"
)

func one(data []string) int {
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

	var out int
	for j := 20; j < len(cycleVals); j += 40 {
		out += j * cycleVals[j]
	}

	return out
}

func main() {
	data := readinput.ReadData("input.txt")
	// data := readinput.ReadData("input_test.txt")
	fmt.Println("First result:", one(data))
}
