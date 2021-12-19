package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := part1(input)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(in string) (int, int) {
	scanners := load(in)
	scanners[0].readings = scanners[0].orientedReadings[0]
	fixedScanners := []scannerInfo{scanners[0]}
	unfixedScanners := append([]scannerInfo{}, scanners[1:]...)

	for changed := true; changed && len(unfixedScanners) > 0; {
		changed = false
		for _, fixed := range fixedScanners {
			for i := 0; i < len(unfixedScanners); i++ {
				unfixed := unfixedScanners[i]
				matched, nowFixed := matches(fixed, unfixed)
				if matched {
					changed = true
					fixedScanners = append(fixedScanners, nowFixed)
					unfixedScanners = append(unfixedScanners[:i], unfixedScanners[i+1:]...)
					i--
				}
			}
		}
	}
	allPoints := pointSet{}
	for _, scanner := range fixedScanners {
		for pos := range scanner.readings {
			allPoints[pos] = true
		}
	}
	maxDist := 0
	for _, a := range fixedScanners {
		for _, b := range fixedScanners {
			dist := a.pos.point.sub(b.pos.point).magnitude()
			if dist > maxDist {
				maxDist = dist
			}
		}
	}
	return len(allPoints), maxDist
}

func matches(fixed, unfixed scannerInfo) (bool, scannerInfo) {
	for facing, possible := range unfixed.orientedReadings {
		for fB := range fixed.readings {
			for pB := range possible {
				offset := fB.sub(pB)
				count := 0
				for adjPB := range possible {
					adjPB = adjPB.add(offset)
					if fixed.readings[adjPB] {
						count++
						if count >= 12 {
							adjustedReadings := make(pointSet, len(possible))
							for pos := range possible {
								adjustedReadings[pos.add(offset)] = true
							}
							nowFixed := scannerInfo{
								id: unfixed.id,
								pos: locationInfo{
									point:       offset,
									orientation: orientation(facing),
								},
								readings: adjustedReadings,
							}
							return true, nowFixed
						}
					}
				}
			}
		}
	}
	return false, unfixed
}

func load(in string) []scannerInfo {
	blocks := strings.Split(strings.TrimSpace(in), "\n\n")
	scanners := make([]scannerInfo, 0, len(blocks))
	for scannerID, block := range blocks {
		lines := strings.Split(strings.TrimSpace(block), "\n")
		scanner := scannerInfo{
			id:               scannerID,
			readings:         make(pointSet, len(lines)),
			orientedReadings: make([]pointSet, 24),
		}
		for i := 0; i < 24; i++ {
			scanner.orientedReadings[i] = make(pointSet, len(lines))
		}
		for lineNum, line := range lines {
			if lineNum == 0 {
				continue
			}
			coords := utils.GetInts(line)
			pos := point{coords[0], coords[1], coords[2]}
			for or, oriented := range genOrientations(pos) {
				scanner.orientedReadings[or][oriented] = true
			}

		}
		scanners = append(scanners, scanner)
	}
	return scanners
}

type scannerInfo struct {
	id               int
	pos              locationInfo
	readings         pointSet
	orientedReadings []pointSet
}
type locationInfo struct {
	point       point
	orientation orientation
}
type orientation byte

type pointSet map[point]bool

func (ps pointSet) String() string {
	sb := strings.Builder{}

	first := true
	for pos := range ps {
		if !first {
			sb.WriteByte(' ')
		}
		first = false
		sb.WriteString(pos.String())
	}
	return sb.String()
}

type point struct{ x, y, z int }

func (p point) String() string {
	sb := strings.Builder{}

	sb.WriteByte('[')
	sb.Write([]byte(strconv.Itoa(p.x)))
	sb.WriteByte(',')
	sb.Write([]byte(strconv.Itoa(p.y)))
	sb.WriteByte(',')
	sb.Write([]byte(strconv.Itoa(p.z)))
	sb.WriteByte(']')

	return sb.String()
}

func (p point) add(q point) point { return point{p.x + q.x, p.y + q.y, p.z + q.z} }
func (p point) sub(q point) point { return point{p.x - q.x, p.y - q.y, p.z - q.z} }
func (p point) magnitude() int {
	return utils.IntAbs(p.x) + utils.IntAbs(p.y) + utils.IntAbs(p.z)
}
func (p point) roll() point {
	return point{p.x, p.z, -p.y}
}
func (p point) turn() point {
	return point{-p.y, p.x, p.z}
}

func genOrientations(pos point) []point {
	// Inspired by https://stackoverflow.com/questions/16452383/how-to-get-all-24-rotations-of-a-3-dimensional-array
	orientations := make([]point, 0, 24)
	for cycle := 0; cycle < 2; cycle++ {
		for step := 0; step < 3; step++ {
			orientations = append(orientations, pos)
			pos = pos.roll()
			for i := 0; i < 3; i++ {
				orientations = append(orientations, pos)
				pos = pos.turn()
			}
		}
		pos = pos.roll().turn().roll()
	}
	return orientations
}

var benchmark = false
