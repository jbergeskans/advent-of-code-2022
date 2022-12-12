package main

import (
	"aoc/2022/aoc2022/readinput"
	"fmt"
	"strings"
)

type node struct {
	selfVal                          rune
	xPos, yPos, f, g, h, nodeId      int
	top, right, bottom, left, parent *node
	visited                          bool
	visVal, takenDir                 string
}

type nodeList struct {
	nodes []*node
	maxY  int
}

func (nl *nodeList) GenerateNode(x, y, nodeId int, r rune, sR string) {
	if y > nl.maxY {
		nl.maxY = y
	}
	newNode := &node{
		selfVal:  r,
		nodeId:   nodeId,
		xPos:     x,
		yPos:     y,
		f:        0,
		g:        0,
		h:        0,
		top:      nil,
		right:    nil,
		bottom:   nil,
		left:     nil,
		parent:   nil,
		visited:  false,
		visVal:   sR,
		takenDir: ".",
	}
	nl.nodes = append(nl.nodes, newNode)
}

func (nl *nodeList) LinkNodes() {
	for idx := range nl.nodes {
		for subIdx := range nl.nodes {
			if int(nl.nodes[idx].selfVal) >= (int(nl.nodes[subIdx].selfVal))-1 { // Only populate if potential path can be traversed
				if nl.nodes[idx].xPos == nl.nodes[subIdx].xPos {
					if nl.nodes[idx].yPos == (nl.nodes[subIdx].yPos + 1) {
						// If X is same and Y is one less, subIdx node is at bottom
						nl.nodes[idx].bottom = nl.nodes[subIdx]
					} else if nl.nodes[idx].yPos == (nl.nodes[subIdx].yPos - 1) {
						// If X is same and Y is one more, subIdx node is at top
						nl.nodes[idx].top = nl.nodes[subIdx]
					}
				} else if nl.nodes[idx].yPos == nl.nodes[subIdx].yPos {
					if nl.nodes[idx].xPos == (nl.nodes[subIdx].xPos + 1) {
						// If Y is same and X is one less, subIdx node is at left
						nl.nodes[idx].left = nl.nodes[subIdx]
					} else if nl.nodes[idx].xPos == (nl.nodes[subIdx].xPos - 1) {
						// If Y is same and X is one more, subIdx node is at right
						nl.nodes[idx].right = nl.nodes[subIdx]
					}
				}
			}
		}
	}
}

func (nL *nodeList) PrintNodes() {
	for _, e := range nL.nodes {
		fmt.Println(e)
	}
}

func (nL *nodeList) VisNodes() {
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
					fmt.Print(e.takenDir)
				}
			}
		}
		fmt.Print("\n")
	}
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

func one(data [][]string) int {
	var open, closed []*node
	var endNode *node
	var out int
	nL := genNodeList(data)
	nL.LinkNodes()

	// 0. Add start note to open
	for idx := range nL.nodes {
		if nL.nodes[idx].visVal == "S" {
			nL.nodes[idx].visited = true
			open = append(open, nL.nodes[idx])
		} else if nL.nodes[idx].visVal == "E" {
			endNode = nL.nodes[idx]
		}
	}

	end := false
	for {
		var lowest *node
		// fmt.Println(open)
		// fmt.Println("l1", lowest)

		// 1. Find lowest f value and place in closed
		for idx := range open {
			if lowest != nil {
				if open[idx].f <= lowest.f && !open[idx].visited {
					lowest = open[idx]
				}
			} else {
				lowest = open[idx]
			}
		}
		lowest.visited = true
		lowest.g++
		lowest.h = calcAbs(lowest.xPos-endNode.xPos) + calcAbs(lowest.yPos-endNode.yPos)
		lowest.f = lowest.g + lowest.h
		lowest.takenDir = "X"
		closed = append(closed, lowest)

		// 2. Check adjecent nodes
		adjecent := []*node{lowest.top, lowest.right, lowest.bottom, lowest.left}

		// fmt.Println("lowest:", lowest.nodeId, lowest.f, lowest)

		for a_idx := range adjecent {
			// nL.VisNodes()
			if adjecent[a_idx] == endNode {
				// fmt.Println("ADASDASDASDASDAS", lowest.parent.f+1)
				adjecent[a_idx].takenDir = "O"
				for {
					if lowest.parent != nil {
						out++
						lowest = lowest.parent
					} else {
						lowest.takenDir = "O"
						out++
						end = true
						break
					}
				}
			} else if adjecent[a_idx] != nil {
				// 2.1 Checking closed list
				var inClosed, inOpen *node

				for c_idx := range closed {
					if closed[c_idx].nodeId == adjecent[a_idx].nodeId {
						inClosed = closed[c_idx]
					}
				}

				// 2.2. Checking open list
				for o_idx := range open {
					if open[o_idx].nodeId == adjecent[a_idx].nodeId {
						inOpen = open[o_idx]
					}
				}

				if inClosed == nil {
					adjecent[a_idx].parent = lowest
					adjecent[a_idx].g = adjecent[a_idx].parent.g + 1
					adjecent[a_idx].h = calcAbs(adjecent[a_idx].xPos-endNode.xPos) + calcAbs(adjecent[a_idx].yPos-endNode.yPos)
					adjecent[a_idx].f = adjecent[a_idx].g + adjecent[a_idx].h
					// fmt.Println("Adj", adjecent[a_idx].nodeId, adjecent[a_idx].f, adjecent[a_idx])
					if inOpen == nil {
						open = append(open, adjecent[a_idx])
					} else {
						if adjecent[a_idx].g < inOpen.g {
							inOpen.visited = true
							open = append(open, adjecent[a_idx])
						}
					}
				}
			}
		}
		// fmt.Println("open", open)
		// fmt.Println("closed", closed)
		// fmt.Println("=======")
		// nL.VisNodes()
		if end {
			break
		}
	}
	nL.VisNodes()
	return out
}

func genNodeList(data [][]string) *nodeList {
	nL := nodeList{}
	nodeId := 0
	for y, e := range data { // each row of input, y here is actual_y - 1
		for x, f := range e { // each col of input, x here is actual_x - 1
			var fR []rune
			if f == "S" {
				fR = []rune("a")
			} else if f == "E" {
				fR = []rune("z")
			} else {
				fR = []rune(f)
			}
			if len(fR) == 1 {
				nL.GenerateNode(x+1, y+1, nodeId, fR[0], f)
				nodeId++
			}
		}
	}

	return &nL
}

func calcAbs(inVal int) int {
	if inVal < 0 {
		return -inVal
	}
	return inVal
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

func main() {
	// data := readinput.ReadData("input_test.txt")
	data := readinput.ReadData("input.txt")
	reversedData := reverseSlice(data)
	fmt.Println("First", one(reversedData))
}
