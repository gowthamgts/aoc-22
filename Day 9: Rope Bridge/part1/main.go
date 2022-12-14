package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	row int
	col int
}

func (p Position) Add(n Position) Position {
	return Position{row: p.row + n.row, col: p.col + n.col}
}

func (p Position) Substract(n Position) Position {
	return Position{row: p.row - n.row, col: p.col - n.col}
}

func (p Position) Move(n Position) Position {
	return Position{row: p.row + Sign(n.row), col: p.col + Sign(n.col)}
}

func (p Position) String() string {
	return fmt.Sprintf("%d_%d", p.row, p.col)
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Sign(n int) int {
	if n < 0 {
		return -1
	}

	if n > 0 {
		return +1
	}

	return 0
}

func MovePosition(position string) Position {
	switch position {
	case "U":
		return Position{0, -1}
	case "D":
		return Position{0, 1}
	case "L":
		return Position{-1, 0}
	case "R":
		return Position{1, 0}
	}
	return Position{0, 0}
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("failed to open input.txt file")
	}
	defer file.Close()

	H, T := Position{row: 0, col: 0}, Position{row: 0, col: 0}

	var empty struct{}
	touchMap := map[Position]interface{}{T: empty}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := strings.TrimSpace(scanner.Text())
		direction := currentLine[0:1]
		times, _ := strconv.Atoi(currentLine[2:])

		moveDifference := MovePosition(direction)
		for i := 0; i < times; i++ {
			H = H.Add(moveDifference)

			tailDifference := H.Substract(T)
			if Abs(tailDifference.row) > 1 || Abs(tailDifference.col) > 1 {
				// move tail
				T = T.Move(Position{row: tailDifference.row, col: tailDifference.col})
				touchMap[T] = empty
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(touchMap))
}
