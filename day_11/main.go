package main

import (
	"aoc/2022/aoc2022/readinput"
	"fmt"
	"strconv"
	"strings"
)

type monkey struct {
	id           int
	currItems    []int
	operation    string
	divBy        int
	passTrue     int
	passFalse    int
	inspectCount int
}

func one(monkeys []monkey) int {
	for i := 0; i < 20; i++ {
		for idx := range monkeys {
			for _, item := range monkeys[idx].currItems {
				var worryLevel int
				op := strings.Split(monkeys[idx].operation, " ")
				val, err := strconv.Atoi(op[1])

				monkeys[idx].inspectCount++

				if op[0] == "mul" {
					if err != nil {
						worryLevel = item * item
					} else {
						worryLevel = item * val
					}
				} else if op[0] == "add" {
					worryLevel = item + val
				}
				worryLevel = int(worryLevel / 3)
				if worryLevel%monkeys[idx].divBy == 0 {
					monkeys[monkeys[idx].passTrue].currItems = append(monkeys[monkeys[idx].passTrue].currItems, worryLevel)
				} else {
					monkeys[monkeys[idx].passFalse].currItems = append(monkeys[monkeys[idx].passFalse].currItems, worryLevel)
				}
			}
			monkeys[idx].currItems = nil
		}
	}

	var one, two int
	for _, mnk := range monkeys {
		if mnk.inspectCount > one {
			two = one
			one = mnk.inspectCount
		} else if mnk.inspectCount > two && mnk.inspectCount < one {
			two = mnk.inspectCount
		}
	}
	return one * two
}

func two(monkeys []monkey, end int, sm int) int {
	// for _, e := range monkeys {
	// 	fmt.Println(e.id, e.inspectCount, e.currItems)
	// }
	for i := 0; i < end; i++ {
		// fmt.Println(monkeys)
		for idx := range monkeys {
			// fmt.Println("Running for", idx, "Svals:", monkeys[idx].currItems)
			// fmt.Println(idx, len(monkeys[idx].currItems))
			for _, item := range monkeys[idx].currItems {
				var worryLevel int
				op := strings.Split(monkeys[idx].operation, " ")
				val, _ := strconv.Atoi(op[1])

				monkeys[idx].inspectCount++

				if op[0] == "mul" {
					if val == 0 {
						// fmt.Println(idx, val)
						worryLevel = item * item
					} else {
						worryLevel = item * val
					}
				} else if op[0] == "add" {
					worryLevel = item + val
				}

				// fmt.Println(idx, worryLevel, worryLevel%monkeys[idx].divBy, monkeys[idx].passTrue, monkeys[idx].passFalse)
				if worryLevel%monkeys[idx].divBy == 0 {
					monkeys[monkeys[idx].passTrue].currItems = append(monkeys[monkeys[idx].passTrue].currItems, worryLevel%sm)
					// fmt.Println("==", monkeys[idx].passTrue, monkeys[monkeys[idx].passTrue].currItems)
				} else {
					monkeys[monkeys[idx].passFalse].currItems = append(monkeys[monkeys[idx].passFalse].currItems, worryLevel%sm)
					// fmt.Println("==", monkeys[idx].passFalse, monkeys[monkeys[idx].passFalse].currItems)
				}
			}
			monkeys[idx].currItems = nil
		}
	}

	var one, two int
	for _, mnk := range monkeys {
		if mnk.inspectCount > one {
			two = one
			one = mnk.inspectCount
		} else if mnk.inspectCount > two && mnk.inspectCount < one {
			two = mnk.inspectCount
		}
	}
	return one * two
}

func genMonkeys(data []string) ([]monkey, int) {
	var monkeyStats [][]string
	var tmp []string
	var monkeyList []monkey
	sm := 1

	for _, e := range data {
		if e == "" {
			monkeyStats = append(monkeyStats, tmp)
			tmp = nil
		} else {
			tmp = append(tmp, e)
		}
	}
	monkeyStats = append(monkeyStats, tmp)

	for _, input := range monkeyStats {
		inId := strings.Split(input[0], "")
		mnkId, _ := strconv.Atoi(inId[len(inId)-2])

		inItems := strings.Split(input[1], ": ")
		inItems = strings.Split(inItems[1], ", ")
		var mnkItems []int
		for _, e := range inItems {
			inItemInt, _ := strconv.Atoi(e)
			mnkItems = append(mnkItems, inItemInt)
		}

		inOp := strings.Split(input[2], " ")
		var mnkInOp string
		if inOp[len(inOp)-2] == "*" {
			mnkInOp = "mul " + inOp[len(inOp)-1]
		} else {
			mnkInOp = "add " + inOp[len(inOp)-1]
		}

		inTest := strings.Split(input[3], " ")
		mnkTest, _ := strconv.Atoi(inTest[len(inTest)-1])
		sm *= mnkTest

		inTrue := strings.Split(input[4], " ")
		mnkTrue, _ := strconv.Atoi(inTrue[len(inTrue)-1])

		inFalse := strings.Split(input[5], " ")
		mnkFalse, _ := strconv.Atoi(inFalse[len(inFalse)-1])

		newMonkey := monkey{mnkId, mnkItems, mnkInOp, mnkTest, mnkTrue, mnkFalse, 0}
		monkeyList = append(monkeyList, newMonkey)
	}

	return monkeyList, sm
}

func main() {
	// data := readinput.ReadData("input_test.txt")
	data := readinput.ReadData("input.txt")
	monkeys, _ := genMonkeys(data)
	fmt.Println("Part one:", one(monkeys))
	monkeys, sm := genMonkeys(data)
	fmt.Println("Part two:", two(monkeys, 10000, sm))
}
