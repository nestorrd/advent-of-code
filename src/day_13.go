package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Day_13(input string) {
	var inputLines []string
	file, err := os.Open(input)
	CheckError(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputLines = append(inputLines, line)
	}
	file.Close()

	fmt.Println("Part 1:")
	day_13_part_1(Part1, inputLines)
	fmt.Println("Part 2:")
	day_13_part_2(Part2, inputLines)
}

func day_13_part_1(part int, inputLines []string) {
	pairs := parseInputLines(part, inputLines)
	valid := 0
	for i, pair := range pairs {
		sub := compare(pair[0], pair[1])
		if sub == 0 || sub == 1 {
			valid += i + 1
		}
	}
	fmt.Println(valid)
}

func day_13_part_2(part int, inputLines []string) {
	pairs := parseInputLines(part, inputLines)
	allPairs := make([]any, 0, len(pairs)*2)
	for _, pair := range pairs {
		allPairs = append(allPairs, pair...)
	}
	sort.Slice(allPairs, func(first, second int) bool {
		return compare(allPairs[first], allPairs[second]) == 1
	})
	decoderKey := 1 // first packet

	for i := 0; i < len(allPairs); i++ {
		packetString := fmt.Sprintf("%v", allPairs[i])
		if packetString == "[[[2]]]" || packetString == "[[[6]]]" { // dividers
			decoderKey = decoderKey * (i + 1)
		}

	}
	fmt.Println(decoderKey)
}

func parseInputLines(part int, inputLines []string) [][]any {
	if part == Part2 {
		inputLines = append(inputLines, []string{"[[2]]", "[[6]]"}...)
	}
	pairs := make([][]any, 0)
	pair := make([]any, 0, 2)
	for _, line := range inputLines {
		if line == "" {
			continue
		}
		found, _ := iterateLine(line)
		pair = append(pair, found)
		if len(pair)%2 == 0 {
			pairs = append(pairs, pair)
			pair = make([]any, 0, 2)
		}
	}
	return pairs
}

func iterateLine(ln string) (any, int) {
	runeArr := []rune(ln)
	intRuneArr := make([]rune, 0)
	lineIterated := make([]any, 0)
	i := 0
	for len(runeArr) > i {
		switch runeArr[i] {
		case '[':
			object, nextSliceIndex := iterateLine(string(runeArr[i+1:]))
			lineIterated = append(lineIterated, object)
			i = i + nextSliceIndex + 1
		case ',':
			if len(intRuneArr) > 0 {
				lineIterated = appendStringToIntIntoSlice(intRuneArr, lineIterated)
			}
			i++
		case ']':
			if len(intRuneArr) > 0 {
				lineIterated = appendStringToIntIntoSlice(intRuneArr, lineIterated)
			}
			i++
			return lineIterated, i
		default:
			intRuneArr = append(intRuneArr, runeArr[i])
			i++
		}
	}
	return lineIterated, i
}

func compare(first any, second any) int {
	defaultReturn := 0

	firstVal, firstIsInt := first.(int)
	secondVal, secondIsInt := second.(int)

	// Check if Integers
	if firstIsInt && secondIsInt {
		// Then compare Integers
		defaultReturn := compareIntegers(firstVal, secondVal)
		return defaultReturn
	}

	firstListVal, firstIsList := first.([]any)
	secondListVal, secondIsList := second.([]any)

	// Check if one or the other are different types (int and slice)
	if !firstIsList {
		firstListVal = []any{firstVal}
	}
	if !secondIsList {
		secondListVal = []any{secondVal}
	}
	defaultReturn = compareSlices(firstListVal, secondListVal)

	return defaultReturn
}

func compareIntegers(first, second int) int {
	if first > second {
		return -1
	} else if first < second {
		return 1
	}
	return 0
}

func compareSlices(first, second []any) int {
	maxLen := len(first)
	if len(first) < len(second) {
		maxLen = len(second)
	}
	for i := 0; i < maxLen; i++ {
		if i >= len(first) {
			return 1
		}
		if i >= len(second) {
			return -1
		}
		sub := compare(first[i], second[i])
		if sub != 0 {
			return sub
		}
	}
	return 0
}

func appendStringToIntIntoSlice(intRuneArr []rune, lineIterated []any) []any {
	num, _ := strconv.Atoi(string(intRuneArr))
	lineIterated = append(lineIterated, num)
	return lineIterated
}
