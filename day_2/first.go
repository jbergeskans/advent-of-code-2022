package main

import (
	"aoc/2022/aoc2022/readinput"
	"fmt"
	"strings"
)

func one(data []string, pointMap map[string]int, roundResMap map[string]int) {
	var totalScore int
	for _, e := range data {
		roundRes := strings.Replace(e, " ", "", -1)
		totalScore += roundResMap[roundRes]

		vals := strings.Split(e, " ")
		myVal := vals[1]
		totalScore += pointMap[myVal]
	}

	fmt.Println(totalScore)
}

func second(
	data []string,
	pointMap map[string]int,
	roundResMap map[string]int,
	signMap map[string][3]string) {
	var totalScore int
	for _, e := range data {
		vals := strings.Split(e, " ")
		opponentValue := vals[0]
		roundOutcome := vals[1]
		totalScore += roundResMap[roundOutcome]

		var idx int
		if roundOutcome == "X" { // lose round
			idx = 0
		} else if roundOutcome == "Y" { // draw round
			idx = 1
		} else { // win round
			idx = 2
		}

		sign := signMap[opponentValue][idx]
		totalScore += pointMap[sign]
	}
	fmt.Println(totalScore)
}

func main() {
	pointMap := map[string]int{"X": 1, "Y": 2, "Z": 3}
	roundResMap := map[string]int{
		"AX": 3, "AY": 6, "AZ": 0,
		"BX": 0, "BY": 3, "BZ": 6,
		"CX": 6, "CY": 0, "CZ": 3,
	}
	data := readinput.ReadData("indata_1.txt")
	one(data, pointMap, roundResMap)

	roundResMap = map[string]int{"X": 0, "Y": 3, "Z": 6}
	pointMap = map[string]int{"A": 1, "B": 2, "C": 3}
	x := map[string][3]string{
		"A": {"C", "A", "B"}, // input: lose, draw, win
		"B": {"A", "B", "C"},
		"C": {"B", "C", "A"},
	}
	data = readinput.ReadData("indata_2.txt")
	second(data, pointMap, roundResMap, x)
}
