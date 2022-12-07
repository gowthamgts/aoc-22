package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	totalColumns     = 9
	crateRegex       = regexp.MustCompile(`(?:([\[A-Z\]])|\s{3})`)
	instructionRegex = regexp.MustCompile(`move\s(\d+)\sfrom\s(\d+)\sto\s(\d+)`)
)

type Crate struct {
	value []string
	size  int
}

func (s *Crate) PushToFront(item string) {
	if item == "" {
		return
	}
	s.value = append([]string{item}, s.value...)
	s.size++
}

func (s *Crate) Push(item string) {
	if item == "" {
		return
	}
	s.value = append(s.value, item)
	s.size++
}

func (s *Crate) PushBulk(items []string) {
	s.value = append(s.value, items...)
	s.size += len(items)
}

func (s *Crate) Pop() string {
	if s.size == 0 {
		return ""
	}
	value := s.value[s.size-1]
	s.value = s.value[:s.size-1]
	s.size--
	return value
}

func (s *Crate) PopBulk(n int) []string {
	var items []string
	if s.size-n < 0 {
		items = s.value
		s.value = []string{}
		s.size = 0
	} else {
		items = s.value[s.size-n:]
		s.value = s.value[:s.size-n]
		s.size = s.size - n
	}
	return items
}

func NewCrate() *Crate {
	return &Crate{}
}

func (s *Crate) String() string {
	return fmt.Sprintf("{values: %v, size: %d}", s.value, s.size)
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("failed to open input.txt file")
	}
	defer file.Close()

	// create initial stacks
	var stacks []*Crate
	for i := 0; i < totalColumns; i++ {
		stacks = append(stacks, NewCrate())
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()

		if crateRegex.MatchString(currentLine) {
			for i := 0; i+3 <= len(currentLine); i = i + 4 {
				if currentLine[i] == ' ' {
					// empty space continue
					continue
				}
				if i+4 > len(currentLine) {
					stacks[i/4].PushToFront(currentLine[i:])
				} else {
					stacks[i/4].PushToFront(currentLine[i : i+4])
				}
			}
		} else if instructionRegex.MatchString(currentLine) {
			// parse instructions
			matches := instructionRegex.FindStringSubmatch(currentLine)
			boxesToBeMoved, _ := strconv.Atoi(matches[1])
			fromCrate, _ := strconv.Atoi(matches[2])
			toCrate, _ := strconv.Atoi(matches[3])

			items := stacks[fromCrate-1].PopBulk(boxesToBeMoved)
			if len(items) > 0 {
				stacks[toCrate-1].PushBulk(items)
			}
		}
	}

	items := ""
	for _, s := range stacks {
		item := s.Pop()
		if item != "" {
			items = items + string(item[1])
		}
	}

	fmt.Println(items)
}
