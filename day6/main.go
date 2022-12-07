package main

import (
	"bufio"
	_ "bufio"
	"fmt"
	_ "fmt"
	_ "io"
	"os"
	_ "os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkDoubles(sequence []string) bool {
	for i := 0; i < len(sequence)-1; i++ {
		for j := i + 1; j < len(sequence); j++ {
			if sequence[i] == sequence[j] {
				return true
			}
		}
	}
	return false
}

func main() {
	test := false
	fileName := "day6/input.txt"
	if test {
		fileName = "day6/test.txt"
	}
	dat, err := os.Open(fileName)
	check(err)
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var sequence []string
		row := fileScanner.Text()
		fmt.Println(row)
		chars := row
		for i := 0; i < len(chars); i++ {
			char := chars[i]
			if len(sequence) >= 14 {
				if !checkDoubles(sequence) {
					fmt.Println("Marker is ", string(char), i)
					break
				}
				sequence = sequence[1:]
			}
			sequence = append(sequence, string(char))
		}
	}

	dat.Close()
}
