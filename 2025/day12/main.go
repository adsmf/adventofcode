package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
	}
}

func part1() int {
	shapeSizes := []int{}
	packable := 0
	utils.EachSectionMB(input, "\n\n", func(secIdx int, section string) (done bool) {
		if section[1] == ':' {
			numSpaces := 0
			for x := range 3 {
				for y := range 3 {
					if section[3+y*4+x] == '#' {
						numSpaces++
					}
				}
			}
			shapeSizes = append(shapeSizes, numSpaces)
			return
		}
		utils.EachLine(section, func(lineIdx int, line string) (done bool) {
			vals := utils.GetInts(line)
			width, height := vals[0], vals[1]
			vals = vals[2:]
			areaAvail := width * height
			maxSpace := 0
			areqReq := 0
			for shapeIdx, count := range vals {
				areqReq += shapeSizes[shapeIdx] * count
				maxSpace += 9 * count
			}
			if areqReq > areaAvail {
				return
			}
			if (maxSpace <= areaAvail) != (areqReq <= areaAvail) {
				fmt.Println(areaAvail, maxSpace, maxSpace <= areaAvail, areqReq, areqReq <= areaAvail)
				panic("Maybe we can't just count slots")
			}
			packable++
			return
		})
		return
	})
	return packable
}

type point struct{ x, y int }

var benchmark = false
