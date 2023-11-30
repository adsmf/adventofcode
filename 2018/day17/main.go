package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1, p2 := part1()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() (int, int) {
	scan := load("input.txt")
	scan.flow()
	if !isTest {
		fmt.Println(scan)
	}
	initialSource := point{500, 0}
	allWater := 0
	settledWater := 0
	for pos, tile := range scan.grid {
		if pos == initialSource {
			continue
		}
		if pos.y < scan.min.y || pos.y > scan.max.y {
			continue
		}
		switch tile {
		case tileSettledWater:
			settledWater++
			fallthrough
		case tileFlowingWater, tileSpring:
			allWater++
		}
	}
	return allWater, settledWater
}

type scanData struct {
	grid     map[point]scanTile
	springs  map[point]bool
	min, max point
}

func (sd *scanData) flow() {
	requiresReflow := false
	for sources := sd.springs; requiresReflow || len(sources) > 0; {
		if requiresReflow && len(sources) == 0 {
			sources = map[point]bool{{500, 0}: true}
			for pos, tile := range sd.grid {
				if tile == tileSpring {
					sources[pos] = true
				}
			}
			requiresReflow = false
		}
		newSources := map[point]bool{}

		for source := range sources {
			for pos := source; pos.y <= sd.max.y; pos.y++ {
				halt := false
				switch sd.grid[pos] {
				case tileSpring:
					if pos != source {
						break
					}
				case tileFlowingWater:
				case tileSand:
					sd.grid[pos] = tileFlowingWater
				case tileSettledWater:
					above := point{pos.x, pos.y - 1}
					if sd.grid[above] == tileSpring {
						sd.grid[above] = tileFlowingWater
						requiresReflow = true
						halt = true
						break
					}
					fallthrough
				case tileClay:
					walls := 0
					newFlow := []point{}
					for _, xStep := range []int{-1, 1} {
						spillHitsWall := false
						spillFrom := point{pos.x, pos.y - 1}
						for spillPos := spillFrom; ; spillPos.x += xStep {
							spillHalt := false
							belowPos := point{spillPos.x, spillPos.y + 1}
							switch sd.grid[spillPos] {
							case tileSpring:
								if sd.grid[belowPos] != tileSettledWater {
									spillHalt = true
								}
							case tileSettledWater:
								spillHalt = true
								spillHitsWall = true
							case tileClay:
								spillHalt = true
								spillHitsWall = true
							case tileSand:
								switch sd.grid[belowPos] {
								case tileClay, tileSettledWater:
									sd.grid[spillPos] = tileFlowingWater
									newFlow = append(newFlow, spillPos)
								case tileSand, tileFlowingWater:
									sd.grid[spillPos] = tileSpring
									newSources[spillPos] = true
									spillHalt = true
								}
							case tileFlowingWater:
								newFlow = append(newFlow, spillPos)
							}
							if spillHalt {
								break
							}
						}
						if spillHitsWall {
							walls++
						}
					}
					if walls == 2 {
						for _, pos := range newFlow {
							sd.grid[pos] = tileSettledWater
						}
						newSources[source] = true
					}
					halt = true
				}
				if halt {
					break
				}
			}
		}

		sources = newSources
	}
}

func (sd scanData) String() string {
	sb := &strings.Builder{}
	height := sd.max.y - sd.min.y + 1
	width := sd.max.x - sd.min.x + 1
	sb.Grow(height * (width + 1))

	for y := sd.min.y; y <= sd.max.y; y++ {
		for x := sd.min.x; x <= sd.max.x; x++ {
			sb.WriteByte(sd.grid[point{x, y}].AsByte())
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

type point struct{ x, y int }

func minPointByOrd(a, b point) point {
	min := a
	if b.x < min.x {
		min.x = b.x
	}
	if b.y < min.y {
		min.y = b.y
	}
	return min
}
func maxPointByOrd(a, b point) point {
	max := a
	if b.x > max.x {
		max.x = b.x
	}
	if b.y > max.y {
		max.y = b.y
	}
	return max
}

type scanTile int

const (
	tileSand scanTile = iota
	tileClay
	tileSpring
	tileSettledWater
	tileFlowingWater
)

var tileReps = map[scanTile]byte{tileSand: '.', tileClay: '#', tileSpring: '+', tileSettledWater: '~', tileFlowingWater: '|'}

func (st scanTile) AsByte() byte   { return tileReps[st] }
func (st scanTile) String() string { return string(st.AsByte()) }

func load(filename string) scanData {
	spring := point{500, 0}
	scan := scanData{
		grid:    map[point]scanTile{},
		springs: map[point]bool{spring: true},
	}
	var min, max point
	first := true

	for _, line := range utils.ReadInputLines(filename) {
		var range1, range2 string
		var range1c, range2c byte
		fmt.Sscanf(line, "%c=%s %c=%s", &range1c, &range1, &range2c, &range2)
		var xVals, yVals []int
		if range1c == 'x' {
			xVals = rangeInts(range1)
			yVals = rangeInts(range2)
		} else {
			yVals = rangeInts(range1)
			xVals = rangeInts(range2)
		}
		for _, y := range yVals {
			for _, x := range xVals {
				pos := point{x, y}
				scan.grid[pos] = tileClay
				if first {
					first = false
					min = pos
					max = pos
				} else {
					min = minPointByOrd(min, pos)
					max = maxPointByOrd(max, pos)
				}
			}
		}
	}

	scan.grid[spring] = tileSpring

	min.x--
	max.x++

	scan.max = max
	scan.min = min

	return scan
}

func rangeInts(input string) []int {
	input = strings.Trim(input, ",")
	val, err := strconv.Atoi(input)
	if err == nil {
		return []int{val}
	}
	ordinals := strings.Split(input, "..")
	ints := []int{}
	for i := utils.MustInt[int](ordinals[0]); i <= utils.MustInt[int](ordinals[1]); i++ {
		ints = append(ints, i)
	}
	return ints
}

var benchmark = false
var isTest = false
