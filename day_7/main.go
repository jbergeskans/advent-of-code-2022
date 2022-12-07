package main

import (
	"aoc/2022/aoc2022/readinput"
	"fmt"
	"strconv"
	"strings"
)

func one(data []string) map[string]int {
	// var pwd string
	pwd := "root"
	sizeMap := make(map[string]int)
	for _, e := range data {
		input := strings.Split(e, " ")
		if input[0] == "$" && input[1] == "cd" { // change dir
			if input[2] != ".." { // cd newdir
				if input[2] != "/" {
					pwd += "/" + input[2]
				}
			} else { // cd ..
				tmp_pwd := strings.Split(pwd, "/")
				pwd = strings.Join(tmp_pwd[:len(tmp_pwd)-1], "/")
			}
		} else if input[0] != "dir" { // file entry
			val, _ := strconv.Atoi(input[0])
			sizeMap[pwd] += val
		}
	}

	out := make(map[string]int)
	for k, v := range sizeMap {
		out[k] = v
		for sk, sv := range sizeMap {
			if k != sk && strings.Contains(sk, k) {
				out[k] += sv
			}
		}
	}

	return out
}

func calcOne(out map[string]int) {
	var output int
	for _, j := range out {
		if j <= 100000 {
			output += j
		}
	}

	fmt.Println(output)
}

func calcTwo(out map[string]int) {
	freeSpace := 70000000 - out["root"]
	missingSpace := 30000000 - freeSpace
	currMin := 70000000

	for _, v := range out {
		if missingSpace < v && v < currMin {
			currMin = v
		}
	}

	fmt.Println(missingSpace)
	fmt.Println(currMin)
}

func main() {
	// data := readinput.ReadData("input_test.txt")
	data := readinput.ReadData("input.txt")
	calcOne(one(data))
	calcTwo(one(data))
}
