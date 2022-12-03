package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func parseData(data []string) {
	var output int

	for i := 0; i < len(data); i += 3 {
		elfOne := strings.Split(data[i], "")
		elfTwo := strings.Split(data[i+1], "")
		elfThree := strings.Split(data[i+2], "")
		var subRes string

		for _, e := range elfOne {
			for _, x := range elfTwo {
				if e == x {
					for _, y := range elfThree {
						if e == y {
							subRes = e
						}
					}
				}
			}
		}

		for len(subRes) > 0 {
			r, size := utf8.DecodeRuneInString(subRes)
			if r > 96 { // Lower case
				r -= 96
			} else {
				r -= 38
			}
			output += int(r)

			subRes = subRes[size:]
		}
	}

	fmt.Println(output)
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
	data := readData("input.txt")
	parseData(data)
}
