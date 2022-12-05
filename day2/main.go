package main

import (
	"bufio"
	_ "bufio"
	"fmt"
	_ "fmt"
	_ "io"
	"os"
	_ "os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func points(score string) int {
	hands := strings.Split(score, " ")
	if hands[1] == "X" { // Rock
		if hands[0] == "B" { // Paper
			return 1
		} else if hands[0] == "A" { // Rock
			return 4
		} else { // Scissors
			return 7
		}
	}
	if hands[1] == "Y" { // Paper
		if hands[0] == "C" { // Scissors
			return 2
		} else if hands[0] == "B" { // Paper
			return 5
		} else { // Rock
			return 8
		}
	}
	if hands[1] == "Z" { // Scissors
		if hands[0] == "A" { // Rock
			return 3
		} else if hands[0] == "C" { // Scissors
			return 6
		} else { // Paper
			return 9
		}
	}
	return 0
}

func matchFix(score string) int {
	hands := strings.Split(score, " ")
	if hands[1] == "X" { // lose
		if hands[0] == "A" { // Rock
			return 3
		} else if hands[0] == "B" { // Paper
			return 1
		} else { // Scissors
			return 2
		}
	}
	if hands[1] == "Y" { // draw
		if hands[0] == "A" {
			return 1 + 3
		} else if hands[0] == "B" {
			return 2 + 3
		} else {
			return 3 + 3
		}
	}
	if hands[1] == "Z" { // win
		if hands[0] == "A" { // Rock
			return 2 + 6
		} else if hands[0] == "B" { // Paper
			return 3 + 6
		} else { // Scissors
			return 1 + 6
		}
	}
	return 0
}

func main() {
	test := false
	fileName := "day2/input.txt"
	if test {
		fileName = "day2/test.txt"
	}
	dat, err := os.Open(fileName)
	check(err)
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)

	total := 0
	totalMatchfix := 0

	for fileScanner.Scan() {
		score := fileScanner.Text()
		pnts := points(score)
		total = total + pnts

		mpnts := matchFix(score)
		totalMatchfix = totalMatchfix + mpnts

	}

	fmt.Println(total)
	fmt.Println(totalMatchfix)

	dat.Close()
}
