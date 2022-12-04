package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day_4(input string) {
	counter1, counter2 := 0, 1000
	file, err := os.Open(input)
	CheckError(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		firstElf, secondElf, _ := strings.Cut(line, ",")

		var firstElfNumbers [2]int
		var seconfElfNumbers [2]int

		firstElfNumbers = transformToInteger(firstElf)
		seconfElfNumbers = transformToInteger(secondElf)

		counter1 = day_4_part_1(line, counter1, firstElfNumbers, seconfElfNumbers)
		counter2 = day_4_part_2(line, counter2, firstElfNumbers, seconfElfNumbers)
	}
	fmt.Printf("Part1 number obtained: %v\n", counter1)
	fmt.Printf("Part2 number obtained: %v\n", counter2)
}

func day_4_part_1(line string, counter int, firstElfNumbers, seconfElfNumbers [2]int) int {
	if (firstElfNumbers[0] >= seconfElfNumbers[0] && firstElfNumbers[1] <= seconfElfNumbers[1]) ||
		(firstElfNumbers[0] <= seconfElfNumbers[0] && firstElfNumbers[1] >= seconfElfNumbers[1]) {
		counter++
		fmt.Printf("%v\n", line)
	}
	return counter
}

func day_4_part_2(line string, counter int, firstElfNumbers, seconfElfNumbers [2]int) int {
	if (firstElfNumbers[1] < seconfElfNumbers[0]) ||
		(firstElfNumbers[0] > seconfElfNumbers[1]) {
		counter--
	}
	return counter
}

func transformToInteger(first string) [2]int {
	var intArray [2]int
	var middleStrings [2]string
	middleStrings[0], middleStrings[1], _ = strings.Cut(first, "-")
	intArray[0], _ = strconv.Atoi(middleStrings[0])
	intArray[1], _ = strconv.Atoi(middleStrings[1])
	return intArray
}
