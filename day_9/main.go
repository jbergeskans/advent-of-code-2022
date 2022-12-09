package main

import (
	"aoc/2022/aoc2022/readinput"
	"fmt"
	"strconv"
	"strings"
)

type knot struct {
	xPos int
	yPos int
}

func one(moves []string) int {
	tailPositions := make(map[string]int)
	head := knot{1, 1}
	tail := knot{1, 1}

	tailStartStr := strconv.Itoa(tail.xPos) + "_" + strconv.Itoa(tail.yPos)
	tailPositions[tailStartStr] = 1

	for _, move := range moves {
		splitMove := strings.Split(move, " ")
		direction := splitMove[0]
		steps, _ := strconv.Atoi(splitMove[1])

		// Calculate new position for head
		switch direction {
		case "R":
			head.xPos += steps
		case "L":
			head.xPos -= steps
		case "U":
			head.yPos += steps
		case "D":
			head.yPos -= steps
		}

		// Calculate tails response to position change
		if calcAbs((head.xPos - tail.xPos)) > 1 {
			for calcAbs((head.xPos - tail.xPos)) != 1 {
				switch direction {
				case "R":
					tail.xPos++
				case "L":
					tail.xPos--
				}

				if (head.yPos - tail.yPos) > 0 {
					tail.yPos++
				} else if (head.yPos - tail.yPos) < 0 {
					tail.yPos--
				}
				tailNewPos := strconv.Itoa(tail.xPos) + "_" + strconv.Itoa(tail.yPos)
				tailPositions[tailNewPos] = 1
			}
		} else if calcAbs((head.yPos - tail.yPos)) > 1 {
			for calcAbs((head.yPos - tail.yPos)) != 1 {
				switch direction {
				case "U":
					tail.yPos++
				case "D":
					tail.yPos--
				}

				if (head.xPos - tail.xPos) > 0 {
					tail.xPos++
				} else if (head.xPos - tail.xPos) < 0 {
					tail.xPos--
				}
				tailNewPos := strconv.Itoa(tail.xPos) + "_" + strconv.Itoa(tail.yPos)
				tailPositions[tailNewPos] = 1
			}
		}
	}

	// fmt.Println(head, tail)
	// fmt.Println(tailPositions)
	return len(tailPositions)
}

func calcAbs(inVal int) int {
	if inVal < 0 {
		return -inVal
	}
	return inVal
}

func main() {
	data := readinput.ReadData("input.txt")
	// data := readinput.ReadData("input_test.txt")
	fmt.Println("Answer one:", one(data))
}
