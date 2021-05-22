package main

import (
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	loadSketches()
	p1, p2 := run(5, 18)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

var sketches sketchbook

func run(c1, c2 int) (int, int) {
	grid := initial()
	p1 := -1
	for i := 0; i < c2; i++ {
		if i == c1 {
			p1 = countPixels(grid)
		}
		grid = grid.enhance()
	}
	return p1, countPixels(grid)
}

func countPixels(grid pattern) int {
	pixels := 0
	for _, row := range grid {
		for _, pixel := range row {
			if pixel {
				pixels++
			}
		}
	}
	return pixels
}

func loadSketches() {
	sketches = sketchbook{}
	for _, line := range utils.ReadInputLines("input.txt") {
		parts := strings.Split(line, " => ")

		input := processPattern(parts[0])
		output := processPattern(parts[1])
		for _, inputOption := range input.transforms() {
			sketches[inputOption.hash()] = output
		}
	}
}

func processPattern(input string) pattern {
	rows := strings.Split(input, "/")
	pat := make(pattern, 0, len(rows))
	for _, row := range rows {
		rowPixels := make([]bool, 0, len(row))
		for _, pixel := range row {
			rowPixels = append(rowPixels, (pixel == '#'))
		}
		pat = append(pat, rowPixels)
	}
	return pat
}

func initial() pattern {
	return pattern{
		[]bool{false, true, false},
		[]bool{false, false, true},
		[]bool{true, true, true},
	}
}

type sketchbook map[string]pattern

var benchmark = false

type image [][]pattern

func (i image) merge() pattern {
	n := len(i[0][0])
	stitched := pattern{}
	for _, row := range i {
		for j := 0; j < n; j++ {
			stitchedRow := make([]bool, 0)
			for _, pat := range row {
				stitchedRow = append(stitchedRow, pat[j]...)
			}
			stitched = append(stitched, stitchedRow)
		}
	}
	return stitched
}

type pattern [][]bool

func (p pattern) split() image {
	n := 3
	if len(p)%2 == 0 {
		n = 2
	}
	img := image{}
	for i := 0; i*n < len(p); i++ {
		imgRow := []pattern{}
		rows := p[i*n : (i+1)*n]
		for j := 0; j < len(p); j += n {
			pat := pattern{}
			for _, row := range rows {
				pat = append(pat, row[j:j+n])
			}
			imgRow = append(imgRow, pat)
		}
		img = append(img, imgRow)
	}
	return img
}

func (p pattern) hash() string {
	return fmt.Sprintf("%v", p)
}

func (p pattern) String() string {
	sb := strings.Builder{}
	for i, row := range p {
		if i > 0 {
			sb.WriteByte('/')
		}
		for _, c := range row {
			if c {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
	}
	return sb.String()
}

func (p pattern) transforms() []pattern {
	patterns := []pattern{}
	cur := p.clone()
	for i := 0; i < 4; i++ {
		patterns = append(patterns, cur, cur.flip())
		cur = cur.rotate90()
	}
	return patterns
}

func (p pattern) enhance() pattern {
	switch len(p) {
	case 2, 3:
		return p.lookupEnhanced()
	default:
		img := p.split()
		for y, row := range img {
			for x, pat := range row {
				img[y][x] = pat.enhance()
			}
		}
		return img.merge()
	}
}

func (p pattern) lookupEnhanced() pattern {
	if enhanced, found := sketches[p.hash()]; found {
		return enhanced.clone()
	}
	panic(fmt.Errorf("Could't find pattern %v in sketches", p))
}

func (p pattern) rotate90() pattern {
	N := len(p) - 1
	copy := p.clone()
	for i, j := 0, N; i <= N/2; i, j = i+1, j-1 {
		copy[0][j], copy[i][0], copy[N][i], copy[j][N] = copy[i][0], copy[N][i], copy[j][N], copy[0][j]
	}
	return copy
}

func (p pattern) flip() pattern {
	copy := make(pattern, 0, len(p))
	for _, origRow := range p {
		newRow := make([]bool, len(origRow))
		for i, j := 0, len(origRow)-1; i < len(origRow); i, j = i+1, j-1 {
			newRow[i] = origRow[j]
		}
		copy = append(copy, newRow)
	}
	return copy
}

func (p pattern) clone() pattern {
	copy := make(pattern, 0, len(p))
	for _, row := range p {
		newRow := make([]bool, 0, len(row))
		newRow = append(newRow, row...)
		copy = append(copy, newRow)
	}
	return copy
}
