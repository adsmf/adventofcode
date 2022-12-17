package main

import (
	_ "embed"
	"fmt"
	"strings"
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
	c := chamber{
		rows: []byte{},
	}
	curShape := 0
	inputPos := 0
	countTo := 1000000000000
	const maxHeight = 3000
	seen := make(map[gridHash]int, maxHeight)
	heights := [maxHeight]int{}
	cyclesAdd := 0
	for pieceNum := 0; pieceNum < countTo; pieceNum++ {
		inputPos = c.dropPiece(curShape, inputPos)
		curShape++
		if curShape == 5 {
			curShape = 0
		}
		if cyclesAdd == 0 {
			heights[pieceNum] = c.curTop
			if c.curTop > _hash_lookback {
				hash := c.fnvHash(inputPos, curShape)
				if seen[hash] > 0 {
					seenAtPiece := seen[hash]
					cycleLen := pieceNum - seenAtPiece
					numCycles := (countTo - pieceNum) / cycleLen
					cycleHeight := (heights[pieceNum] - heights[seenAtPiece])
					cyclesAdd = numCycles * cycleHeight
					countTo = (countTo-pieceNum)%cycleLen + pieceNum
				}
				seen[hash] = pieceNum
			}
		}
	}
	return heights[2021], c.curTop + cyclesAdd
}

type gridHash uint32

const _hash_lookback = 8

const (
	fnvp32 uint32 = 0x01000193
	fnvo32 uint32 = 0x811c9dc5
)

type chamber struct {
	rows                []byte
	falling             []point
	fallLeft, fallRight int
	curTop              int
}

func (c chamber) fnvHash(inputPos int, curShape int) gridHash {
	hash := fnvo32
	for i := 0; i < _hash_lookback; i += 4 {
		y := c.curTop - i - 1
		hash ^= uint32(c.rows[y])<<24 |
			uint32(c.rows[y-1])<<16 |
			uint32(c.rows[y-2])<<8 |
			uint32(c.rows[y-2])
		hash *= fnvp32
	}
	hash ^= uint32(inputPos)<<8 | uint32(curShape)
	hash *= fnvp32
	return gridHash(hash)
}

func (c *chamber) dropPiece(shape int, inputPos int) int {
	c.falling = []point{}
	for _, tile := range rockShapes[shape] {
		c.falling = append(c.falling, point{tile.x + 2, tile.y + c.curTop + 3})
	}
	c.fallLeft = 2
	c.fallRight = 2 + rockWidths[shape]
	inputsUsed := 0
	for ; ; inputsUsed++ {
		inputPos = c.pushPiece(inputPos)
		done := c.stepPieceDown()
		if done {
			for _, tile := range c.falling {
				if tile.y >= c.curTop {
					c.curTop = tile.y + 1
				}
				for tile.y >= len(c.rows) {
					c.rows = append(c.rows, 0)
				}
				c.rows[tile.y] |= 1 << tile.x
			}
			break
		}
	}
	return inputPos
}

func (c *chamber) pushPiece(inputPos int) int {
	ch := input[inputPos]
	inputPos++
	if input[inputPos] == '\n' {
		inputPos = 0
	}
	switch ch {
	case '<':
		if c.fallLeft == 0 {
			return inputPos
		}
		for _, tile := range c.falling {
			if tile.y < len(c.rows) && c.rows[tile.y]&(1<<(tile.x-1)) > 0 {
				return inputPos
			}
		}
		c.fallLeft--
		c.fallRight--
		for idx := range c.falling {
			c.falling[idx].x--
		}
	case '>':
		if c.fallRight == 7 {
			return inputPos
		}
		for _, tile := range c.falling {
			if tile.y < len(c.rows) && c.rows[tile.y]&(1<<(tile.x+1)) > 0 {
				return inputPos
			}
		}
		c.fallLeft++
		c.fallRight++
		for idx := range c.falling {
			c.falling[idx].x++
		}
	}
	return inputPos
}

func (c *chamber) stepPieceDown() bool {
	for _, tile := range c.falling {
		if tile.y == 0 {
			return true
		}
		if tile.y <= len(c.rows) {
			if c.rows[tile.y-1]&(1<<(tile.x)) > 0 {
				return true
			}
		}
	}
	for idx := range c.falling {
		c.falling[idx].y--
	}
	return false
}

func (c chamber) String() string {
	sb := strings.Builder{}

	for y := c.curTop + 6; y >= 0; y-- {
		sb.WriteByte('|')
		rowBytes := [7]byte{'.', '.', '.', '.', '.', '.', '.'}
		if y < c.curTop {
			for x := 0; x < 7; x++ {
				if c.rows[y]&(1<<x) > 0 {
					rowBytes[x] = '#'
				}
			}
		}
		if c.falling != nil {
			for _, fallPos := range c.falling {
				if fallPos.y == y {
					rowBytes[fallPos.x] = '@'
				}
			}
		}
		for x := 0; x < 7; x++ {
			sb.WriteByte(rowBytes[x])
		}
		sb.WriteByte('|')
		sb.WriteByte('\n')
	}
	sb.WriteString("+-------+\n")

	return sb.String()
}

type point struct{ x, y int }
type rockPiece []point

var rockShapes = [...]rockPiece{
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}},         // H-Line
	{{0, 1}, {1, 0}, {2, 1}, {1, 2}, {1, 1}}, // Plus
	{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}, // Backwards L
	{{0, 0}, {0, 1}, {0, 2}, {0, 3}},         // V-Line
	{{0, 0}, {0, 1}, {1, 0}, {1, 1}},         //Block
}
var rockWidths = [...]int{4, 3, 3, 1, 2}

var benchmark = false
