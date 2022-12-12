package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 2536
type knots struct {
	id               int
	positionsVisited map[string]int
	head             bool
	x                int
	y                int
}

func Day_9(input string) {
	var inputFile []string
	file, err := os.Open(input)
	CheckError(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputFile = append(inputFile, scanner.Text())
	}

	day_9_part_1(inputFile)
	//day_9_part_2(inputFile)

}

func day_9_part_1(inputFile []string) {

	head := knots{positionsVisited: make(map[string]int), x: 0, y: 0}
	tail := knots{positionsVisited: make(map[string]int), x: 0, y: 0}

	for _, val := range inputFile {
		line := strings.Split(val, " ")
		direction := line[0]
		moves, _ := strconv.Atoi(line[1])

		for i := 0; i < moves; i++ {

			head.x, head.y, tail.x, tail.y = iter(direction, head.x, head.y, tail.x, tail.y)

			position := fmt.Sprintf("(%v|%v)", tail.x, tail.y)
			if val, ok := tail.positionsVisited[position]; ok {
				tail.positionsVisited[position] = val + 1
			} else {
				tail.positionsVisited[position] = 1
			}
		}
	}
	fmt.Println(len(tail.positionsVisited))
}

func day_9_part_2(inputFile []string) {
	var a [10]knots
	a[0] = knots{id: 0, head: true, positionsVisited: make(map[string]int), x: 0, y: 0}
	a[1] = knots{id: 1, head: false, positionsVisited: make(map[string]int), x: 0, y: 0}
	a[2] = knots{id: 2, head: false, positionsVisited: make(map[string]int), x: 0, y: 0}
	a[3] = knots{id: 3, head: false, positionsVisited: make(map[string]int), x: 0, y: 0}
	a[4] = knots{id: 4, head: false, positionsVisited: make(map[string]int), x: 0, y: 0}
	a[5] = knots{id: 5, head: false, positionsVisited: make(map[string]int), x: 0, y: 0}
	a[6] = knots{id: 6, head: false, positionsVisited: make(map[string]int), x: 0, y: 0}
	a[7] = knots{id: 7, head: false, positionsVisited: make(map[string]int), x: 0, y: 0}
	a[8] = knots{id: 8, head: false, positionsVisited: make(map[string]int), x: 0, y: 0}
	a[9] = knots{id: 9, head: false, positionsVisited: make(map[string]int), x: 0, y: 0}

	for _, val := range inputFile {
		line := strings.Split(val, " ")
		direction := line[0]
		moves, _ := strconv.Atoi(line[1])

		for i := 0; i < moves; i++ {
			fmt.Printf("Move: %v\n", i+1)
			for m := 0; m < len(a); m++ {

				if m == 9 {
					break
				}
				xh := a[m].x
				yh := a[m].y
				xn := a[m+1].x
				yn := a[m+1].y

				xh, yh, xn, yn = iter2(direction, xh, yh, xn, yn, a[m].head)

				a[m].x, a[m].y, a[m+1].x, a[m+1].y = xh, yh, xn, yn

				position := fmt.Sprintf("(%v|%v)", xh, yh)
				if val, ok := a[m].positionsVisited[position]; ok {
					a[m].positionsVisited[position] = val + 1
				} else {
					a[m].positionsVisited[position] = 1
				}

				position = fmt.Sprintf("(%v|%v)", xn, yn)
				if val, ok := a[m+1].positionsVisited[position]; ok {
					a[m+1].positionsVisited[position] = val + 1
				} else {
					a[m+1].positionsVisited[position] = 1
				}

			}
		}
	}
	fmt.Printf("Tail positions visited %v\n", len(a[len(a)-1].positionsVisited))
}

func iter2(direction string, xh, yh, xt, yt int, head bool) (int, int, int, int) {
	switch direction {
	case "L":
		if head {
			xh--
			if xh-xt < -1 {
				if yh-yt < 0 {
					yt--
				}
				if yh-yt > 0 {
					yt++
				}
				xt--
			}
		} else {
			if xh-xt < -1 {
				if yh-yt < 0 {
					yt--
				}
				if yh-yt > 0 {
					yt++
				}
				xt--
			} else if yh-yt < -1 {
				if xh-xt < 0 {
					xt--
				}
				if xh-xt > 0 {
					xt++
				}
				yt--
			}
		}

	case "R":
		if head {
			xh++
			if xh-xt > 1 {
				if yh-yt < 0 {
					yt--
				}
				if yh-yt > 0 {
					yt++
				}
				xt++
			}
		} else {
			if xh-xt > 1 {
				if yh-yt < 0 {
					yt--
				}
				if yh-yt > 0 {
					yt++
				}
				xt++
			} else if yh-yt > 1 {
				if xh-xt < 0 {
					xt--
				}
				if xh-xt > 0 {
					xt++
				}
				yt++
			}
		}

	case "U":
		if head {
			yh++
			if yh-yt > 1 {
				if xh-xt < 0 {
					xt--
				}
				if xh-xt > 0 {
					xt++
				}
				yt++
			}
		} else {
			if yh-yt > 1 {
				if xh-xt < 0 {
					xt--
				}
				if xh-xt > 0 {
					xt++
				}
				yt++
			} else if xh-xt > 1 {
				if yh-yt < 0 {
					yt--
				}
				if yh-yt > 0 {
					yt++
				}
				xt++
			}
		}

	case "D":
		if head {
			yh--
			if yh-yt < -1 {
				if xh-xt < 0 {
					xt--
				}
				if xh-xt > 0 {
					xt++
				}
				yt--
			}
		} else {
			if yh-yt < -1 {
				if xh-xt < 0 {
					xt--
				}
				if xh-xt > 0 {
					xt++
				}
				yt--
			} else if xh-xt < -1 {
				if yh-yt < 0 {
					yt--
				}
				if yh-yt > 0 {
					yt++
				}
				xt--
			}
		}

	}
	return xh, yh, xt, yt
}

func iter(direction string, xh, yh, xt, yt int) (int, int, int, int) {
	switch direction {
	case "L":
		xh--
		if xh-xt < -1 {
			if yh-yt < 0 {
				yt--
			}
			if yh-yt > 0 {
				yt++
			}
			xt--
		}

	case "R":
		xh++
		if xh-xt > 1 {
			if yh-yt < 0 {
				yt--
			}
			if yh-yt > 0 {
				yt++
			}
			xt++
		}
	case "U":
		yh++
		if yh-yt > 1 {
			if xh-xt < 0 {
				xt--
			}
			if xh-xt > 0 {
				xt++
			}
			yt++
		}

	case "D":
		yh--
		if yh-yt < -1 {
			if xh-xt < 0 {
				xt--
			}
			if xh-xt > 0 {
				xt++
			}
			yt--
		}
	}
	return xh, yh, xt, yt
}
