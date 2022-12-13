package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type coord struct {
	x, y int
}

var (
	limLeft  int
	limRight int
	limUp    int
	limDown  int
)

func Day_12(input string) {
	fmt.Println("Part 1:")
	day_12_part_1(Part1, input)
	fmt.Println("Part 2:")
	day_12_part_2(Part2, input)
}

func readInputFile_12(part int, input string) ([][]rune, coord, coord, []coord) {
	file, err := os.Open(input)
	CheckError(err)
	scanner := bufio.NewScanner(file)

	runeArr := make([][]rune, 0)
	var start, end coord
	var startCoordinates []coord

	for scanner.Scan() {
		var line []rune
		for row, val := range scanner.Text() {
			if part == 1 {
				if val == 'S' {
					start = coord{row, len(runeArr)}
					val = 'a'
				}
			} else if part == 2 {
				if val == 'S' || val == 'a' {
					startCoordinates = append(startCoordinates, coord{row, len(runeArr)})
				}
			}
			if val == 'E' {
				end = coord{row, len(runeArr)}
			}
			line = append(line, val)
		}
		runeArr = append(runeArr, line)
	}
	file.Close()
	return runeArr, start, end, startCoordinates
}

func day_12_part_1(part int, input string) {
	runeArr, start, end, _ := readInputFile_12(part, input)
	findShortestPath(runeArr, part, 0, start, end)
}

func day_12_part_2(part int, input string) {
	runeArr, _, end, startCoordinates := readInputFile_12(part, input)

	var shortestPath int
	for _, start := range startCoordinates {
		shortestPath = findShortestPath(runeArr, part, shortestPath, start, end)
	}
	fmt.Println(shortestPath)
}

func findShortestPath(runeArr [][]rune, part, shortestPath int, start, end coord) int {
	visited := make(map[coord]bool)
	toVisit := []coord{start}
	distanceFromStart := map[coord]int{start: 0}
	limLeft, limRight, limUp, limDown = 0, len(runeArr[0]), 0, len(runeArr)

	for {
		if part == 2 && len(toVisit) == 0 {
			break
		}
		currentPoint := toVisit[0]
		toVisit = toVisit[1:]
		visited[currentPoint] = true

		if currentPoint == end {
			if part == 1 {
				if currentPoint == end {
					fmt.Println(distanceFromStart[end])
				}
			} else if part == 2 {
				if distanceFromStart[end] < shortestPath || shortestPath == 0 {
					shortestPath = distanceFromStart[end]
				}
			}
			break
		}

		neighbourhood := [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
		for _, neighbour := range neighbourhood {
			x, y := neighbour[1], neighbour[0]
			nextPoint := coord{currentPoint.x + x, currentPoint.y + y}
			if nextPoint.x >= limLeft && nextPoint.y >= limUp && nextPoint.x < limRight && nextPoint.y < limDown && (runeArr[nextPoint.y][nextPoint.x]-runeArr[currentPoint.y][currentPoint.x] <= 1) && !visited[nextPoint] {
				if distanceFromStart[nextPoint] == 0 {
					toVisit = append(toVisit, nextPoint)
					distanceFromStart[nextPoint] = distanceFromStart[currentPoint] + 1
				}
				if distanceFromStart[nextPoint] >= distanceFromStart[currentPoint]+1 {
					distanceFromStart[nextPoint] = distanceFromStart[currentPoint] + 1
				}
			}
		}
		sort.Slice(toVisit, func(i, j int) bool {
			return distanceFromStart[toVisit[i]] < distanceFromStart[toVisit[j]]
		})
	}
	return shortestPath
}
