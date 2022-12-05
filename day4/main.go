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

func fullyContains(assignmentA string, assignmentB string) bool {
	nrsA := strings.Split(assignmentA, "-")
	nrsB := strings.Split(assignmentB, "-")
	fromA, err := strconv.Atoi(nrsA[0])
	toA, err := strconv.Atoi(nrsA[1])
	fromB, err := strconv.Atoi(nrsB[0])
	toB, err := strconv.Atoi(nrsB[1])
	check(err)
	return fromA <= fromB && toA >= toB
}

func overlap(assignmentA string, assignmentB string) bool {
	nrsA := strings.Split(assignmentA, "-")
	nrsB := strings.Split(assignmentB, "-")
	//fromA, err := strconv.Atoi(nrsA[0])
	toA, err := strconv.Atoi(nrsA[1])
	fromB, err := strconv.Atoi(nrsB[0])
	toB, err := strconv.Atoi(nrsB[1])
	check(err)
	return toA >= fromB && toA <= toB

}

func main() {
	test := false
	fileName := "day4/input.txt"
	if test {
		fileName = "day4/test.txt"
	}
	dat, err := os.Open(fileName)
	check(err)
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)

	pairs := 0
	overlapPair := 0

	for fileScanner.Scan() {
		row := fileScanner.Text()
		assignments := strings.Split(row, ",")
		if fullyContains(assignments[0], assignments[1]) {
			pairs++
		} else if fullyContains(assignments[1], assignments[0]) {
			pairs++
		}

		if overlap(assignments[0], assignments[1]) {
			overlapPair++
		} else if overlap(assignments[1], assignments[0]) {
			overlapPair++
		}
	}

	fmt.Println(pairs)
	fmt.Println(overlapPair)

	dat.Close()
}
