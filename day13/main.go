package main

import (
	"fmt"
	"sort"

	"github.com/adsmf/adventofcode2018/utils"
)

func main() {
	grid, carts := loadData("input.txt")
	curTick := 0
	var crashes [][]int
	for {
		curTick++
		crashes = tick(nil, grid, carts)
		if len(crashes) > 0 {
			break
		}
	}
	fmt.Printf("Crashed %d,%d during tick %d\n", crashes[0][0], crashes[0][1], curTick)
}

func tick(logFunction func(string, ...interface{}), grid gridType, carts []*cart) [][]int {
	// TODO: sort based on position rather than cart index!

	sortedCarts := []*cart{}
	for _, cart := range carts {
		sortedCarts = append(sortedCarts, cart)
	}

	sort.Sort(cartSorter{sortedCarts})

	crashes := [][]int{}
	for _, cart := range sortedCarts {
		switch cart.facing {
		case facingEast:
			cart.x++
		case facingSouth:
			cart.y++
		case facingWest:
			cart.x--
		case facingNorth:
			cart.y--
		}
		newTrack := grid[cart.y][cart.x]
		switch newTrack.trackType {
		case trackTypeCorner:
			switch {
			case (newTrack.exits.n) && (cart.facing != facingSouth):
				cart.facing = facingNorth
			case (newTrack.exits.e) && (cart.facing != facingWest):
				cart.facing = facingEast
			case (newTrack.exits.s) && (cart.facing != facingNorth):
				cart.facing = facingSouth
			case (newTrack.exits.w) && (cart.facing != facingEast):
				cart.facing = facingWest
			}
		case trackTypeIntersection:
			// oldFacing := cart.facing
			cart.facing = facing((int(cart.facing) + int(cart.nextTurn)) % int(facingEND))

			// oldNextTurn := cart.nextTurn
			cart.nextTurn++
			if cart.nextTurn > turnRight {
				cart.nextTurn = turnLeft
			}
			// if logFunction != nil {
			// 	logFunction("Cart has reached intersection.  facing %d->%d; nextTurn %d->%d", oldFacing, cart.facing, oldNextTurn, cart.nextTurn)
			// }
		}
		if checkCrashed(cart, carts) {
			crashes = append(crashes, []int{cart.x, cart.y})
		}
	}
	return crashes
}

func checkCrashed(curCart *cart, carts []*cart) bool {
	for _, checkCart := range carts {
		if checkCart == curCart {
			continue
		}
		if curCart.x == checkCart.x && curCart.y == checkCart.y {
			curCart.crashed = true
			checkCart.crashed = true
			return true
		}
	}
	return false
}

func loadData(filename string) (gridType, []*cart) {
	lines := utils.ReadInputLines(filename)
	gridWidth := len(lines[0])
	gridHeight := len(lines)
	grid := makeGrid(gridWidth, gridHeight)
	carts := []*cart{}

	for y, line := range lines {
		for x, symbol := range line {
			switch symbol {
			case '-':
				grid[y][x] = point{
					trackType: trackTypeStraight,
					exits: exits{
						e: true,
						w: true,
					},
				}
			case '|':
				grid[y][x] = point{
					trackType: trackTypeStraight,
					exits: exits{
						n: true,
						s: true,
					},
				}
			case '/':
				grid[y][x] = point{
					trackType: trackTypeCorner,
				}
			case '\\':
				grid[y][x] = point{
					trackType: trackTypeCorner,
				}
			case '+':
				grid[y][x] = point{
					trackType: trackTypeIntersection,
					exits: exits{
						n: true,
						e: true,
						s: true,
						w: true,
					},
				}
			case ' ':
			case '^':
				carts = append(carts, &cart{
					facing:   facingNorth,
					nextTurn: turnLeft,
					x:        x,
					y:        y,
				})
				grid[y][x] = point{
					trackType: trackTypeStraight,
					exits: exits{
						n: true,
						s: true,
					},
				}
			case '>':
				carts = append(carts, &cart{
					facing:   facingEast,
					nextTurn: turnLeft,
					x:        x,
					y:        y,
				})
				grid[y][x] = point{
					trackType: trackTypeStraight,
					exits: exits{
						e: true,
						w: true,
					},
				}
			case 'v':
				carts = append(carts, &cart{
					facing:   facingSouth,
					nextTurn: turnLeft,
					x:        x,
					y:        y,
				})
				grid[y][x] = point{
					trackType: trackTypeStraight,
					exits: exits{
						n: true,
						s: true,
					},
				}
			case '<':
				carts = append(carts, &cart{
					facing:   facingWest,
					nextTurn: turnLeft,
					x:        x,
					y:        y,
				})
				grid[y][x] = point{
					trackType: trackTypeStraight,
					exits: exits{
						e: true,
						w: true,
					},
				}
			default:
				fmt.Printf("Unknown symbol %c\n", symbol)
			}
		}
	}
	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridWidth; x++ {
			if grid[y][x].trackType == trackTypeCorner {
				if x > 0 && grid[y][x-1].exits.e {
					grid[y][x].exits.w = true
				}
				if x < gridWidth-1 && grid[y][x+1].exits.w {
					grid[y][x].exits.e = true
				}

				if y > 0 && grid[y-1][x].exits.s {
					grid[y][x].exits.n = true
				}
				if y < gridHeight-1 && grid[y+1][x].exits.n {
					grid[y][x].exits.s = true
				}
			}
		}
	}
	return grid, carts
}

func makeGrid(width, height int) gridType {
	newGrid := make(gridType, height)
	for i := 0; i < height; i++ {
		newGrid[i] = make(gridRow, width)
	}
	return newGrid
}

type trackType int

const (
	trackTypeStraight trackType = iota
	trackTypeCorner
	trackTypeIntersection
)

type facing int

const (
	facingNorth facing = iota
	facingEast
	facingSouth
	facingWest
	facingEND
)

type turn int

const (
	turnLeft     turn = -1
	turnStraight turn = 0
	turnRight    turn = 1
)

type gridType []gridRow
type gridRow []point

type point struct {
	trackType trackType
	exits     exits
}
type exits struct {
	n, e, s, w bool
}

type cart struct {
	facing   facing
	nextTurn turn
	x, y     int
	crashed  bool
}

type cartSorter struct {
	carts []*cart
}

// func (cs *cartSorter)

// Len is part of sort.Interface.
func (cs cartSorter) Len() int {
	return len(cs.carts)
}

// Swap is part of sort.Interface.
func (cs cartSorter) Swap(i, j int) {
	cs.carts[i], cs.carts[j] = cs.carts[j], cs.carts[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (cs cartSorter) Less(i, j int) bool {
	return cs.carts[i].y <= cs.carts[j].y && cs.carts[i].x <= cs.carts[j].x
}
