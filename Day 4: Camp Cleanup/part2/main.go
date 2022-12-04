package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var empty struct{}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("failed to open input.txt file")
	}
	defer file.Close()

	totalMatchingPairs := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()

		if isInvalidPair(currentLine) {
			totalMatchingPairs++
		}
	}

	fmt.Println(totalMatchingPairs)
}

func isInvalidPair(currentLine string) bool {
	pairs := strings.Split(currentLine, ",")
	pair1 := strings.Split(pairs[0], "-")
	pair2 := strings.Split(pairs[1], "-")

	// parse first pair
	a1, _ := strconv.Atoi(pair1[0])
	a2, _ := strconv.Atoi(pair1[1])

	// parse second pair
	b1, _ := strconv.Atoi(pair2[0])
	b2, _ := strconv.Atoi(pair2[1])

	return (a1 >= b1 && a1 <= b2) || (a2 >= b1 && a2 <= b2) || (b1 >= a1 && b1 <= a2) || (b2 >= a1 && b2 <= a2)
}
