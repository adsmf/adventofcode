package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	tiles := load("input.txt")
	goodHorizontal := map[tilePair]bool{}
	goodVertical := map[tilePair]bool{}
	for tile1 := range tiles {
		for tile2 := range tiles {
			for o1 := orient0; o1 < flippingOrientMax; o1++ {
				for o2 := orient0; o2 < flippingOrientMax; o2++ {
					if tiles[tile1].row(9, o1) == tiles[tile2].row(0, o2) {
						pair := tilePair{
							tile1, tile2,
							o1, o2,
						}
						goodVertical[pair] = true
					}
					if tiles[tile1].col(9, o1) == tiles[tile2].col(0, o2) {
						pair := tilePair{
							tile1, tile2,
							o1, o2,
						}
						goodHorizontal[pair] = true
					}
				}
			}
		}
	}
	tileIDs := make([]int, 0, len(tiles))
	for id := range tiles {
		tileIDs = append(tileIDs, id)
	}
	tilesPerRow := int(math.Sqrt(float64(len(tiles))))

	order, rotations, _ := tryArrangement(tiles, tilesPerRow, []int{}, tileIDs, []orientation{}, goodHorizontal, goodVertical)

	cornerMultiple := order[0]
	cornerMultiple *= order[tilesPerRow-1]
	cornerMultiple *= order[len(tiles)-1]
	cornerMultiple *= order[len(tiles)-tilesPerRow]

	image := map[point]bool{}
	for tileRow := 0; tileRow < len(order)/tilesPerRow; tileRow++ {
		for tileColOffset := 0; tileColOffset < tilesPerRow; tileColOffset++ {
			imX, imY := tileColOffset*8, tileRow*8
			tile := tiles[order[tileRow*tilesPerRow+tileColOffset]]
			tileRotation := rotations[tileRow*tilesPerRow+tileColOffset]
			for y := 0; y < 8; y++ {
				rowPixels := tile.row(y+1, tileRotation)
				for x := 0; x < 8; x++ {
					if rowPixels[x+1] == '#' {
						image[point{imX + x, imY + y}] = true
					}
				}
			}
		}
	}
	for rotation := orient0; rotation < flippingOrientMax; rotation++ {
		numMonsters := 0
		inMonster := map[point]bool{}
		rotatedMonster := rotatedMonster(rotation)
		for x := 0; x < 12*tilesPerRow; x++ {
			for y := 0; y < 12*tilesPerRow; y++ {
				isMonster := true
				for monsterPoint := range rotatedMonster {
					if !image[point{x + monsterPoint.x, y + monsterPoint.y}] {
						isMonster = false
						break
					}
				}
				if isMonster {
					for monsterPoint := range rotatedMonster {
						inMonster[point{x + monsterPoint.x, y + monsterPoint.y}] = true
					}
					numMonsters++
				}
			}
		}
		if numMonsters > 0 {
			count := 0
			for pos := range image {
				if _, found := inMonster[pos]; !found {
					count++
				}
			}

			return cornerMultiple, count
		}
	}

	panic("Could not solve!")
}

//                     #
//   #    ##    ##    ###
//    #  #  #  #  #  #
//   01234567890123456789
var monsterPoints = map[point]bool{
	{18, 0}: true,
	{0, 1}:  true, {5, 1}: true, {6, 1}: true, {11, 1}: true, {12, 1}: true, {17, 1}: true, {18, 1}: true, {19, 1}: true,
	{1, 2}: true, {4, 2}: true, {7, 2}: true, {10, 2}: true, {13, 2}: true, {16, 2}: true,
}

func rotatedMonster(rot orientation) map[point]bool {
	rotated := make(map[point]bool, 15)
	for pos := range monsterPoints {
		newPos := pos
		height, width := 2, 19
		calcWidth := width

		switch rot % orientMax {
		case orient90:
			newPos.x, newPos.y = height-newPos.y, newPos.x
			calcWidth = height
		case orient180:
			newPos.x, newPos.y = width-newPos.x, height-newPos.y
		case orient270:
			newPos.x, newPos.y = newPos.y, width-newPos.x
			calcWidth = height
		}

		if (rot & flipX) > 0 {
			newPos.x = calcWidth - newPos.x
		}

		rotated[newPos] = true
	}
	return rotated
}

type tilePair struct {
	t1, t2 int
	o1, o2 orientation
}

func tryArrangement(tiles tileset, tilesPerRow int, order, remaining []int, rotations []orientation, goodHorizontal, goodVertical map[tilePair]bool) ([]int, []orientation, bool) {
	if len(order) > 1 {
		curTileIndex := len(order) - 1
		curTileX := curTileIndex % tilesPerRow
		curTileY := curTileIndex / tilesPerRow
		curTile := tiles[order[curTileIndex]]
		if curTileY > 0 {
			aboveIndex := curTileIndex - tilesPerRow
			aboveTile := tiles[order[aboveIndex]]
			if !goodVertical[tilePair{aboveTile.id, curTile.id, rotations[aboveIndex], rotations[curTileIndex]}] {
				return nil, nil, false
			}
		}
		if curTileX > 0 {
			leftIndex := curTileIndex - 1
			leftTile := tiles[order[leftIndex]]
			if !goodHorizontal[tilePair{leftTile.id, curTile.id, rotations[leftIndex], rotations[curTileIndex]}] {
				return nil, nil, false
			}
		}
	}

	if len(remaining) > 0 {
		for tryIndex, tryTile := range remaining {
			newRemaining := make([]int, 0, len(remaining)-1)
			newRemaining = append(newRemaining, remaining[:tryIndex]...)
			newRemaining = append(newRemaining, remaining[tryIndex+1:]...)
			newOrder := make([]int, len(order), len(order)+1)
			copy(newOrder, order)
			newOrder = append(newOrder, tryTile)
			for rot := 0; rot < flippingOrientMax; rot++ {
				nextOrientations := make([]orientation, len(rotations), len(rotations)+1)
				copy(nextOrientations, rotations)
				nextOrientations = append(nextOrientations, orientation(rot))
				finalOrder, finalRotations, valid := tryArrangement(tiles, tilesPerRow, newOrder, newRemaining, nextOrientations, goodHorizontal, goodVertical)
				if valid {
					return finalOrder, finalRotations, true
				}
			}
		}
		return nil, nil, false
	}
	return order, rotations, true
}

type tileset map[int]*imageTile

type tro struct {
	tile, row   int
	orientation orientation
	vertical    bool
}

var knownEdges = map[tro]string{}

type imageTile struct {
	id     int
	pixels map[orientedPoint]bool
}

func (t imageTile) row(rowNum int, rot orientation) string {
	if edge, found := knownEdges[tro{t.id, rowNum, rot, false}]; found {
		return edge
	}
	row := make([]byte, 0, 10)
	start, step := 0, 1
	if (rot & flipX) > 0 {
		start, step = 9, -1
	}
	for x := start; x < 10 && x >= 0; x += step {
		if t.pixels[orientedPoint{x, rowNum, rot % orientMax}] {
			row = append(row, '#')
		} else {
			row = append(row, '.')
		}
	}
	knownEdges[tro{t.id, rowNum, rot, false}] = string(row)
	return string(row)
}

func (t imageTile) col(colNum int, rot orientation) string {
	if edge, found := knownEdges[tro{t.id, colNum, rot, true}]; found {
		return edge
	}
	row := make([]byte, 0, 10)
	useCol := colNum
	if (rot & flipX) > 0 {
		useCol = 9 - useCol
	}
	for y := 0; y < 10; y++ {
		if t.pixels[orientedPoint{useCol, y, rot % orientMax}] {
			row = append(row, '#')
		} else {
			row = append(row, '.')
		}
	}
	knownEdges[tro{t.id, colNum, rot, true}] = string(row)
	return string(row)
}

type orientation int

const (
	orient0 orientation = iota
	orient90
	orient180
	orient270
	orientMax         = 1 << 2
	flipX             = 1 << 2
	flippingOrientMax = 1 << 3
)

type point struct{ x, y int }
type orientedPoint struct {
	x, y        int
	orientation orientation
}

func load(filename string) tileset {
	inputBytes, _ := ioutil.ReadFile(filename)
	tiles := tileset{}

	for _, block := range strings.Split(string(inputBytes), "\n\n") {
		lines := strings.Split(block, "\n")
		id := 0
		fmt.Sscanf(lines[0], "Tile %4d:", &id)
		tile := imageTile{
			id:     id,
			pixels: map[orientedPoint]bool{},
		}
		for y := 0; y < 10; y++ {
			line := lines[y+1]
			for x, char := range line {
				switch char {
				case '#':
					tile.pixels[orientedPoint{x, y, orient0}] = true
					tile.pixels[orientedPoint{9 - y, x, orient90}] = true
					tile.pixels[orientedPoint{9 - x, 9 - y, orient180}] = true
					tile.pixels[orientedPoint{y, 9 - x, orient270}] = true
				}
			}
		}
		tiles[id] = &tile
	}

	return tiles
}

var benchmark = false
