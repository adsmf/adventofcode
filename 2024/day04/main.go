package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	p1, p2 := 0, 0
	w := 0
	for ; input[w] != '\n'; w++ {
	}
	lineLen := w + 1
	h := len(input) / lineLen
	chAt := func(row, col int) byte {
		return input[col+row*lineLen]
	}
	const p1word = "XMAS"
	p1Dirs := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for row := range h {
		for col := range w {
			switch chAt(row, col) {
			case 'X':
				for _, off := range p1Dirs {
					if col+off[0]*3 < 0 || row+off[1]*3 < 0 || col+off[0]*3 >= w || row+off[1]*3 >= h {
						continue
					}
					valid := true
					for i := 1; i < len(p1word); i++ {
						if chAt(row+off[1]*i, col+off[0]*i) != p1word[i] {
							valid = false
							break
						}
					}
					if valid {
						p1++
					}
				}
			case 'A':
				if col < 1 || col >= w-1 || row < 1 || row >= h-1 {
					continue
				}
				tl := chAt(row-1, col-1)
				tr := chAt(row-1, col+1)
				bl := chAt(row+1, col-1)
				br := chAt(row+1, col+1)
				if ((tl == 'M' && br == 'S') || (br == 'M' && tl == 'S')) &&
					((tr == 'M' && bl == 'S') || (bl == 'M' && tr == 'S')) {
					p2++
				}
			}
		}
	}
	return p1, p2
}

var benchmark = false
