package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	fileSystemTotalSize  = 70000000
	fileSystemNeededSize = 30000000
)

type directory struct {
	name        string
	totalSize   int
	root        bool
	parent      *directory
	files       map[string]int
	directories map[string]*directory
}

func Day_7(input string) {
	var inputFile []string
	file, err := os.Open(input)
	CheckError(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputFile = append(inputFile, scanner.Text())
	}
	day_7_part_1(inputFile)
}

func day_7_part_1(inputFile []string) {
	var arr []int
	mainDir := directory{name: "/", totalSize: 0, root: true, files: make(map[string]int), directories: map[string]*directory{}}

	lineAnalyzed := 1

	analyzeLine(&mainDir, inputFile, lineAnalyzed)
	totalSum := iterateOverDirectories(&mainDir, 0)
	fmt.Printf("Total sum of directories with size <10000: %v.\n", totalSum)

	fmt.Printf("Root directory size: %v\n", mainDir.totalSize)
	t := mainDir.totalSize - fileSystemTotalSize + fileSystemNeededSize
	fmt.Printf("Total size needed: %v\n", t)

	totalSize := getNeededDirectorySize(&mainDir, arr, int(t))
	sort.Ints(totalSize)
	fmt.Printf("Total sum: %v.\n", totalSize[0:1])
}

func analyzeLine(currentDirectory *directory, inputFile []string, lineAnalyzed int) {
	var newCurrentDir *directory
	currentLine := inputFile[lineAnalyzed]

	if strings.HasPrefix(currentLine, "$ ls") {
	} else if strings.HasPrefix(currentLine, "dir") {
		var newDirectory directory
		dirNameFound := strings.ReplaceAll(currentLine, "dir ", "")
		_, exists := currentDirectory.directories[dirNameFound]
		if !exists {
			newDirectory = directory{name: dirNameFound, totalSize: 0, root: false, parent: currentDirectory, files: make(map[string]int), directories: map[string]*directory{}}
			currentDirectory.directories[dirNameFound] = &newDirectory
		}

	} else if strings.HasPrefix(currentLine, "$ cd") {
		destDirName := strings.ReplaceAll(currentLine, "$ cd ", "")
		if destDirName == ".." {
			newCurrentDir = currentDirectory.parent
		} else {
			newCurrentDir = currentDirectory.directories[destDirName]
		}
		currentDirectory = newCurrentDir
	} else {
		var runeSize []rune
		var size int
		var file string
		for _, val := range inputFile[lineAnalyzed] {
			if val >= '0' && val <= '9' {
				runeSize = append(runeSize, val)
			} else if val == ' ' {
				break
			}
		}
		file = strings.ReplaceAll(currentLine, string(runeSize), "")
		size, _ = strconv.Atoi(string(runeSize))
		currentDirectory.files[file] = size
		currentDirectory.totalSize = currentDirectory.totalSize + size
		updateParentsTotalSize(currentDirectory, size)

	}
	if lineAnalyzed < len(inputFile)-1 {
		lineAnalyzed++
		analyzeLine(currentDirectory, inputFile, lineAnalyzed)
	}
}

func updateParentsTotalSize(currentDirectory *directory, size int) {
	if currentDirectory.root == false {
		currentDirectory.parent.totalSize = currentDirectory.parent.totalSize + size
		updateParentsTotalSize(currentDirectory.parent, size)
	}
}

func iterateOverDirectories(directory *directory, sum int) int {
	totalSum := sum

	if directory.totalSize <= 100000 {
		totalSum = totalSum + directory.totalSize
	}

	for _, val := range directory.directories {
		totalSum = totalSum + iterateOverDirectories(val, sum)
	}
	return totalSum
}

func getNeededDirectorySize(directory *directory, arr []int, maxSize int) []int {
	var arr2 []int
	if directory.totalSize >= maxSize {
		arr2 = append(arr2, directory.totalSize)
		for _, val := range directory.directories {
			newArr := getNeededDirectorySize(val, arr2, maxSize)
			arr2 = append(arr2, newArr...)
		}
	} else {
		for _, val := range directory.directories {
			newArr := getNeededDirectorySize(val, arr2, maxSize)
			arr2 = append(arr2, newArr...)
		}
	}
	return arr2
}
