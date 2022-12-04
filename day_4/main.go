package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func one(data []string) int {
	var output int

	for _, e := range data {
		ass := strings.Split(e, ",")
		elfOne, elfTwo := strings.Split(ass[0], "-"), strings.Split(ass[1], "-")

		elfOne_low, _ := strconv.Atoi(elfOne[0])
		elfOne_high, _ := strconv.Atoi(elfOne[1])
		elfTwo_low, _ := strconv.Atoi(elfTwo[0])
		elfTwo_high, _ := strconv.Atoi(elfTwo[1])

		if elfOne_low <= elfTwo_low && elfOne_high >= elfTwo_high {
			output += 1
		} else if elfTwo_low <= elfOne_low && elfTwo_high >= elfOne_high {
			output += 1
		}
	}

	return output
}

func two(data []string) int {
	var output int

	for _, e := range data {
		var elfOne_low, elfOne_high, elfTwo_low, elfTwo_high int
		_, err := fmt.Sscanf(e, "%d-%d,%d-%d", &elfOne_low, &elfOne_high, &elfTwo_low, &elfTwo_high)

		if err != nil {
			panic(err)
		}

		if elfTwo_low <= elfOne_high && elfTwo_high >= elfOne_low {
			output++
		} else if elfOne_low <= elfTwo_high && elfOne_high >= elfTwo_low {
			output++
		}
	}

	return output
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
	data := readData("input.txt")
	fmt.Println("Part one:", one(data))
	fmt.Println("Part two:", two(data))
}
