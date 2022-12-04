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

	totalSum := 0

	var lines [3]string
	var index uint32 = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines[index%3] = scanner.Text()

		index++
		if index%3 == 0 {
			totalSum += calculatePriority(lines)
		}
	}

	fmt.Println(totalSum)
}

func calculatePriority(lines [3]string) int {
	line1Map := make(map[byte]struct{})
	line2Map := make(map[byte]struct{})

	// push all elements in the map
	for i := 0; i < len(lines[0]); i++ {
		line1Map[lines[0][i]] = empty
	}

	// iterate over line 2 and push common items to the map
	for i := 0; i < len(lines[1]); i++ {
		if _, ok := line1Map[lines[1][i]]; ok {
			line2Map[lines[1][i]] = empty
		}
	}

	// iterate over line 3 and decide the final item
	var commonItem byte
	for i := 0; i < len(lines[2]); i++ {
		if _, ok := line2Map[lines[2][i]]; ok {
			commonItem = lines[2][i]
			break
		}
	}

	if commonItem >= 'a' && commonItem <= 'z' {
		return int(commonItem) - 'a' + 1
	}
	return int(commonItem) - 'A' + 27
}
