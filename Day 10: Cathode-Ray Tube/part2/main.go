package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Matrix [6][40]bool

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("failed to open input.txt file")
	}
	defer file.Close()

	cycle := 0
	x := 1

	var matrix Matrix

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := strings.TrimSpace(scanner.Text())

		fields := strings.Fields(currentLine)

		DrawLine(&matrix, &x, &cycle)
		if fields[0] == "addx" {
			// perform addx
			DrawLine(&matrix, &x, &cycle)
			n, _ := strconv.Atoi(fields[1])
			x += n
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == true {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}

}

func DrawLine(matrix *Matrix, x, cycle *int) {
	row, col := *cycle/40, *cycle%40
	// see whether we have to draw
	if *cycle%40 <= *x+1 && *cycle%40 >= *x-1 {
		// draw it
		(*matrix)[row][col] = true
	}

	*cycle++
}
