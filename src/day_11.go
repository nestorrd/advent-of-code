package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	id          int
	items       []int64
	operation   string
	divisible   int64
	ifTrue      int
	ifFalse     int
	inspections int
}

func Day_11(input string) {
	getMonkeyBusiness(Part1, 20, readInputFile(input))
	getMonkeyBusiness(Part2, 10000, readInputFile(input))
}

func readInputFile(input string) []monkey {
	var monkeys []monkey
	currentMonkey := 0

	file, err := os.Open(input)
	CheckError(err)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Monkey") {
			monkeys = append(monkeys, monkey{id: currentMonkey, operation: "", divisible: 0, ifTrue: 0, ifFalse: 0, inspections: 0})
		}
		if strings.HasPrefix(line, "  Starting items: ") {
			line = strings.ReplaceAll(line, "  Starting items: ", "")
			lineArr := strings.Split(line, ", ")
			for _, v := range lineArr {
				item, _ := strconv.Atoi(v)
				monkeys[currentMonkey].items = append(monkeys[currentMonkey].items, int64(item))
			}
		}
		if strings.HasPrefix(line, "  Operation: ") {
			line = strings.ReplaceAll(line, "  Operation: ", "")
			monkeys[currentMonkey].operation = line
		}
		if strings.HasPrefix(line, "  Test: divisible by ") {
			line = strings.ReplaceAll(line, "  Test: divisible by ", "")
			monkeys[currentMonkey].divisible, _ = strconv.ParseInt(line, 0, 64)
		}
		if strings.HasPrefix(line, "    If true: throw to monkey ") {
			line = strings.ReplaceAll(line, "    If true: throw to monkey ", "")
			monkeyId, _ := strconv.Atoi(line)
			monkeys[currentMonkey].ifTrue = monkeyId
		}
		if strings.HasPrefix(line, "    If false: throw to monkey ") {
			line = strings.ReplaceAll(line, "    If false: throw to monkey ", "")
			monkeyId, _ := strconv.Atoi(line)
			monkeys[currentMonkey].ifFalse = monkeyId
			currentMonkey++
		}
	}
	file.Close()
	return monkeys
}

func getMonkeyBusiness(part int, maxRounds int, monkeys []monkey) {

	for i := 0; i < maxRounds; i++ {
		for key, monkey := range monkeys {
			for _, item := range monkey.items {

				newWorryLvl := performOperation(monkey.operation, item)
				switch part {
				case 1:
					newWorryLvl = newWorryLvl / 3
				case 2:
					newWorryLvl = newWorryLvl % getMinimumCommonMultiple(monkeys)
				}

				isDivisible := checkIfDivisible(newWorryLvl, monkey.divisible)
				if isDivisible {
					monkeys[monkey.ifTrue].items = append(monkeys[monkey.ifTrue].items, newWorryLvl)
				} else {
					monkeys[monkey.ifFalse].items = append(monkeys[monkey.ifFalse].items, newWorryLvl)
				}
				monkeys[key].inspections++
			}
			monkeys[key].items = []int64{}
		}
	}
	var allInsp []int
	for _, monkey := range monkeys {
		allInsp = append(allInsp, monkey.inspections)
	}
	sort.Ints(allInsp)
	fmt.Printf("Part %v, Monkey Business: %v\n", part, allInsp[len(allInsp)-1]*allInsp[len(allInsp)-2])
}

func performOperation(operation string, item int64) int64 {
	operationParsed := strings.Fields(operation)
	var new, old int64
	var operator int64
	old = item
	if operationParsed[4] != "old" {
		operator, _ = (strconv.ParseInt(operationParsed[4], 0, 64))
	} else {
		operator = old
	}
	switch operationParsed[3] {
	case "+":
		new = old + operator
	case "*":
		new = old * operator
	}

	return new
}

func checkIfDivisible(worryLvl, divisor int64) bool {
	isDivisible := false
	if worryLvl%divisor == 0 {
		isDivisible = true
	}
	return isDivisible
}

func getMinimumCommonMultiple(monkeys []monkey) int64 {
	mcm := int64(1)
	for _, monkey := range monkeys {
		mcm = int64(monkey.divisible) * mcm
	}
	return mcm
}
