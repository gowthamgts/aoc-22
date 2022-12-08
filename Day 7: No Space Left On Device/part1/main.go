package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	fileRegex             = regexp.MustCompile(`(\d+)\s*`)
	directoryRegex        = regexp.MustCompile(`dir\s(.*)`)
	commandRegex          = regexp.MustCompile(`\$\s(.*)`)
	totalSum       uint32 = 0
)

type Directory struct {
	name        string
	filesSize   uint32
	parent      *Directory
	directories map[string]*Directory
}

func (d *Directory) String() string {
	if d.parent == nil {
		return fmt.Sprintf("name: %s, filesSize: %d, parent: %v, directories: %v", d.name, d.filesSize, nil, d.directories)
	}
	return fmt.Sprintf("name: %s, filesSize: %d, parent: %v, directories: %v", d.name, d.filesSize, d.parent.name, d.directories)
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("failed to open input.txt file")
	}
	defer file.Close()

	var root, current *Directory

	root = &Directory{
		name:        "/",
		parent:      nil,
		directories: make(map[string]*Directory),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()

		if commandRegex.MatchString(currentLine) {
			if currentLine == "$ cd /" {
				current = root
			} else if currentLine == "$ cd .." {
				// go up one level
				current = current.parent
			} else if strings.HasPrefix(currentLine, "$ cd ") {
				// we're trying to go into a new directory
				current = current.directories[strings.Split(currentLine, " cd ")[1]]
			} else if currentLine == "$ ls" {
				// ignore
			}
		} else if directoryRegex.MatchString(currentLine) {
			// add the directory to the current folder
			newDirectory := Directory{
				name:        directoryRegex.FindStringSubmatch(currentLine)[1],
				parent:      current,
				directories: make(map[string]*Directory),
			}
			current.directories[newDirectory.name] = &newDirectory
		} else if fileRegex.MatchString(currentLine) {
			// add the current file size to the directory size
			fileSize, _ := strconv.ParseUint(fileRegex.FindStringSubmatch(currentLine)[1], 10, 32)
			current.filesSize += uint32(fileSize)
		}
		// fmt.Println(currentLine)
	}

	fmt.Println(getCurrentDirSize(root))
	fmt.Println(totalSum)
}

func getCurrentDirSize(dir *Directory) uint32 {
	for _, d := range dir.directories {
		dir.filesSize += getCurrentDirSize(d)
	}

	if dir.filesSize <= 100000 {
		totalSum += dir.filesSize
	}
	return dir.filesSize
}
