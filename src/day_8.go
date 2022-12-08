package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	height = 99
	width  = 99
)

func Day_8(input string) {
	var inputArr [height][width]int
	file, err := os.Open(input)
	CheckError(err)
	scanner := bufio.NewScanner(file)
	lineCounter, visibleCounter, treePunct := 0, (height*2)+(width*2)-4, 0
	for scanner.Scan() {
		line := scanner.Text()
		for key, val := range line {
			num, _ := strconv.Atoi(string(val))
			inputArr[lineCounter][key] = num
		}
		lineCounter++
	}

	for i := 1; i < len(inputArr)-1; i++ {
		for j := 1; j < (len(inputArr[i]))-1; j++ {
			if isVisible(j, i, inputArr) {
				visibleCounter++
			}
			pnts := calculateTreeWithBetterViews(j, i, inputArr)
			if pnts > treePunct {
				fmt.Printf("New tree with better views [%v][%v]=%v with %v points\n", i, j, inputArr[j][i], pnts)
				treePunct = pnts
			}
		}
	}
	fmt.Println(visibleCounter)
	fmt.Println(treePunct)
}

func calculateTreeWithBetterViews(x, y int, arr [height][width]int) int {
	cl, cr, ct, cb := x-1, x+1, y-1, y+1
	tree, nextTree, pnts, tl, tr, tt, tb := arr[y][x], 0, 0, 0, 0, 0, 0

	for {
		if cl >= 0 {
			nextTree = arr[y][cl]
			if tree > nextTree {
				tl++
				cl--
				if cl < 0 {
					break
				}
			} else if tree <= nextTree {
				tl++
				break
			}

		} else {
			break
		}
	}
	for {
		if cr < width {
			nextTree = arr[y][cr]
			if tree > nextTree {
				tr++
				cr++
				if cr == width {
					break
				}

			} else if tree <= nextTree {
				tr++
				break
			}
		} else {
			break
		}
	}

	for {
		if ct >= 0 {
			nextTree = arr[ct][x]
			if tree > nextTree {
				tt++
				ct--
				if ct < 0 {
					break
				}
			} else if tree <= nextTree {
				tt++
				break
			}
		} else {
			break
		}
	}

	for {
		if cb < height {
			nextTree = arr[cb][x]
			if tree > nextTree {
				tb++
				cb++
				if cb < 0 {
					break
				}

			} else if tree <= nextTree {
				tb++
				break
			}
		} else {
			break
		}
	}

	pnts = tl * tr * tt * tb
	fmt.Printf("[%v][%v]=%v %v * %v * %v * %v = %v\n", y, x, arr[y][x], tl, tr, tt, tb, pnts)
	return pnts

}

func isVisible(x, y int, arr [height][width]int) bool {
	vl, vr, vt, vb := true, true, true, true
	tree, nextTree, cl, cr, ct, cb := arr[y][x], 0, x-1, x+1, y-1, y+1

	for {
		if cl >= 0 {
			nextTree = arr[y][cl]
			if tree > nextTree {
				cl--
				if cl < 0 {
					break
				}
			} else if tree <= nextTree {
				vl = false
				break
			}
		} else {
			break
		}
	}
	if vl {
		return vl
	}

	for {
		if cr < width {
			nextTree = arr[y][cr]
			if tree > nextTree {
				cr++
				if cr == width {
					break
				}
			} else if tree <= nextTree {
				vr = false
				break
			}
		} else {
			break
		}
	}
	if vr {
		return vr
	}

	for {
		if ct >= 0 {
			nextTree = arr[ct][x]
			if tree > nextTree {
				ct--
				if ct < 0 {
					break
				}
			} else if tree <= nextTree {
				vt = false
				break
			}
		} else {
			break
		}
	}
	if vt {
		return vt
	}

	for {
		if cb < height {
			nextTree = arr[cb][x]
			if tree > nextTree {
				cb++
				if cb == height {
					break
				}
			} else if tree <= nextTree {
				vb = false
				break
			}
		} else {
			break
		}
	}
	if vb {
		return vb
	}
	return false
}
