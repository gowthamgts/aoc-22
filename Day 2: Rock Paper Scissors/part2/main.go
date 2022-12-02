package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("failed to open input.txt file")
	}
	defer file.Close()

	var totalScore = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()

		totalScore += calculateCurrentScore(currentLine[0], currentLine[2])
	}

	fmt.Println(totalScore)
}

func calculateCurrentScore(opponent, result byte) int {
	valueMap := make(map[byte]int)
	valueMap['A'] = 1
	valueMap['B'] = 2
	valueMap['C'] = 3
	valueMap['X'] = 1
	valueMap['Y'] = 2
	valueMap['Z'] = 3

	resultMap := make(map[int]byte)
	resultMap[0] = 'Z'
	resultMap[1] = 'X'
	resultMap[2] = 'Y'
	resultMap[3] = 'Z'

	var currentScore int

	var player byte
	if result == 'X' {
		// we need to lose
		player = resultMap[(valueMap[opponent]-1)%3]
		currentScore += valueMap[player]
	} else if result == 'Y' {
		// we need to draw
		player = resultMap[valueMap[opponent]]
		currentScore += valueMap[player] + 3
	} else if result == 'Z' {
		// we need to win
		player = resultMap[(valueMap[opponent]+1)%3]
		currentScore += valueMap[player] + 6
	}

	return currentScore
}
