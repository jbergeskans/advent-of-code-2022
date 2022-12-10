package main

import (
	"aoc/2022/aoc2022/readinput"
	"fmt"
	"strconv"
	"strings"
)

type knot struct {
	xPos   int
	yPos   int
	parent *knot
	child  *knot
	pos    int
}

type knotList struct {
	head *knot
}

func one(moves []string) int {
	tailPositions := make(map[string]int)
	head := knot{1, 1, nil, nil, 1}
	tail := knot{1, 1, nil, nil, 1}

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

func (kl *knotList) Insert(pos int) {
	newNode := &knot{1, 1, nil, nil, pos}
	if kl.head == nil {
		kl.head = newNode
	} else {
		p := kl.head
		for p.child != nil {
			p = p.child
		}
		newNode.parent = p
		p.child = newNode
	}
}

func Show(l *knotList) {
	p := l.head
	for p != nil {
		fmt.Printf("-> %v ", p)
		p = p.child
	}
}

func moveTwo(node *knot, direction string, tailPositions map[string]int) {
	for node != nil {
		// Need to move along the x-axis
		if calcAbs((node.parent.xPos - node.xPos)) > 1 {
			// Move knots until they've reached their end position
			for calcAbs((node.parent.xPos - node.xPos)) != 1 {
				if node.parent.xPos-node.xPos > 1 {
					node.xPos++
				} else if node.parent.xPos-node.xPos < -1 {
					node.xPos--
				}

				if (node.parent.yPos - node.yPos) > 0 {
					node.yPos++
				} else if (node.parent.yPos - node.yPos) < 0 {
					node.yPos--
				}
			}
			// Need tom move along the y-axis
		} else if calcAbs((node.parent.yPos - node.yPos)) > 1 {
			// Move knots until they've reached their end position
			for calcAbs((node.parent.yPos - node.yPos)) != 1 {
				if node.parent.yPos-node.yPos > 1 {
					node.yPos++
				} else if node.parent.yPos-node.yPos < -1 {
					node.yPos--
				}

				if (node.parent.xPos - node.xPos) > 0 {
					node.xPos++
				} else if (node.parent.xPos - node.xPos) < 0 {
					node.xPos--
				}
			}
		}

		if node.pos == 9 {
			tailNewPos := strconv.Itoa(node.xPos) + "_" + strconv.Itoa(node.yPos)
			tailPositions[tailNewPos] = 1
		}

		node = node.child
	}
}

func two(moves []string) int {
	tailPositions := make(map[string]int)
	tailPositions["1_1"] = 1
	kl := knotList{}

	// Create knots
	for i := 0; i < 10; i++ {
		kl.Insert(i)
	}

	for _, move := range moves {
		splitMove := strings.Split(move, " ")
		direction := splitMove[0]
		steps, _ := strconv.Atoi(splitMove[1])

		// Calculate head movement
		for i := 0; i < steps; i++ {
			switch direction {
			case "R":
				kl.head.xPos += 1
			case "L":
				kl.head.xPos -= 1
			case "U":
				kl.head.yPos += 1
			case "D":
				kl.head.yPos -= 1
			}

			// Calculate how knots react to head movement
			moveTwo(kl.head.child, direction, tailPositions)
		}
	}

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
	fmt.Println("Answer two:", two(data))
}
