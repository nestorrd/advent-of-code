package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Lose = 0
	Draw = 3
	Win  = 6
)

type competition struct {
	opponentPoints int
	yourPoints     int
	whoDefeatsWho  map[string]string
	opponentPlay   map[string]string
	yourPlay       map[string]string
}

// https://adventofcode.com/2022/day/2
func Day_2(input string) {
	whoDefeatsWho := map[string]string{
		"Rock":     "Scissors",
		"Scissors": "Paper",
		"Paper":    "Rock"}

	opponentPlay := map[string]string{"A": "Rock", "B": "Paper", "C": "Scissors"}

	part1 := competition{
		opponentPoints: 0,
		yourPoints:     0,
		whoDefeatsWho:  whoDefeatsWho,
		opponentPlay:   opponentPlay,
		yourPlay:       map[string]string{"X": "Rock", "Y": "Paper", "Z": "Scissors"}}

	part2 := competition{
		opponentPoints: 0,
		yourPoints:     0,
		whoDefeatsWho:  whoDefeatsWho,
		opponentPlay:   opponentPlay,
		yourPlay:       map[string]string{"X": "Lose", "Y": "Draw", "Z": "Win"}}

	file, err := os.Open(input)
	CheckError(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		part1 = day_2_part_1(line, part1)
		part2 = day_2_part_2(line, part2)
	}
	fmt.Printf("Part 1. Your results: %v | Opponent results: %v\n", part1.yourPoints, part1.opponentPoints)
	fmt.Printf("Part 2. Your results: %v | Opponent results: %v\n", part2.yourPoints, part2.opponentPoints)
	file.Close()
}

func day_2_part_1(line string, comp competition) competition {
	var yourChoice, opponentChoice string
	if v, ok := comp.opponentPlay[line[0:1]]; ok {
		opponentChoice = v
	}
	if v, ok := comp.yourPlay[line[len(line)-1:]]; ok {
		yourChoice = v
	}
	if comp.whoDefeatsWho[opponentChoice] == yourChoice {
		comp.opponentPoints, comp.yourPoints = addResult("Lose", opponentChoice, yourChoice, comp.opponentPoints, comp.yourPoints)
	} else if opponentChoice == yourChoice {
		comp.opponentPoints, comp.yourPoints = addResult("Draw", opponentChoice, yourChoice, comp.opponentPoints, comp.yourPoints)
	} else {
		comp.opponentPoints, comp.yourPoints = addResult("Win", opponentChoice, yourChoice, comp.opponentPoints, comp.yourPoints)
	}
	return comp
}

func day_2_part_2(line string, comp competition) competition {
	var expectedResult, yourChoice, opponentChoice string

	if v, ok := comp.opponentPlay[line[0:1]]; ok {
		opponentChoice = v
	}
	if v, ok := comp.yourPlay[line[len(line)-1:]]; ok {
		expectedResult = v
	}
	switch expectedResult {
	case "Lose":
		yourChoice = comp.whoDefeatsWho[opponentChoice]
		comp.opponentPoints, comp.yourPoints = addResult(expectedResult, opponentChoice, yourChoice, comp.opponentPoints, comp.yourPoints)
	case "Draw":
		yourChoice = opponentChoice
		comp.opponentPoints, comp.yourPoints = addResult(expectedResult, opponentChoice, yourChoice, comp.opponentPoints, comp.yourPoints)
	case "Win":
		for key, value := range comp.whoDefeatsWho {
			if value == opponentChoice {
				yourChoice = key
			}
		}
		comp.opponentPoints, comp.yourPoints = addResult(expectedResult, opponentChoice, yourChoice, comp.opponentPoints, comp.yourPoints)
	}
	return comp
}

func addResult(result, opponentChoice, yourChoice string, opponentPoints, yourPoints int) (int, int) {
	choicePuntuation := map[string]int{"Rock": 1, "Paper": 2, "Scissors": 3}
	switch result {
	case "Lose":
		opponentPoints = opponentPoints + choicePuntuation[opponentChoice] + Win
		yourPoints = yourPoints + choicePuntuation[yourChoice]
	case "Draw":
		opponentPoints = opponentPoints + choicePuntuation[opponentChoice] + Draw
		yourPoints = yourPoints + choicePuntuation[yourChoice] + Draw
	case "Win":
		opponentPoints = opponentPoints + choicePuntuation[opponentChoice]
		yourPoints = yourPoints + choicePuntuation[yourChoice] + Win
	}
	return opponentPoints, yourPoints
}
