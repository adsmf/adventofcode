import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve(input)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve(in string) (int, int) {
	vals := utils.GetInts(in)
	minX, maxX, minY, maxY := vals[0], vals[1], vals[2], vals[3]
	bestY := 0
	countValid := 0
	for dY := minY; dY < -minY; dY++ {
		for dX := maxX; dX > 0; dX-- {
			maxY, hit := runSim(dX, dY, minX, maxX, minY, maxY)
			if hit {
				if bestY < maxY {
					bestY = maxY
				}
				countValid++
			}
		}
	}
	return bestY, countValid
}

func runSim(dX, dY int, minX, maxX, minY, maxY int) (int, bool) {
	x, y := 0, 0
	bestY := 0
	for x <= maxX && y >= minY {
		x += dX
		y += dY
		if y > bestY {
			bestY = y
		}
		if minX <= x && x <= maxX && minY <= y && y <= maxY {
			return bestY, true
		}
		if dX > 0 {
			dX--
		}
		dY--
	}
	return 0, false
}
