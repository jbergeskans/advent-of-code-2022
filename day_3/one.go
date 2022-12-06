package main

import (
	"aoc/2022/aoc2022/readinput"
	"fmt"
	"strings"
	"unicode/utf8"
)

func two(data []string) {
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

func one(data []string) {
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

func main() {
	data := readinput.ReadData("input.txt")
	one(data)
	two(data)
}
