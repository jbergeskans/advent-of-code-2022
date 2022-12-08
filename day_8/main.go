package main

import (
	"aoc/2022/aoc2022/readinput"
	"fmt"
	"strconv"
	"strings"
)

type tree struct {
	xPos   int
	yPos   int
	height int
	up     *tree
	down   *tree
	left   *tree
	right  *tree
}

func calcOne(trees []tree, maxCoord int) int {
	var visCount int

	for _, tree := range trees {
		if tree.xPos == 1 || tree.xPos == maxCoord || tree.yPos == 1 || tree.yPos == maxCoord {
			visCount++
		} else {
			curr := tree
			baseHeight := tree.height
			foundExit := false

			for {
				if curr.left == nil {
					visCount++
					foundExit = true
					break
				} else if baseHeight > curr.left.height {
					curr = *curr.left
				} else {
					break
				}
			}

			if !foundExit {
				curr = tree
				for {
					if curr.right == nil {
						visCount++
						foundExit = true
						break
					} else if baseHeight > curr.right.height {
						curr = *curr.right
					} else {
						break
					}
				}
			}

			if !foundExit {
				curr = tree
				for {
					if curr.up == nil {
						visCount++
						foundExit = true
						break
					} else if baseHeight > curr.up.height {
						curr = *curr.up
					} else {
						break
					}
				}
			}

			if !foundExit {
				curr = tree
				for {
					if curr.down == nil {
						visCount++
						foundExit = true
						break
					} else if baseHeight > curr.down.height {
						curr = *curr.down
					} else {
						break
					}
				}
			}
		}
	}

	return visCount
}

func calcTwo(trees []tree, maxCoord int) int {
	var visCountOut [][]int
	for _, tree := range trees {
		if tree.xPos != 1 {
			visCount := 0
			var visCountInner []int
			curr := tree
			baseHeight := tree.height
			// foundExit := false

			for {
				if curr.left != nil {
					visCount++
					if baseHeight > curr.left.height {
						curr = *curr.left
					} else {
						visCountInner = append(visCountInner, visCount)
						break
					}
				} else {
					visCountInner = append(visCountInner, visCount)
					break
				}
			}

			curr = tree
			visCount = 0
			for {
				if curr.right != nil {
					visCount++
					if baseHeight > curr.right.height {
						curr = *curr.right
					} else {
						visCountInner = append(visCountInner, visCount)
						break
					}
				} else {
					visCountInner = append(visCountInner, visCount)
					break
				}
			}

			curr = tree
			visCount = 0
			for {
				if curr.up != nil {
					visCount++
					if baseHeight > curr.up.height {
						curr = *curr.up
					} else {
						visCountInner = append(visCountInner, visCount)
						break
					}
				} else {
					visCountInner = append(visCountInner, visCount)
					break
				}
			}

			curr = tree
			visCount = 0
			for {
				if curr.down != nil {
					visCount++
					if baseHeight > curr.down.height {
						curr = *curr.down
					} else {
						visCountInner = append(visCountInner, visCount)
						break
					}
				} else {
					visCountInner = append(visCountInner, visCount)
					break
				}
			}

			visCountOut = append(visCountOut, visCountInner)
		} else {
			visCountOut = append(visCountOut, []int{10, 0, 0, 0})
		}
	}

	currMax := 0
	for _, e := range visCountOut {
		currVal := e[0] * e[1] * e[2] * e[3]

		if currVal > currMax {
			currMax = currVal
		}
	}

	return currMax
}

func genTrees(data [][]string) []tree {
	var treeSlice []tree
	for idxY, valY := range data {
		// fmt.Println(valY)
		for idxX, valX := range valY {
			height, _ := strconv.Atoi(valX)
			tree := tree{idxX + 1, idxY + 1, height, nil, nil, nil, nil}
			treeSlice = append(treeSlice, tree)
		}
	}

	for idx, e := range treeSlice {
		for idxtwo, i := range treeSlice {
			if e.yPos == i.yPos {
				if e.xPos == i.xPos-1 {
					treeSlice[idx].right = &treeSlice[idxtwo]
				} else if e.xPos == i.xPos+1 {
					treeSlice[idx].left = &treeSlice[idxtwo]
				}
			} else if e.xPos == i.xPos {
				if e.yPos == i.yPos+1 {
					treeSlice[idx].down = &treeSlice[idxtwo]
				} else if e.yPos == i.yPos-1 {
					treeSlice[idx].up = &treeSlice[idxtwo]
				}
			}
		}
	}

	return treeSlice
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
	parsedData := reverseSlice(data)
	trees := genTrees(parsedData)
	fmt.Println("Solution one:", calcOne(trees, len(data)))
	fmt.Println("Solution two:", calcTwo(trees, len(data)))
}
