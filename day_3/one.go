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
	for _, e := range data {
		var res string
		splitPos := len(e) / 2
		compartmentOne := strings.Split(e[:splitPos], "")
		compartmentTwo := strings.Split(e[splitPos:], "")

		for _, i := range compartmentOne {
			for _, j := range compartmentTwo {
				if i == j {
					res = i
					break
				}
			}
		}

		for len(res) > 0 {
			r, size := utf8.DecodeRuneInString(res)
			if r > 96 { // Lower case
				r -= 96
			} else { // Upper case
				r -= 38
			}
			output += int(r)

			res = res[size:]
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
