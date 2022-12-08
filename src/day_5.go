package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	part1 = 1
	part2 = 2
)

func Day_5(input string) {
	var moves []string
	stacksSaved := false

	file, err := os.Open(input)
	CheckError(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if stacksSaved {
			moves = append(moves, line)
		} else {
			if line == "" {
				stacksSaved = true
			}
		}
	}

	day_5_executePart(part1, moves)
	day_5_executePart(part2, moves)
}

func day_5_executePart(part int, moves []string) {
	stacks := [][]string{{"R", "N", "P", "G"},
		{"T", "J", "B", "L", "C", "S", "V", "H"},
		{"T", "D", "B", "M", "N", "L"},
		{"R", "V", "P", "S", "B"},
		{"G", "C", "Q", "S", "W", "M", "V", "H"},
		{"W", "Q", "S", "C", "D", "B", "J"},
		{"F", "Q", "L"},
		{"W", "M", "H", "T", "D", "L", "F", "V"},
		{"L", "P", "B", "V", "M", "J", "F"}}

	for _, newString := range moves {
		strArr := strings.Split(newString, " ")
		move, _ := strconv.Atoi(strArr[1])
		from, _ := strconv.Atoi(strArr[3])
		to, _ := strconv.Atoi(strArr[5])
		from, to = from-1, to-1
		var iterations int

		if part == 1 {
			iterations = move
			move = 1
		} else if part == 2 {
			iterations = 1
		}

		for i := 0; i < iterations; i++ {
			itm := stacks[from][len(stacks[from])-move:]
			stacks[from] = stacks[from][:len(stacks[from])-move]
			stacks[to] = append(stacks[to], itm...)
		}
	}

	var a string
	for k, _ := range stacks {
		a = a + stacks[k][len(stacks[k])-1]
	}
	fmt.Printf("Part %v: %v\n", part, a)
}
