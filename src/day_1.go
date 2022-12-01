package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	id       int
	calories int
}

// https://adventofcode.com/2022/day/1
func Day_1(inputFile string) {

	file, err := os.Open(inputFile)
	CheckError(err)

	elf_counter, elf_calories, top_three_calories := 0, 0, 0
	var elfs []Elf

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			strConverted, err := strconv.Atoi(scanner.Text())
			CheckError(err)
			elf_calories = elf_calories + strConverted
		} else {
			elf_counter++
			elfs = append(elfs, Elf{id: elf_counter, calories: elf_calories})
			elf_calories = 0
		}
	}
	if elf_calories > 0 {
		elf_counter++
		elfs = append(elfs, Elf{id: elf_counter, calories: elf_calories})
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}

	sort.SliceStable(elfs, func(i, j int) bool {
		return elfs[i].calories < elfs[j].calories
	})
	top_three_calories = elfs[len(elfs)-1].calories + elfs[len(elfs)-2].calories + elfs[len(elfs)-3].calories
	fmt.Printf("The Elf with more calories has %v calories.\n", elfs[len(elfs)-1].calories)
	fmt.Printf("Top 3 elfs have a total amount of %v calories.\n", top_three_calories)
}
