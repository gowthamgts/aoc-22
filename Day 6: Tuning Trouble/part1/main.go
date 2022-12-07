package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var empty struct{}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("failed to open input.txt file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()

		fmt.Println(findMarker(currentLine))
	}
}

func findMarker(currentLine string) int {
	set := make(map[byte]uint)

	for i := 0; i < len(currentLine); i++ {
		if _, ok := set[currentLine[i]]; ok {
			set[currentLine[i]] = set[currentLine[i]] + 1
			// continue
		} else {
			set[currentLine[i]] = 1
		}

		if len(set) == 4 {
			// match. call the marker
			return i + 1
		}

		// delete the first element and insert the last one
		if i > 2 {
			set[currentLine[i-3]] = set[currentLine[i-3]] - 1
			if set[currentLine[i-3]] == 0 {
				delete(set, currentLine[i-3])
			}
		}
	}

	return -1
}
