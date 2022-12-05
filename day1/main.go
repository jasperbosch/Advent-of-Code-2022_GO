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
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Sort(elves []int) []int {
	for i := 0; i < len(elves)-1; i++ {
		for j := i + 1; j < len(elves); j++ {
			if elves[i] < elves[j] {
				tmp := elves[i]
				elves[i] = elves[j]
				elves[j] = tmp
			}
		}
	}
	return elves
}

func main() {
	test := false
	fileName := "day1/input.txt"
	if test {
		fileName = "day1/test.txt"
	}
	dat, err := os.Open(fileName)
	check(err)
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)

	calorieCount := 0
	maxCalories := 0
	var elves []int

	for fileScanner.Scan() {
		calories := fileScanner.Text()
		if calories == "" {
			fmt.Println(calorieCount)
			elves = append(elves, calorieCount)
			maxCalories = Max(maxCalories, calorieCount)
			calorieCount = 0
		} else {
			i, err := strconv.Atoi(calories)
			check(err)
			calorieCount = calorieCount + i
		}
	}
	fmt.Println(calorieCount)
	elves = append(elves, calorieCount)
	maxCalories = Max(maxCalories, calorieCount)
	fmt.Println(maxCalories)

	elves = Sort(elves)
	fmt.Println(elves[0] + elves[1] + elves[2])

	dat.Close()
}
