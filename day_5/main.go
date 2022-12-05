package main

import (
	"bufio"
	"fmt"
	"os"
)

func makeMoves(stack [][]string, instructions [][]int) [][]string {
	for _, instr := range instructions {
		numToMove := instr[0]
		fromStack := instr[1] - 1
		toStack := instr[2] - 1

		subStackFrom := stack[fromStack]
		subStackTo := stack[toStack]

		for i := 0; i < numToMove; i++ {
			val := subStackFrom[0]
			subStackTo = append([]string{val}, subStackTo...)
			subStackFrom = subStackFrom[1:]

			stack[fromStack] = subStackFrom
			stack[toStack] = subStackTo
		}
	}

	return stack
}

func parseInstructions(instructions []string) [][]int {
	var output [][]int
	for _, e := range instructions {
		var moveNumber, fromStack, toStack int
		_, err := fmt.Sscanf(e, "move %d from %d to %d", &moveNumber, &fromStack, &toStack)

		if err != nil {
			fmt.Println(err)
		}

		output = append(output, []int{moveNumber, fromStack, toStack})
	}

	return output
}

func parseStack(stack []string) [][]string {
	var out [][]string
	tempStack := stack[:len(stack)-1]
	for i := 1; i < len(stack[0]); i += 4 {
		var tmp []string
		for _, e := range tempStack {
			if string(e[i]) != " " {
				tmp = append(tmp, string(e[i]))
			}
		}
		out = append(out, tmp)
	}

	return out
}

func splitData(data []string) ([]string, []string) {
	var stack, instructions []string
	for i, e := range data {
		if e == "" {
			return data[:i], data[i+1:]
		}
	}

	return stack, instructions
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

func one(stack [][]string, instructions [][]int) {
	finalStack := makeMoves(stack, instructions)
	for _, e := range finalStack {
		fmt.Print(e[0])
	}
}

func main() {
	data := readData("input.txt")
	stack, instructions := splitData(data)
	newStack := parseStack(stack)
	newInstructions := parseInstructions(instructions)
	one(newStack, newInstructions)
}
