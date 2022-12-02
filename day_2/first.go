package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func run(data []string, pointMap map[string]int, roundResMap map[string]int) {
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
	pointMap := map[string]int{"X": 1, "Y": 2, "Z": 3}
	roundResMap := map[string]int{
		"AX": 3, "AY": 6, "AZ": 0,
		"BX": 0, "BY": 3, "BZ": 6,
		"CX": 6, "CY": 0, "CZ": 3,
	}
	data := readData("indata_1.txt")
	run(data, pointMap, roundResMap)
}
