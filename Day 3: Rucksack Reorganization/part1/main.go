package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var empty struct{}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("failed to open input.txt file")
	}
	defer file.Close()

	totalSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()

		totalSum += calculatePriority(currentLine)
	}

	fmt.Println(totalSum)
}

func calculatePriority(currentLine string) int {
	splitIndex := len(currentLine) / 2
	split1 := currentLine[0:splitIndex]
	split2 := currentLine[splitIndex:]

	split1Map := make(map[byte]struct{})

	for i := 0; i < len(split1); i++ {
		split1Map[split1[i]] = empty
	}

	var commonItem byte
	for i := 0; i < len(split2); i++ {
		if _, ok := split1Map[split2[i]]; ok {
			commonItem = split2[i]
			break
		}
	}

	if commonItem >= 'a' && commonItem <= 'z' {
		return int(commonItem) - 'a' + 1
	}
	return int(commonItem) - 'A' + 27
}
