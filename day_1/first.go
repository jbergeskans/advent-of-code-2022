package main

import (
	"aoc/2022/aoc2022/readinput"
	"fmt"
	"sort"
	"strconv"
)

func findMaxSubsetOne(data [][]int) int {
	var max int
	for _, e := range data {
		var sum int
		for _, i := range e {
			sum += i
		}

		if sum > max {
			max = sum
		}
	}

	return max
}

func parseDataOne(data []string) [][]int {
	var subSlice []int
	var outputSlice [][]int

	for _, e := range data {
		if e != "" {
			val, _ := strconv.Atoi(e)
			subSlice = append(subSlice, val)
		} else {
			outputSlice = append(outputSlice, subSlice)
			subSlice = nil
		}
	}

	return outputSlice
}

func findMaxSubsetTwo(data [][]int) int {
	var calCounts []int
	var sum int

	for _, e := range data {
		var sum int
		for _, i := range e {
			sum += i
		}
		calCounts = append(calCounts, sum)
	}

	sort.Ints(calCounts)
	no_elfs := len(calCounts)

	for i := 1; i < 4; i++ {
		sum += calCounts[no_elfs-i]
	}

	return sum
}

func parseDataTwo(data []string) [][]int {
	var subSlice []int
	var outputSlice [][]int

	for _, e := range data {
		if e != "" {
			val, _ := strconv.Atoi(e)
			subSlice = append(subSlice, val)
		} else {
			outputSlice = append(outputSlice, subSlice)
			subSlice = nil
		}
	}

	outputSlice = append(outputSlice, subSlice)

	return outputSlice
}

func main() {
	data := readinput.ReadData("indata_1.txt")
	elfs := parseDataOne(data)
	res := findMaxSubsetOne(elfs)
	fmt.Println(res)

	data = readinput.ReadData("indata_2.txt")
	elfs = parseDataTwo(data)
	res = findMaxSubsetTwo(elfs)
	fmt.Println(res)
}
