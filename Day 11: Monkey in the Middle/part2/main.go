package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	REGEX  = `^Monkey \d:(?:[\n])\s{2}(?:Starting items: )([\S|\s]*)(?:[\n])\s{2}(?:Operation: new = )([\S|\s]*)(?:[\n])\s{2}(?:Test:\sdivisible by )(\d+)(?:[\n])\s{4}(?:If true: throw to monkey )(\d+)(?:[\n])\s{4}(?:If false: throw to monkey )(\d+)$`
	ROUNDS = 10_000
)

var (
	regex = regexp.MustCompile(REGEX)
)

type Monkey struct {
	items         []int
	divisor       int
	success       int
	fail          int
	inspectCount  int
	operationFunc func(old int) int
}

func NewMonkey(matches []string) *Monkey {
	var (
		m               = new(Monkey)
		itemFields      = strings.Split(matches[0], ", ")
		operationFields = strings.Fields(matches[1])
		operationVar, _ = strconv.Atoi(operationFields[2])
		divisor, _      = strconv.Atoi(matches[2])
		success, _      = strconv.Atoi(matches[3])
		fail, _         = strconv.Atoi(matches[4])
	)

	// assign items
	m.items = make([]int, 0)
	for _, item := range itemFields {
		n, _ := strconv.Atoi(item)
		m.items = append(m.items, n)
	}

	// assign operationFunc
	switch operationFields[1] {
	case "+":
		m.operationFunc = func(old int) int {
			if operationFields[2] == "old" {
				return old + old
			}
			return old + operationVar
		}
	case "*":
		m.operationFunc = func(old int) int {
			if operationFields[2] == "old" {
				return old * old
			}
			return old * operationVar
		}
	}

	m.divisor = divisor
	m.success = success
	m.fail = fail

	return m
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("failed to open input.txt file")
	}
	defer file.Close()

	// var declaractions
	var (
		monkeys       []*Monkey
		scanner              = bufio.NewScanner(file)
		monkeyDetails string = ""
	)

	for scanner.Scan() {
		currentLine := scanner.Text()

		if strings.TrimSpace(currentLine) == "" {
			// parse regex

			matches := regex.FindStringSubmatch(monkeyDetails)
			monkeys = append(monkeys, NewMonkey(matches[1:]))

			monkeyDetails = ""
			continue
		}

		if monkeyDetails == "" {
			monkeyDetails += currentLine
		} else {
			monkeyDetails += "\n" + currentLine
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	commonDivisor := 1
	for _, m := range monkeys {
		commonDivisor *= m.divisor
	}

	// process for each monkey
	for i := 0; i < ROUNDS; i++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				var (
					newLevel     = m.operationFunc(item)
					updatedLevel = newLevel % commonDivisor
				)

				if updatedLevel%m.divisor == 0 {
					monkeys[m.success].items = append(monkeys[m.success].items, updatedLevel)
				} else {
					monkeys[m.fail].items = append(monkeys[m.fail].items, updatedLevel)
				}

				m.inspectCount++
			}
			m.items = make([]int, 0)
		}
	}

	inspectArr := []int{}
	for _, m := range monkeys {
		inspectArr = append(inspectArr, m.inspectCount)
	}

	// sort the array
	sort.Sort(sort.Reverse(sort.IntSlice(inspectArr)))
	fmt.Println(inspectArr[0] * inspectArr[1])
}
