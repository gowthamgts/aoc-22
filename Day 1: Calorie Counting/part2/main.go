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

	var calories [3]int
	currentCalories := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		if currentLine == "" {
			if currentCalories >= calories[2] {
				calories[0] = calories[1]
				calories[1] = calories[2]
				calories[2] = currentCalories
			} else if currentCalories >= calories[1] {
				calories[0] = calories[1]
				calories[1] = currentCalories
			} else if currentCalories >= calories[0] {
				calories[0] = currentCalories
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

	fmt.Println(calories[0] + calories[1] + calories[2])
}
