package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("failed to open input.txt file")
	}
	defer file.Close()

	var rows [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := strings.TrimSpace(scanner.Text())

		row := []int{}

		for i := 0; i < len(currentLine); i++ {
			height, _ := strconv.Atoi(currentLine[i : i+1])
			row = append(row, height)
		}

		rows = append(rows, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(findHighestScenicScore(rows))
}

func findHighestScenicScore(matrix [][]int) int {
	height := len(matrix)
	width := len(matrix[0])

	highestScenicScore := 0

	for row := 1; row < height-1; row++ {
		for col := 1; col < width-1; col++ {
			currentHeight := matrix[row][col]
			top, bottom, left, right := 0, 0, 0, 0

			// from current to top
			for i := row - 1; i >= 0; i-- {
				top++
				if matrix[i][col] >= currentHeight {
					break
				}
			}

			// from current to bottom
			for i := row + 1; i < height; i++ {
				bottom++
				if matrix[i][col] >= currentHeight {
					break
				}
			}

			// from current to left
			for i := col - 1; i >= 0; i-- {
				left++
				if matrix[row][i] >= currentHeight {
					break
				}
			}

			// from current to right
			for i := col + 1; i < width; i++ {
				right++
				if matrix[row][i] >= currentHeight {
					break
				}
			}

			currentScore := top * bottom * left * right
			if currentScore > highestScenicScore {
				highestScenicScore = currentScore
			}
		}
	}

	return highestScenicScore
}
