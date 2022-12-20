package main

import (
	"aoc/2022/aoc2022/readinput"
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	top, right, bottom, left *node
	f, g, h, xPos, yPos      int
	selfVal                  rune
	visVal                   string
}

type nodeList struct {
	nodes []*node
	maxY  int
}

func (nL *nodeList) VisNodesStr() {
	for currLine := nL.maxY; currLine != 0; currLine-- {
		var tmp []*node
		for _, e := range nL.nodes {
			if e.yPos == currLine {
				tmp = append(tmp, e)
			}
		}

		for xPos := 0; xPos <= len(tmp); xPos++ {
			for _, e := range tmp {
				if e.xPos == xPos {
					fmt.Print(e.visVal + " ")
				}
			}
		}
		fmt.Print("\n")
	}
}

func (nL *nodeList) GenerateNode(x, y int, r rune, sR string) {
	if y > nL.maxY {
		nL.maxY = y
	}

	newNode := &node{
		top:     nil,
		right:   nil,
		bottom:  nil,
		left:    nil,
		f:       0,
		g:       0,
		h:       0,
		xPos:    x,
		yPos:    y,
		selfVal: r,
		visVal:  sR,
	}
	nL.nodes = append(nL.nodes, newNode)
}

func (nL *nodeList) LinkNodes() {
	for idx := range nL.nodes {
		for subIdx := range nL.nodes {
			if int(nL.nodes[idx].selfVal) >= (int(nL.nodes[subIdx].selfVal))-1 { // Only populate if potential path can be traversed
				if nL.nodes[idx].xPos == nL.nodes[subIdx].xPos {
					if nL.nodes[idx].yPos == (nL.nodes[subIdx].yPos + 1) {
						// If X is same and Y is one less, subIdx node is at bottom
						nL.nodes[idx].bottom = nL.nodes[subIdx]
					} else if nL.nodes[idx].yPos == (nL.nodes[subIdx].yPos - 1) {
						// If X is same and Y is one more, subIdx node is at top
						nL.nodes[idx].top = nL.nodes[subIdx]
					}
				} else if nL.nodes[idx].yPos == nL.nodes[subIdx].yPos {
					if nL.nodes[idx].xPos == (nL.nodes[subIdx].xPos + 1) {
						// If Y is same and X is one less, subIdx node is at left
						nL.nodes[idx].left = nL.nodes[subIdx]
					} else if nL.nodes[idx].xPos == (nL.nodes[subIdx].xPos - 1) {
						// If Y is same and X is one more, subIdx node is at right
						nL.nodes[idx].right = nL.nodes[subIdx]
					}
				}
			}
		}
	}
}

func genNodeList(data [][]string) *nodeList {
	nL := nodeList{}
	for y, e := range data {
		for x, f := range e {
			var r []rune
			if f == "S" {
				r = []rune("a")
			} else if f == "E" {
				r = []rune("z")
			} else {
				r = []rune(f)
			}
			nL.GenerateNode(x+1, y+1, r[0], f)
		}
	}
	return &nL
}

func two(nL *nodeList) int {
	var open, closed []*node
	var endNode *node

	for idx, e := range nL.nodes {
		if e.visVal == "S" {
			open = append(open, nL.nodes[idx])
		} else if e.visVal == "E" {
			endNode = nL.nodes[idx]
		}
	}

	for {
		var currNode *node
		var nodeIdx int

		// Identify lowest f
		for idx := range open {
			if currNode == nil { // First iteration only
				currNode = open[idx]
			} else if open[idx].f < currNode.f {
				currNode = open[idx]
				nodeIdx = idx
			}
		}

		// Remove lowest f node from open
		open = slicePop(open, nodeIdx)

		// Check all adjeceten nodes
		adjecent := []*node{currNode.top, currNode.right, currNode.bottom, currNode.left}
		for idx := range adjecent {
			if adjecent[idx] != nil {
				if adjecent[idx] == endNode {
					return currNode.g + 1
				} else {
					skip := false
					adjecent[idx].g = currNode.g + 1
					adjecent[idx].h = calcAbs(adjecent[idx].xPos-endNode.xPos) + calcAbs(adjecent[idx].yPos-endNode.yPos)
					adjecent[idx].f = adjecent[idx].g + adjecent[idx].h
					adjecent[idx].visVal = "*"
					for sidx := range open {
						if open[sidx].xPos == adjecent[idx].xPos && open[sidx].yPos == adjecent[idx].yPos {
							if open[sidx].f <= adjecent[idx].f {
								skip = true
							}
						}
					}
					for didx := range closed {
						if closed[didx].xPos == adjecent[idx].xPos && closed[didx].yPos == adjecent[idx].yPos {
							if closed[didx].f <= adjecent[idx].f {
								skip = true
							}
						}
					}
					if !skip {
						open = append(open, adjecent[idx])
					}
				}
			}
		}
		closed = append(closed, currNode)
		if len(open) == 0 {
			break
		}
	}
	return 0
}

func one(data [][]string) int {
	var open, closed []*node
	var endNode *node
	nL := genNodeList(data)
	nL.LinkNodes()

	for idx, e := range nL.nodes {
		if e.visVal == "S" {
			open = append(open, nL.nodes[idx])
		} else if e.visVal == "E" {
			endNode = nL.nodes[idx]
		}
	}

	for {
		var currNode *node
		var nodeIdx int

		// Identify lowest f
		for idx := range open {
			if currNode == nil { // First iteration only
				currNode = open[idx]
			} else if open[idx].f < currNode.f {
				currNode = open[idx]
				nodeIdx = idx
			}
		}

		// Remove lowest f node from open
		open = slicePop(open, nodeIdx)

		// Check all adjeceten nodes
		adjecent := []*node{currNode.top, currNode.right, currNode.bottom, currNode.left}
		for idx := range adjecent {
			if adjecent[idx] != nil {
				if adjecent[idx] == endNode {
					return currNode.g + 1
				} else {
					skip := false
					adjecent[idx].g = currNode.g + 1
					adjecent[idx].h = calcAbs(adjecent[idx].xPos-endNode.xPos) + calcAbs(adjecent[idx].yPos-endNode.yPos)
					adjecent[idx].f = adjecent[idx].g + adjecent[idx].h
					adjecent[idx].visVal = "*"
					for sidx := range open {
						if open[sidx].xPos == adjecent[idx].xPos && open[sidx].yPos == adjecent[idx].yPos {
							if open[sidx].f <= adjecent[idx].f {
								skip = true
							}
						}
					}
					for didx := range closed {
						if closed[didx].xPos == adjecent[idx].xPos && closed[didx].yPos == adjecent[idx].yPos {
							if closed[didx].f <= adjecent[idx].f {
								skip = true
							}
						}
					}
					if !skip {
						open = append(open, adjecent[idx])
					}
				}
			}
		}
		closed = append(closed, currNode)
		if len(open) == 0 {
			break
		}
	}
	return 0
}

func calcAbs(inVal int) int {
	if inVal < 0 {
		return -inVal
	}
	return inVal
}

func slicePop(s []*node, idx int) []*node {
	return append(s[:idx], s[idx+1:]...)
}

func reverseSlice(data []string) [][]string {
	var tmpSlice, revSlice [][]string

	for _, e := range data {
		tmpSlice = append(tmpSlice, strings.Split(e, ""))
	}

	for i := range data {
		revSlice = append(revSlice, tmpSlice[len(data)-1-i])
	}

	return revSlice
}

func reGen(graphData [][]string) *nodeList {
	nL := genNodeList(graphData)
	nL.LinkNodes()
	for idx := range nL.nodes {
		if nL.nodes[idx].visVal == "S" {
			nL.nodes[idx].visVal = "a"
		}
	}
	return nL
}

func triggerTwo(graphData [][]string) int {
	var out []int
	wasStart := make(map[string]int)
	for {
		done := true
		nL := reGen(graphData)
		for idx := range nL.nodes {
			if nL.nodes[idx].visVal == "a" {
				tmpStr := strconv.Itoa(nL.nodes[idx].xPos) + "_" + strconv.Itoa(nL.nodes[idx].yPos)
				if wasStart[tmpStr] < 1 {
					nL.nodes[idx].visVal = "S"
					wasStart[tmpStr]++
					done = false
					break
				}
			}
		}
		if done {
			break
		}
		retVal := two(nL)
		if retVal != 0 {
			out = append(out, retVal)
		}
	}

	var low int
	for _, e := range out {
		if e != 0 {
			if e < low {
				low = e
			} else if low == 0 {
				low = e
			}
		}
	}
	return low
}

func main() {
	// data := readinput.ReadData("input_test.txt")
	data := readinput.ReadData("input.txt")
	graphData := reverseSlice(data)
	fmt.Println("First:", one(graphData))
	fmt.Println("Second:", triggerTwo(graphData))
}
