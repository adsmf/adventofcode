package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/adsmf/adventofcode2019/utils/pathfinding/astar"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	m := loadMap("input.txt")
	return m.solve()
}

func part2() int {
	m := loadMap("input.txt")
	m.recursive = true
	return m.solve()
}

type maze struct {
	grid       map[int]map[point]tile
	portals    map[point]point
	start, end point
	minX, minY int
	maxX, maxY int
	recursive  bool
}

func (m *maze) solve() int {
	startTile := m.grid[0][m.start]
	endTile := m.grid[0][m.end]
	route, err := astar.Route(startTile, endTile)
	if err != nil {
		fmt.Printf("Could not find route!\n")
		return 0
	}

	return len(route) - 3
}

func (m *maze) isEdgePortal(p point) bool {
	if p.x == m.minX ||
		p.x >= m.maxX-1 ||
		p.y == m.minY ||
		p.y >= m.maxY-1 {
		return true
	}
	return false
}

func (m *maze) get(pos point, level int) tile {
	if m.grid[level] == nil {
		m.grid[level] = map[point]tile{}
		for copyPos, copyTile := range m.grid[0] {
			copyTile.level = level
			m.grid[level][copyPos] = copyTile
		}
	}
	return m.grid[level][pos]
}

func (m *maze) set(pos point, level int, val tile) {
	val.pos = pos
	val.maze = m
	val.level = level
	if m.grid[level] == nil {
		m.grid[level] = map[point]tile{}
		for copyPos, copyTile := range m.grid[0] {
			copyTile.level = level
			m.grid[level][copyPos] = copyTile
		}
	}
	m.grid[level][pos] = val
	if m.minX > pos.x {
		m.minX = pos.x
	}
	if m.maxX < pos.x {
		m.maxX = pos.x
	}
	if m.minY > pos.y {
		m.minY = pos.y
	}
	if m.maxY < pos.y {
		m.maxY = pos.y
	}
}

func (m maze) String() string {
	newString := ""
	for y := m.minY; y <= m.maxY; y++ {
		for x := m.minX; x <= m.maxX; x++ {
			newString += fmt.Sprintf("%v", m.grid[0][point{x, y}])
		}
		newString += fmt.Sprintln()
	}
	return newString
}

type tile struct {
	tileType tileType
	portalId string
	pos      point
	maze     *maze
	level    int
}

func (t tile) Heuristic(astar.Node) astar.Cost {
	return 1
}
func (t tile) Paths() []astar.Edge {
	edges := []astar.Edge{}

	for _, nPos := range t.pos.neighbours() {
		n := t.maze.get(nPos, t.level)
		if n.tileType == tileTypeEmpty ||
			nPos == t.maze.end {
			n.level = t.level
			edges = append(
				edges,
				astar.Edge{
					To:   n,
					Cost: 1,
				},
			)
		} else if n.tileType == tileTypePortal {
			nextLevel := t.level
			if n.maze.recursive {
				if t.maze.isEdgePortal(nPos) {
					nextLevel--
				} else {
					nextLevel++
				}
			}
			if nextLevel < 0 {
				continue
			}
			portalPos := t.maze.portals[nPos]
			for _, portalNeighbourPos := range portalPos.neighbours() {
				portalNeighbour := t.maze.get(portalNeighbourPos, nextLevel)
				if portalNeighbour.tileType == tileTypeEmpty {
					edges = append(
						edges,
						astar.Edge{
							To:   portalNeighbour,
							Cost: 1,
						},
					)
				}
			}
		}

	}

	return edges
}

func (t tile) String() string {
	switch t.tileType {
	case tileTypeEmpty:
		return "  "
	case tileTypeWall:
		return "██"
	case tileTypePortal:
		return t.portalId
	case tileTypeUnkown:
		return "░░"
	default:
		return "!!"
	}
}

type tileType int

const (
	tileTypeUnkown tileType = iota
	tileTypeEmpty
	tileTypeWall
	tileTypePortal
)

type point struct {
	x, y int
}

func (p point) neighbours() []point {
	return []point{
		point{p.x - 1, p.y},
		point{p.x + 1, p.y},
		point{p.x, p.y - 1},
		point{p.x, p.y + 1},
	}
}

func loadMap(filename string) maze {
	m := maze{
		grid:    map[int]map[point]tile{},
		portals: map[point]point{},
	}
	// m.grid[0] = map[point]tile{}

	// lines := utils.ReadInputLines(filename)
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(raw), "\n")
	tempPortals := map[point]rune{}
	for y, line := range lines {
		for x, char := range line {
			pos := point{x, y}
			switch {
			case char == '.':
				m.set(pos, 0, tile{tileType: tileTypeEmpty})
			case char == '#':
				m.set(pos, 0, tile{tileType: tileTypeWall})
			case 'A' <= char && char <= 'Z':
				// TODO portal links
				m.set(pos, 0, tile{
					tileType: tileTypePortal,
					portalId: string(char),
				})
				tempPortals[pos] = char
			}
		}
	}
	portalLinks := map[string][]point{}
	for pos, char := range tempPortals {
		var portID string
		var isNearest bool
		for _, n := range pos.neighbours() {
			if m.get(n, 0).tileType == tileTypeEmpty {
				isNearest = true
			} else if _, found := tempPortals[n]; found {
				if pos.x > n.x || pos.y > n.y {
					portID = fmt.Sprintf("%c%c", tempPortals[n], char)
				} else {
					portID = fmt.Sprintf("%c%c", char, tempPortals[n])
				}

			}

		}
		if isNearest {
			m.set(pos, 0, tile{tileType: tileTypePortal, portalId: portID})
			if portID == "AA" {
				m.start = pos
				continue
			} else if portID == "ZZ" {
				m.end = pos
				continue
			}
			if _, found := portalLinks[portID]; found {
				portalLinks[portID] = append(portalLinks[portID], pos)
			} else {
				portalLinks[portID] = []point{pos}
			}
		} else {
			m.set(pos, 0, tile{tileType: tileTypeUnkown})
		}
	}
	for _, ends := range portalLinks {
		if len(ends) != 2 {
			fmt.Printf("%d\n", len(ends))
			panic("Portals only have two ends!")
		}
		m.portals[ends[0]] = ends[1]
		m.portals[ends[1]] = ends[0]
	}
	return m
}
