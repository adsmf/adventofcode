package main

import (
	_ "embed"
	"fmt"
	"sort"
)

//go:embed input.txt
var input []byte

func main() {
	g := loadInput()
	p1 := part1(g)
	p2 := part2(g)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(g grid) int {
	return g.countExcluded(2000000)
}

func part2(g grid) int {
	maxVal := 4000000
	excluded := &excludeRanges{}
	for y := 0; y <= maxVal; y++ {
		excluded.reset()
		g.locateExcluded(y, excluded)
		if excluded.len > 1 {
			for i := 0; i < excluded.len-1; i++ {
				gap := excluded.get(i).end + 1
				if gap > 0 && gap <= maxVal {
					return 4000000*gap + y
				}
			}
		}
	}
	return -1
}

func loadInput() grid {
	g := grid{}
	id := 0
	sX, sY, bX, bY := 0, 0, 0, 0
	for pos := 12; pos < len(input); pos += 13 {
		sX, pos = getInt(input, pos)
		sY, pos = getInt(input, pos+4)
		bX, pos = getInt(input, pos+25)
		bY, pos = getInt(input, pos+4)
		g.data[id] = sensorData{
			sensorPos: point{sX, sY},
			beaconPos: point{bX, bY},
		}
		g.len++
		id++
	}
	return g
}

func getInt(in []byte, pos int) (int, int) {
	accumulator := 0
	negative := false
	if in[pos] == '-' {
		negative = true
		pos++
	}
	for ; in[pos] >= '0' && in[pos] <= '9'; pos++ {
		accumulator *= 10
		accumulator += int(in[pos] & 0xf)
	}
	if negative {
		accumulator = -accumulator
	}
	return accumulator, pos
}

type grid struct {
	len  int
	data [30]sensorData
}

type sensorData struct {
	sensorPos point
	beaconPos point
}

type excludeRange struct{ start, end int }
type excludeRanges struct {
	ranges [25]excludeRange
	len    int
}

func (e *excludeRanges) add(exclude excludeRange) {
	e.ranges[e.len] = exclude
	e.len++
}
func (e excludeRanges) get(index int) excludeRange { return e.ranges[index] }
func (e *excludeRanges) reset()                    { e.len = 0 }
func (e excludeRanges) Len() int                   { return e.len }
func (e excludeRanges) Less(i, j int) bool         { return e.ranges[i].start < e.ranges[j].start }
func (e *excludeRanges) Swap(i, j int)             { e.ranges[i], e.ranges[j] = e.ranges[j], e.ranges[i] }

func (e *excludeRanges) consolidate() {
	sort.Sort(e)
	curRange := 0
	for i := 0; i < e.len; i++ {
		if e.ranges[curRange].end >= e.ranges[i].start {
			if e.ranges[i].end > e.ranges[curRange].end {
				e.ranges[curRange].end = e.ranges[i].end
			}
		} else {
			curRange++
			e.ranges[curRange] = e.ranges[i]
		}
	}
	e.len = curRange + 1
}
func (g grid) locateExcluded(row int, ranges *excludeRanges) {
	for i := 0; i < g.len; i++ {
		sensor := g.data[i].sensorPos
		beacon := g.data[i].beaconPos
		coverage := sensor.manhattan(beacon)
		if sensor.y+coverage >= row && sensor.y-coverage <= row {
			excludeDist := coverage - sensor.manhattan(point{sensor.x, row})
			exclude := excludeRange{sensor.x - excludeDist, sensor.x + excludeDist}
			ranges.add(exclude)
		}
	}
	ranges.consolidate()
}

func (g grid) countExcluded(row int) int {
	ranges := &excludeRanges{}
	g.locateExcluded(row, ranges)
	count := 0
	ignorePoints := map[int]bool{}
	for i := 0; i < g.len; i++ {
		sensor := g.data[i].sensorPos
		beacon := g.data[i].beaconPos
		if sensor.y == row {
			ignorePoints[sensor.x] = true
		}
		if beacon.y == row {
			ignorePoints[beacon.x] = true
		}
	}
	for i := 0; i < ranges.len; i++ {
		exclude := ranges.get(i)
		count += exclude.end - exclude.start + 1
		for ignore := range ignorePoints {
			if ignore >= exclude.start && ignore <= exclude.end {
				count--
			}
		}
	}
	return count
}

func intAbs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}

type point struct{ x, y int }

func (p point) sub(o point) point { return point{p.x - o.x, p.y - o.y} }
func (p point) manhattan(o point) int {
	diff := p.sub(o)
	return intAbs(diff.x) + intAbs(diff.y)
}

var benchmark = false
