package main

import (
	"bufio"
	_ "bufio"
	"fmt"
	_ "fmt"
	_ "io"
	"os"
	_ "os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func reverseArray(arrIn []string) []string {
	arrOut := make([]string, 0)
	for i := len(arrIn) - 1; i >= 0; i-- {
		arrOut = append(arrOut, arrIn[i])
	}
	return arrOut
}

func main() {
	test := false
	fileName := "day5/input.txt"
	if test {
		fileName = "day5/test.txt"
	}
	dat, err := os.Open(fileName)
	check(err)
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)

	stacks := make(map[int][]string)

	readingStacks := true
	readingMoves := false

	for fileScanner.Scan() {
		row := fileScanner.Text()
		if readingStacks {
			if row[0:7] == " 1   2 " {
				readingStacks = false
				for i := 1; i <= len(stacks); i++ {
					stacks[i] = reverseArray(stacks[i])
				}
			} else {
				for p := 0; p < len(row); p++ {
					if string(row[p]) != " " && string(row[p]) != "[" && string(row[p]) != "]" {
						position := ((p - 1) / 4) + 1
						if stacks[position] == nil {
							emptyArray := make([]string, 0)
							stacks[position] = emptyArray
						}
						stacks[position] = append(stacks[position], string(row[p]))
					}
				}
			}
		} else if row == "" {
			readingMoves = true
		} else if readingMoves {
			instructions := strings.Split(row, " ")
			amount, err := strconv.Atoi(instructions[1])
			from, err := strconv.Atoi(instructions[3])
			to, err := strconv.Atoi(instructions[5])
			check(err)
			//for x := 0; x < amount; x++ {
			//	stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			//	stacks[from] = stacks[from][:len(stacks[from])-1]
			//}
			tot := len(stacks[from])
			van := tot - amount
			stacks[to] = append(stacks[to], stacks[from][van:tot]...)
			stacks[from] = stacks[from][:len(stacks[from])-amount]

		}
	}

	lastCrates := ""
	for c := 1; c <= len(stacks); c++ {
		last := stacks[c][len(stacks[c])-1]
		lastCrates = lastCrates + last
	}

	dat.Close()

	fmt.Println(lastCrates)
}
