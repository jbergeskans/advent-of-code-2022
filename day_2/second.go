package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func run(
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

func readData(path string) []string {
	var data []string
	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		// val, err := strconv.Atoi(fileScanner.Text())
		val := fileScanner.Text()

		if err != nil {
			fmt.Println(err)
		}

		data = append(data, val)
	}

	return data
}

func main() {
	roundResMap := map[string]int{"X": 0, "Y": 3, "Z": 6}
	pointMap := map[string]int{"A": 1, "B": 2, "C": 3}
	x := map[string][3]string{
		"A": {"C", "A", "B"}, // input: lose, draw, win
		"B": {"A", "B", "C"},
		"C": {"B", "C", "A"},
	}
	data := readData("indata_2.txt")
	run(data, pointMap, roundResMap, x)
}
