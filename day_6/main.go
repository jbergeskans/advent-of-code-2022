package main

import (
	"bufio"
	"fmt"
	"os"
)

func calc(data string, addNum int) {
	for i := 0; i < len(data)-addNum-1; i++ {
		subString := data[i : i+addNum]
		tempMap := make(map[rune]int)

		for _, e := range subString {
			tempMap[e] = 1
		}

		if len(tempMap) == addNum {
			fmt.Println(i + addNum)
			break
		}
	}
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
		val := fileScanner.Text()

		if err != nil {
			fmt.Println(err)
		}

		data = append(data, val)
	}

	return data
}

func main() {
	// data := readData("input_test.txt")
	data := readData("input.txt")
	calc(data[0], 4)
	calc(data[0], 14)
}
