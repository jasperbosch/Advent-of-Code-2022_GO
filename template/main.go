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

func main() {
	test := true
	fileName := "day2/input.txt"
	if test {
		fileName = "day2/test.txt"
	}
	dat, err := os.Open(fileName)
	check(err)
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	dat.Close()
}
