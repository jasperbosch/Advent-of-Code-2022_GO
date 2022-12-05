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

func priority(similar string) int {
	bv := similar[0]
	charcode := int(bv)
	if charcode >= 97 {
		// lowercase
		charcode -= 96
	} else {
		// uppercase
		charcode = charcode - 65 + 27
	}
	return charcode
}

func priority2(similar int) int {
	charcode := similar
	if charcode >= 97 {
		// lowercase
		charcode -= 96
	} else {
		// uppercase
		charcode = charcode - 65 + 27
	}
	return charcode
}

func main() {
	test := false
	fileName := "day3/input.txt"
	if test {
		fileName = "day3/test.txt"
	}
	dat, err := os.Open(fileName)
	check(err)
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	var rucksacs []string

	for fileScanner.Scan() {
		rucksac := fileScanner.Text()
		rucksacs = append(rucksacs, rucksac)

		rContent := strings.Split(rucksac, "")
		rSize := len(rContent)
		rContentA := rContent[0 : rSize/2]
		rContentB := rContent[rSize/2 : rSize]
		similar := ""
		// find in both
		for i := 0; i < len(rContentA) && similar == ""; i++ {
			for j := 0; j < len(rContentB) && similar == ""; j++ {
				if rContentA[i] == rContentB[j] {
					similar = rContentA[i]
				}
			}
		}
		sum = sum + priority(similar)

	}

	fmt.Println(sum)

	dat.Close()

	sum2 := 0
	for i := 0; i < len(rucksacs); i = i + 3 {

		badge := 0
		for x := 0; x < len(rucksacs[i]) && badge == 0; x++ {
			for y := 0; y < len(rucksacs[i+1]) && badge == 0; y++ {
				if rucksacs[i][x] == rucksacs[i+1][y] {
					for z := 0; z < len(rucksacs[i+2]) && badge == 0; z++ {
						if rucksacs[i][x] == rucksacs[i+2][z] {
							badge = priority2(int(rucksacs[i][x]))
							sum2 = sum2 + badge
						}
					}
				}
			}
		}
	}
	fmt.Println(sum2)
}
