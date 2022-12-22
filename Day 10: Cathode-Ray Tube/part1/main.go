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

	cycle := 0
	x := 1
	strength := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := strings.TrimSpace(scanner.Text())

		fields := strings.Fields(currentLine)

		AddCycle(&x, &cycle, &strength)
		if fields[0] == "addx" {
			// perform addx
			AddCycle(&x, &cycle, &strength)
			n, _ := strconv.Atoi(fields[1])
			x += n
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(strength)
}

func AddCycle(x, cycle, strength *int) {
	*cycle++
	if *cycle%40 == 20 {
		*strength += *cycle * *x
	}
}
