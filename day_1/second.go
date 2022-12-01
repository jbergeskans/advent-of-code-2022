package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func findMaxSubset(data [][]int) int {
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

func parseData(data []string) [][]int {
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
	data := readData("indata_2.txt")
	elfs := parseData(data)
	res := findMaxSubset(elfs)
	fmt.Println(res)
}
