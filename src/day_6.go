package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	marker = 14 // change 14 for 4 to trigger part 1
)

func Day_6(input string) {
	file, err := os.Open(input)
	CheckError(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		day_6(line)
	}
}

func day_6(line string) {
	itemsAddedToArray, indexWithFirstData := 0, 0
	var characters [marker]rune
	arrayIsDifferent := false
	for charPosition, val := range line {
		if charRepeated(characters, val) {
			if itemsAddedToArray < marker-1 {
				characters[itemsAddedToArray] = val
				itemsAddedToArray++
				characters, itemsAddedToArray = iterateOverArray(characters, val)
			} else {
				characters[itemsAddedToArray] = val
				characters, itemsAddedToArray = iterateOverArray(characters, val)
			}
			fmt.Println(characters)
		} else {
			if itemsAddedToArray < marker-1 {
				characters[itemsAddedToArray] = val
			} else if itemsAddedToArray == marker-1 {
				characters[itemsAddedToArray] = val
				arrayIsDifferent = arrayIsTotallyDiferent(characters)
			}

			if (itemsAddedToArray == marker-1) && (arrayIsDifferent) {
				indexWithFirstData = charPosition + 1
				break
			}
			if itemsAddedToArray < marker-1 {
				itemsAddedToArray++
			}
			fmt.Println(characters)
		}
		indexWithFirstData = charPosition + 1
	}
	fmt.Printf("First character in %v\n", indexWithFirstData)
}

func charRepeated(arr [marker]rune, char rune) bool {
	for _, val := range arr {
		if val == char {
			return true
		}
	}
	return false
}

func iterateOverArray(arr [marker]rune, char rune) ([marker]rune, int) {
	var newArray [marker]rune
	newElements := false
	elementsInNewArray := 0
	for _, val := range arr {
		if newElements {
			newArray[elementsInNewArray] = val
			if val != 0 {
				elementsInNewArray++
			}
		}
		if val == char {
			newElements = true
		}
	}
	return newArray, elementsInNewArray
}

func arrayIsTotallyDiferent(arr [marker]rune) bool {
	different := true
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if i == j {
				return !different
			}
		}
	}
	return different
}
