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
	type quad [4]int
	x0quads := make([]quad, g.len)
	for i := 0; i < g.len; i++ {
		pair := g.data[i]
		dist := pair.sensorPos.manhattan(pair.beaconPos)
		x0tl := pair.sensorPos.y - dist - 1 - pair.sensorPos.x
		x0br := pair.sensorPos.y + dist + 1 - pair.sensorPos.x

		x0tr := pair.sensorPos.y - dist - 1 + pair.sensorPos.x
		x0bl := pair.sensorPos.y + dist + 1 + pair.sensorPos.x
		x0quads[i] = quad{x0tl, x0br, x0tr, x0bl}
	}
	testMerge := func(x0pos, x0neg int) (bool, point) {
		if x0neg < x0pos {
			return false, point{}
		}
		diff := x0neg - x0pos
		if diff&1 == 1 {
			return false, point{}
		}
		diff >>= 1
		mergePos := point{diff, x0pos + (diff)}
		return g.potentialBeacon(mergePos), mergePos
	}
	for q1idx := 0; q1idx < g.len-1; q1idx++ {
		q1 := x0quads[q1idx]
		for q2idx := q1idx + 1; q2idx < g.len; q2idx++ {
			q2 := x0quads[q2idx]
			pairs := [8][2]int{
				{q1[0], q2[2]},
				{q1[0], q2[3]},
				{q1[1], q2[2]},
				{q1[1], q2[3]},
				{q2[0], q1[2]},
				{q2[0], q1[3]},
				{q2[1], q1[2]},
				{q2[1], q1[3]},
			}
			for _, pair := range pairs {
				valid, pos := testMerge(pair[0], pair[1])
				if valid {
					return 4000000*pos.x + pos.y
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
		sensor := point{sX, sY}
		beacon := point{bX, bY}
		g.data[id] = sensorData{
			sensorPos: sensor,
			beaconPos: beacon,
			dist:      sensor.manhattan(beacon),
		}
		g.len++
		id++
	}
	return g
}

type sensorData struct {
	sensorPos point
	beaconPos point
	dist      int
}

type grid struct {
	len  int
	data [30]sensorData
}

func (g grid) potentialBeacon(pos point) bool {
	for i := 0; i < g.len; i++ {
		data := g.data[i]
		if pos == data.beaconPos || pos == data.sensorPos {
			return false
		}
		if data.sensorPos.manhattan(pos) <= data.dist {
			return false
		}
	}
	return true
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

var benchmark = false
