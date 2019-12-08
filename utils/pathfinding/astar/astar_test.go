package astar

import (
	"fmt"
	"math"
	"testing"

	"github.com/adsmf/adventofcode2019/utils"
	"github.com/stretchr/testify/assert"
)

func TestRouting(t *testing.T) {
	type example struct {
		testMapFile string
		routeLength int
		cost        Cost
	}
	tests := []example{
		example{testMapFile: "map1.txt", routeLength: 7, cost: 7},
		example{testMapFile: "map2.txt", routeLength: 11, cost: 11},
		example{testMapFile: "map3.txt", routeLength: 11, cost: 11},
		example{testMapFile: "map4.txt", routeLength: 15, cost: 15},
		example{testMapFile: "map5.txt", routeLength: 6, cost: 6},
		example{testMapFile: "map6.txt", routeLength: 8, cost: 8},
		example{testMapFile: "map7.txt", routeLength: 6, cost: 10},
		example{testMapFile: "map8.txt", routeLength: 30, cost: 30},
	}
	for id, testDef := range tests {
		t.Run(fmt.Sprintf("Example%d:%s", id+1, testDef.testMapFile), func(t *testing.T) {
			test := loadTestMap(testDef.testMapFile)
			route, err := Route(test.start, test.end)
			test.route = route
			t.Logf("Grid:\n%s", test.String())
			assert.NoError(t, err)
			assert.Equal(t, testDef.routeLength, len(route))
			assert.Equal(t, test.start, route[0])
			assert.Equal(t, test.end, route[len(route)-1])
			routeCost := Cost(0)
			for _, node := range route {
				routeCost = routeCost + node.(*testNode).cost
			}
			assert.Equal(t, testDef.cost, routeCost)
		})
	}
}

func TestUnroutable(t *testing.T) {
	test := loadTestMap("unroutable.txt")
	route, err := Route(test.start, test.end)
	assert.Error(t, err)
	assert.IsType(t, []Node{}, route)
	assert.EqualValues(t, 0, len(route))
}

type testNode struct {
	x, y   int
	cost   Cost
	grid   *testMap
	symbol rune
}

func (n testNode) Heuristic(from Node) Cost {
	fromNode := from.(*testNode)

	cost := fromNode.cost

	if fromNode.cost < math.MaxFloat32 {
		xDiff := fromNode.x - n.x
		if xDiff < 0 {
			xDiff = -xDiff
		}
		yDiff := fromNode.y - n.y
		if yDiff < 0 {
			yDiff = -yDiff
		}
		cost += Cost(xDiff + yDiff)
	}
	return cost
}

func (n testNode) Paths() []Edge {
	options := []Edge{}

	if n.grid == nil {
		return options
	}

	if n.x > 0 {
		left := n.grid.grid[n.x-1][n.y]
		if left != nil {
			options = append(options, Edge{
				To:   left,
				Cost: n.cost,
			})
		}
	}
	if n.y > 0 {
		up := n.grid.grid[n.x][n.y-1]
		if up != nil {
			options = append(options, Edge{
				To:   up,
				Cost: n.cost,
			})
		}
	}
	if n.x < n.grid.maxX {
		right := n.grid.grid[n.x+1][n.y]
		if right != nil {
			options = append(options, Edge{
				To:   right,
				Cost: n.cost,
			})
		}
	}
	if n.y < n.grid.maxY {
		down := n.grid.grid[n.x][n.y+1]
		if down != nil {
			options = append(options, Edge{
				To:   down,
				Cost: n.cost,
			})
		}
	}
	return options
}

type testMap struct {
	grid  map[int]map[int]*testNode
	start *testNode
	end   *testNode
	maxX  int
	maxY  int
	route []Node
}

func (tm *testMap) String() string {
	retString := fmt.Sprintf("%dx%d grid; route length: %d\n\n", tm.maxX, tm.maxY, len(tm.route))
	routed := map[int]rune{}
	for idx, node := range tm.route {
		tn := node.(*testNode)
		sym := rune('0' + ((idx + 1) % 10))
		pos := tn.x + tn.y*(tm.maxX+1)

		routed[pos] = sym
	}
	for row := 0; row <= tm.maxY; row++ {
		for col := 0; col <= tm.maxX; col++ {
			node := tm.grid[col][row]
			var sym rune
			if node == nil {
				sym = '#'
			} else {
				sym = node.symbol
			}
			retString = fmt.Sprintf("%s%c", retString, sym)
		}

		retString = fmt.Sprintf("%s    ", retString)
		for col := 0; col <= tm.maxX; col++ {
			pos := col + row*(tm.maxX+1)
			var sym rune
			if routeSym, found := routed[pos]; found {
				sym = routeSym
			} else {
				node := tm.grid[col][row]
				if node == nil {
					sym = '#'
				} else {
					sym = node.symbol
				}
			}
			retString = fmt.Sprintf("%s%c", retString, sym)
		}

		retString = fmt.Sprintln(retString)
	}
	return retString
}

func loadTestMap(file string) testMap {
	newTestMap := testMap{
		grid: map[int]map[int]*testNode{},
	}
	lines := utils.ReadInputLines("test_fixtures/" + file)
	for row, line := range lines {
		if row > newTestMap.maxY {
			newTestMap.maxY = row
		}
		for col, symbol := range line {
			if col > newTestMap.maxX {
				newTestMap.maxX = col
			}
			if newTestMap.grid[col] == nil {
				newTestMap.grid[col] = make(map[int]*testNode)
			}
			switch symbol {
			case 's':
				newTestMap.start = &testNode{x: col, y: row, cost: Cost(1), grid: &newTestMap, symbol: 's'}
				newTestMap.grid[col][row] = newTestMap.start
			case 'f':
				newTestMap.end = &testNode{x: col, y: row, cost: Cost(1), grid: &newTestMap, symbol: 'f'}
				newTestMap.grid[col][row] = newTestMap.end
			case '_':
				newTestMap.grid[col][row] = &testNode{x: col, y: row, cost: Cost(0.1), grid: &newTestMap, symbol: '_'}
			case '~':
				newTestMap.grid[col][row] = &testNode{x: col, y: row, cost: Cost(5), grid: &newTestMap, symbol: '~'}
			case '.':
				newTestMap.grid[col][row] = &testNode{x: col, y: row, cost: Cost(1), grid: &newTestMap, symbol: '.'}
			}
		}
	}
	return newTestMap
}
