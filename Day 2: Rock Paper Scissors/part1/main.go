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

func calculateCurrentScore(opponent, player byte) int {
	valueMap := make(map[byte]int)
	valueMap['A'] = 1
	valueMap['B'] = 2
	valueMap['C'] = 3
	valueMap['X'] = 1
	valueMap['Y'] = 2
	valueMap['Z'] = 3

	currentScore := valueMap[player]

	// grant 3 points if it's a draw
	if valueMap[opponent] == valueMap[player] {
		currentScore += 3
	}

	if (valueMap[opponent]+1)%3 == valueMap[player]%3 {
		currentScore += 6
	}
	return currentScore
}
