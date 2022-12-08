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
	MAX_SPACE    = uint32(70000000)
	NEEDED_SPACE = uint32(30000000)
)

var (
	fileRegex      = regexp.MustCompile(`(\d+)\s*`)
	directoryRegex = regexp.MustCompile(`dir\s(.*)`)
	commandRegex   = regexp.MustCompile(`\$\s(.*)`)
	sortedSizes    = []uint32{}
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
	}

	usedSpace := getCurrentDirSize(root)
	unusedSpace := MAX_SPACE - usedSpace
	neededSpace := NEEDED_SPACE - unusedSpace

	for _, size := range sortItems(sortedSizes) {
		if size >= neededSpace {
			fmt.Println(size)
			break
		}
	}
}

func getCurrentDirSize(dir *Directory) uint32 {
	for _, d := range dir.directories {
		dir.filesSize += getCurrentDirSize(d)
	}

	sortedSizes = append(sortedSizes, dir.filesSize)
	return dir.filesSize
}

func sortItems(s []uint32) []uint32 {
	// simple bubble sort for now
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				// swap
				s[i], s[j] = s[j], s[i]
			}
		}
	}

	return s
}
