package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	curPosP1, curPosP2 := point{0, 0}, point{0, 0}
	areaP1, areaP2 := 0, 0
	boundaryLenP1, boundaryLenP2 := 0, 0
	utils.EachLine(input, func(index int, line string) (done bool) {
		spaces := 0
		dirP1 := 0
		digLenP1, digLenP2 := 0, 0
		for _, ch := range line {
			if ch == ' ' {
				spaces++
				continue
			}
			switch spaces {
			case 0:
				dirP1 = int(directionLetters[ch])
				continue
			case 1:
				digLenP1 = digLenP1*10 + int(ch-'0')
			case 2:
				switch ch {
				case '(', ')', '#':
				case 'a', 'b', 'c', 'd', 'e', 'f':
					digLenP2 = digLenP2*16 + 10 + int(ch-'a')
				case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
					digLenP2 = digLenP2*16 + int(ch-'0')
				}
			}
		}
		dirP2 := digLenP2 & 0xf
		digLenP2 >>= 4
		offsetP1 := directionOffsets[dirP1]
		offsetP2 := directionOffsets[dirP2]
		nextP1 := curPosP1.add(offsetP1.times(int(digLenP1)))
		nextP2 := curPosP2.add(offsetP2.times(int(digLenP2)))
		boundaryLenP1 += int(digLenP1)
		boundaryLenP2 += int(digLenP2)
		areaP1 += (curPosP1.y + nextP1.y) * (curPosP1.x - nextP1.x)
		areaP2 += (curPosP2.y + nextP2.y) * (curPosP2.x - nextP2.x)
		curPosP1, curPosP2 = nextP1, nextP2
		return false
	})
	if areaP1 < 0 {
		areaP1 *= -1
	}

	if areaP2 < 0 {
		areaP2 *= -1
	}
	dugP1 := areaP1/2 + boundaryLenP1/2 + 1
	dugP2 := areaP2/2 + boundaryLenP2/2 + 1
	return dugP1, dugP2
}

type point struct{ x, y int }

func (p point) add(a point) point      { return point{x: p.x + a.x, y: p.y + a.y} }
func (p point) times(scalar int) point { return point{x: p.x * scalar, y: p.y * scalar} }

type direction int

const (
	dirRight direction = iota
	dirDown
	dirLeft
	dirUp

	dirMAX
	dirNone = -1
)

var directionOffsets = [dirMAX]point{
	dirUp:    {0, -1},
	dirRight: {1, 0},
	dirDown:  {0, 1},
	dirLeft:  {-1, 0},
}

var directionLetters = [...]direction{
	'L': dirLeft,
	'D': dirDown,
	'U': dirUp,
	'R': dirRight,
}

var benchmark = false
