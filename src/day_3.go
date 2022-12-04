package main

import (
	"bufio"
	"fmt"
	"os"
)

func Day_3(input string) {
	letterMapPriority := map[int]int{65: 27, 97: 1} // A:27, a:1 -> https://en.wikipedia.org/wiki/List_of_Unicode_characters
	var inputSaved []string

	file, err := os.Open(input)
	CheckError(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputSaved = append(inputSaved, line)
	}

	day_3_part_1(inputSaved, letterMapPriority)
	day_3_part_2(inputSaved, letterMapPriority)
}

func day_3_part_1(inputSaved []string, letterMapPriority map[int]int) {
	totalPriority := 0
	for _, val := range inputSaved {
		var repeatedChar rune
		firstHalf := val[:len(val)/2]
		secondHalf := val[len(val)/2:]
		for _, ch1 := range firstHalf {
			for _, ch2 := range secondHalf {
				if ch1 == ch2 {
					repeatedChar = ch1
				}
			}
		}
		totalPriority = calculateTotalPriority(repeatedChar, totalPriority, letterMapPriority)
	}
	fmt.Printf("Part 1, total priority: %v.\n", totalPriority)
}

func day_3_part_2(inputSaved []string, letterMapPriority map[int]int) {
	totalPriority := 0
	for i := 0; i < len(inputSaved); i = i + 3 {
		var repeatedChar rune
		firstElf, secondElf, thirdElf := inputSaved[i], inputSaved[i+1], inputSaved[i+2]
		for _, ch1 := range firstElf {
			for _, ch2 := range secondElf {
				for _, ch3 := range thirdElf {
					if ch1 == ch2 && ch2 == ch3 {
						repeatedChar = ch1
					}
				}
			}
		}
		totalPriority = calculateTotalPriority(repeatedChar, totalPriority, letterMapPriority)
	}
	fmt.Printf("Part 2, total priority: %v.\n", totalPriority)
}

func calculateTotalPriority(runeChar rune, priority int, letterMapPriority map[int]int) int {
	var difference int
	if runeChar >= 65 && runeChar <= 90 { //uppercase
		difference = int(runeChar - 65)
		priority = priority + letterMapPriority[65] + difference
	} else if runeChar >= 97 && runeChar <= 122 { //lowercase
		difference = int(runeChar - 97)
		priority = priority + letterMapPriority[97] + difference
	}
	return priority
}
