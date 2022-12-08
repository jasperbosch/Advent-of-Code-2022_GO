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

func main() {
	test := false
	fileName := "day7/input.txt"
	if test {
		fileName = "day7/test.txt"
	}
	dat, err := os.Open(fileName)
	check(err)
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)

	var folders []string
	folderSizes := make(map[string]int)

	for fileScanner.Scan() {
		row := fileScanner.Text()
		parts := strings.Split(row, " ")
		if parts[0] == "$" {
			// command
			if parts[1] == "cd" {
				if parts[2] == ".." {
					// go back
					folders = folders[:len(folders)-1]
				} else {
					folders = append(folders, parts[2])
				}
			}
		} else {
			if parts[0] != "dir" {
				folderName := ""
				for i := 0; i < len(folders); i++ {
					folder := folders[i]
					folderName += "/" + folder
					//if folderSizes[folderName] == 0 {
					size, err := strconv.Atoi(parts[0])
					check(err)
					folderSizes[folderName] = folderSizes[folderName] + size
					//} else {
					//	folderSizes[folderName], err = strconv.Atoi(parts[0])
					//	check(err)
					//}
				}
			}
		}
	}

	dat.Close()

	fmt.Println(folderSizes)

	totalLess := 0

	for _, value := range folderSizes {
		if value < 100000 {
			//fmt.Println(key, value)
			totalLess = totalLess + value
		}
	}

	fmt.Println("total < 100000", totalLess)
	totalUsed := folderSizes["//"]
	fmt.Println("total used", totalUsed)
	fmt.Println("Free", 70000000-totalUsed)
	inNeedOf := 30000000 - (70000000 - totalUsed)
	fmt.Println("To be freed", inNeedOf)

	smallestFolderSize := 30000000

	for _, value := range folderSizes {
		if value > inNeedOf {
			if value < smallestFolderSize {
				smallestFolderSize = value
			}
		}
	}
	fmt.Println(smallestFolderSize)
}
