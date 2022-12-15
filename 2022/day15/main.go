package main

import (
	_ "embed"
	"fmt"
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
	row := 2000000
	excludes := excludeRanges{}
	for i := 0; i < g.len; i++ {
		sensor := g.data[i].sensorPos
		coverage := g.data[i].dist
		if sensor.y+coverage >= row && sensor.y-coverage <= row {
			excludeDist := coverage - sensor.manhattan(point{sensor.x, row})
			exclude := excludeRange{sensor.x - excludeDist, sensor.x + excludeDist}
			excludes.ranges[excludes.len] = exclude
			excludes.len++
		}
	}
	excludes = consolidateRanges(excludes)
	count := 0
	for i := 0; i < excludes.len; i++ {
		exclude := excludes.get(i)
		count += exclude.end - exclude.start + 1
		for beaconIdx := 0; beaconIdx < g.len; beaconIdx++ {
			beaconX := g.data[beaconIdx].beaconPos.x
			if beaconX >= exclude.start && beaconX <= exclude.end {
				count--
				break
			}
		}
	}
	return count
}

func part2(g grid) int {
	type quad [4]int
	x0quads := [30]quad{}
	for i := 0; i < g.len; i++ {
		pair := g.data[i]
		x0tl := pair.sensorPos.y - pair.dist - 1 - pair.sensorPos.x
		x0br := pair.sensorPos.y + pair.dist + 1 - pair.sensorPos.x

		x0tr := pair.sensorPos.y - pair.dist - 1 + pair.sensorPos.x
		x0bl := pair.sensorPos.y + pair.dist + 1 + pair.sensorPos.x
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

type excludeRange struct{ start, end int }
type excludeRanges struct {
	ranges [25]excludeRange
	len    int
}

func (e excludeRanges) get(index int) excludeRange { return e.ranges[index] }

func consolidateRanges(excludes excludeRanges) excludeRanges {
	for moved := true; moved; {
		moved = false
		for i := 0; i < excludes.len-1; i++ {
			if excludes.ranges[i].start > excludes.ranges[i+1].start {
				excludes.ranges[i], excludes.ranges[i+1] = excludes.ranges[i+1], excludes.ranges[i]
				moved = true
			}
		}
	}
	curRange := 0
	for i := 0; i < excludes.len; i++ {
		if excludes.ranges[curRange].end >= excludes.ranges[i].start {
			if excludes.ranges[i].end > excludes.ranges[curRange].end {
				excludes.ranges[curRange].end = excludes.ranges[i].end
			}
		} else {
			curRange++
			excludes.ranges[curRange] = excludes.ranges[i]
		}
	}
	excludes.len = curRange + 1
	return excludes
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
