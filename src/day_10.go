package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type signal struct {
	strength int
	cycle    int
}

type screen struct {
	allLines []string
	line     string
	cursor   int
}

func Day_10(input string) {
	file, err := os.Open(input)
	CheckError(err)
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	day_10_part_1(lines)
	day_10_part_2(lines)

}
func day_10_part_1(lines []string) {
	register := 1
	signal := signal{strength: 0, cycle: 1}
	for _, line := range lines {
		if strings.HasPrefix(line, "noop") {
			signal.cycle, signal.strength = calculateSignalCycle(signal.cycle, signal.strength, register)
		} else if strings.HasPrefix(line, "addx") {
			signal.cycle, signal.strength = calculateSignalCycle(signal.cycle, signal.strength, register)
			signal.cycle, signal.strength = calculateSignalCycle(signal.cycle, signal.strength, register)

			register += increaseRegisterValue(line, register)
		}
	}
	fmt.Println("Part 1:")
	fmt.Println(signal.strength)
}
func day_10_part_2(lines []string) {
	register := 1
	screen := screen{allLines: make([]string, 0), cursor: 0}
	for _, line := range lines {
		if strings.HasPrefix(line, "noop") {
			screen.allLines, screen.line, screen.cursor = calculateScreenCycle(screen.allLines, screen.line, screen.cursor, register)
		} else if strings.HasPrefix(line, "addx") {
			screen.allLines, screen.line, screen.cursor = calculateScreenCycle(screen.allLines, screen.line, screen.cursor, register)
			screen.allLines, screen.line, screen.cursor = calculateScreenCycle(screen.allLines, screen.line, screen.cursor, register)

			register += increaseRegisterValue(line, register)
		}
	}
	fmt.Println("Part 2:")
	fmt.Println(strings.Join(screen.allLines, "\n"))
	fmt.Println("´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´´")
}

func increaseRegisterValue(line string, register int) int {
	lineArr := strings.Split(line, " ")
	val, _ := strconv.Atoi(lineArr[1])
	return val
}

func calculateSignalCycle(cycle, strenght, register int) (int, int) {
	if cycle%40 == 20 {
		strenght += cycle * register
	}
	cycle++
	return cycle, strenght
}

func calculateScreenCycle(lines []string, line string, cursor, register int) ([]string, string, int) {
	if cursor <= register+1 && cursor >= register-1 {
		line += "´´"
	} else {
		line += "[["
	}
	cursor++
	if cursor >= 40 {
		lines = append(lines, line)
		line = ""
		cursor = 0
	}
	return lines, line, cursor
}
