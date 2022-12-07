package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	EMPTY          struct{}
	DISTINCT_CHARS int = 14
)

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

		if len(set) == DISTINCT_CHARS {
			// match. call the marker
			return i + 1
		}

		// delete the first element and insert the last one
		if i > (DISTINCT_CHARS - 2) {
			set[currentLine[i-(DISTINCT_CHARS-1)]] = set[currentLine[i-(DISTINCT_CHARS-1)]] - 1
			if set[currentLine[i-(DISTINCT_CHARS-1)]] == 0 {
				delete(set, currentLine[i-(DISTINCT_CHARS-1)])
			}
		}
	}

	return -1
}
