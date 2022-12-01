package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("failed to open input.txt file")
	}
	defer file.Close()

	maxCalories := 0
	currentCalories := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		if currentLine == "" {
			if currentCalories >= maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
			continue
		}

		n, err := strconv.Atoi(currentLine)
		if err != nil {
			log.Fatal("error occurred while converting string to int", err)
		}

		currentCalories += n
	}

	fmt.Println(maxCalories)
}
