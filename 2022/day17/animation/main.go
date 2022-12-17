package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

//go:embed input.txt
var input []byte

func main() {
	c := &chamber{
		rows: []byte{},
	}
	app := tview.NewApplication()
	window := tview.NewBox()
	window.SetBorder(true).SetTitle(" Advent of Code - 2022 day 17 - @adsmf ")
	window.SetDrawFunc(c.Draw)
	go app.SetRoot(window, true).Run()
	curShape := 0
	inputPos := 0
	countTo := 1000000000000
	const maxHeight = 3000
	seen := make(map[gridHash]int, maxHeight)
	heights := [maxHeight]int{}
	cyclesAdd := 0
	c.delay = 100000000
	c.delayMult = 995
	draw := func(force bool) {
		if c.delay == 0 && !force {
			return
		}
		app.Draw()
		c.delay = c.delay * c.delayMult / 1000
		time.Sleep(time.Duration(c.delay))
	}
	for pieceNum := 0; pieceNum < countTo; pieceNum++ {
		inputPos = c.dropPiece(draw, curShape, inputPos)
		draw(true)
		curShape++
		if curShape == 5 {
			curShape = 0
		}
		if pieceNum == 2021 {
			c.thoughtText = []string{
				"2022 pieces... made it!",
			}

			draw(true)
			time.Sleep(5 * time.Second)
			c.thoughtText = []string{
				"What?!? 1000000000000!?! Uh oh...",
				"... best get started then...",
			}
			draw(true)
			time.Sleep(5 * time.Second)
			c.delay = 100000000
			c.thoughtText = []string{}
		}
		if pieceNum+c.skippedPieces == 1000000000000-101 {
			c.thoughtText = []string{
				"Nearly there....",
			}
		}
		if pieceNum+c.skippedPieces == 1000000000000-2 {
			c.thoughtText = []string{
				"Phew, done!",
			}
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
					c.thoughtText = []string{
						"Hmmm... this seems familiar...",
						fmt.Sprintf("I saw this %d pieces ago!", cycleHeight),
					}
					c.delay = 2000
					c.delayMult = 1001
					draw(true)
					time.Sleep(3 * time.Second)
					c.thoughtText = []string{
						"Let's skip ahead a little...",
					}
					draw(true)
					time.Sleep(3 * time.Second)
					c.thoughtText = []string{}
					c.skippedPieces = numCycles * cycleLen
					c.skippedHeight = cyclesAdd
				}
				seen[hash] = pieceNum
			}
		}
	}
	time.Sleep(5 * time.Second)
	app.Stop()
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
	counter             int
	delay               int
	delayMult           int
	thoughtText         []string
	skippedPieces       int
	skippedHeight       int
}

func (c *chamber) Draw(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	startY := c.curTop
	drawHeight := height - 4
	if startY < drawHeight {
		startY = drawHeight
	}
	endY := startY - drawHeight - 1
	xOffset := width/2 - 4
	yCenter := height / 2
	counterMsg := fmt.Sprintf("Pieces: %d", c.counter+c.skippedPieces)
	if c.skippedPieces > 0 {
		counterMsg += fmt.Sprintf(" (skipped %d)", c.skippedPieces)
	}
	tview.Print(screen, counterMsg, 4, 4, 99, tview.AlignLeft, tcell.ColorDefault)
	tview.Print(screen, fmt.Sprintf("Height: %d", c.curTop+c.skippedHeight), 4, 5, 99, tview.AlignLeft, tcell.ColorDefault)
	// tview.Print(screen, fmt.Sprintf("Delay: %d", c.delay), 4, 6, 99, tview.AlignLeft, tcell.ColorDefault)
	for line, message := range c.thoughtText {
		tview.Print(screen, message, xOffset+10, yCenter+line, 40, tview.AlignLeft, tcell.ColorDefault)
	}
	for y := startY; y >= endY && y >= 0; y-- {
		rowBytes := []byte{'|', '.', '.', '.', '.', '.', '.', '.', '|'}
		if y < c.curTop {
			for x := 0; x < 7; x++ {
				if c.rows[y]&(1<<x) > 0 {
					rowBytes[x+1] = '#'
				}
			}
		}
		if c.falling != nil {
			for _, fallPos := range c.falling {
				if fallPos.y == y {
					rowBytes[fallPos.x+1] = '@'
				}
			}
		}
		tview.Print(screen, string(rowBytes[:]), xOffset, startY-y+1, 9, tview.AlignCenter, tcell.ColorDefault)
	}
	if endY == -1 {
		tview.Print(screen, "+-------+", xOffset, height-2, 9, tview.AlignCenter, tcell.ColorDefault)
	}
	return x + 1, y + 1, width - 2, height - 2
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

func (c *chamber) dropPiece(draw func(bool), shape int, inputPos int) int {
	c.falling = []point{}
	for _, tile := range rockShapes[shape] {
		c.falling = append(c.falling, point{tile.x + 2, tile.y + c.curTop + 3})
	}
	c.fallLeft = 2
	c.fallRight = 2 + rockWidths[shape]
	inputsUsed := 0
	c.counter++
	for ; ; inputsUsed++ {
		inputPos = c.pushPiece(inputPos)
		draw(false)
		done := c.stepPieceDown()
		draw(false)
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
	startY := c.curTop + 6
	drawHeight := 10
	if startY < drawHeight {
		startY = drawHeight
	}
	endY := startY - drawHeight

	for y := startY; y >= endY; y-- {
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
