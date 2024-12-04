package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	count := 0
	grid := utils.GetLines(input)
	h, w := len(grid), len(grid[0])
	word := "XMAS"

	checkFrom := func(col, row int) {
		for offRow := -1; offRow <= 1; offRow++ {
			for offCol := -1; offCol <= 1; offCol++ {
				if offRow == 0 && offCol == 0 {
					continue
				}
				valid := true
				for i := 1; i < len(word); i++ {
					checkCol := col + offCol*i
					checkRow := row + offRow*i
					if checkCol < 0 || checkRow < 0 || checkRow >= h || checkCol >= w {
						valid = false
						break
					}
					if grid[checkRow][checkCol] != word[i] {
						valid = false
						break
					}
				}
				if valid {
					count++
				}
			}
		}
	}
	for row, line := range grid {
		for col, ch := range line {
			if ch != 'X' {
				continue
			}
			checkFrom(col, row)
		}
	}
	return count
}

func part2() int {
	count := 0
	grid := utils.GetLines(input)
	h, w := len(grid), len(grid[0])

	checkFrom := func(col, row int) {
		if col < 1 || col >= w-1 || row < 1 || row >= h-1 {
			return
		}

		diag1 := (grid[row-1][col-1] == 'M' && grid[row+1][col+1] == 'S') ||
			(grid[row+1][col+1] == 'M' && grid[row-1][col-1] == 'S')
		diag2 := (grid[row+1][col-1] == 'M' && grid[row-1][col+1] == 'S') ||
			(grid[row-1][col+1] == 'M' && grid[row+1][col-1] == 'S')
		if diag1 && diag2 {
			count++
		}
	}
	for row, line := range grid {
		for col, ch := range line {
			if ch != 'A' {
				continue
			}
			checkFrom(col, row)
		}
	}
	return count
}

var benchmark = false
