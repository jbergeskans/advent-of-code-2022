package main

import (
	"aoc/2022/aoc2022/readinput"
	"fmt"
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

func main() {
	data := readinput.ReadData("input.txt")
	calc(data[0], 4)
	calc(data[0], 14)
}
