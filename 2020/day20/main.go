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

	tilesPerRow := int(math.Sqrt(float64(len(tiles))))

	order, rotations := align(tiles, tilesPerRow)

	cornerMultiple := order[0] *
		order[tilesPerRow-1] *
		order[len(tiles)-1] *
		order[len(tiles)-tilesPerRow]

	image := stitchTiles(tiles, order, rotations, tilesPerRow)

	nonMonsterTiles := banishMonsters(image, tilesPerRow)

	return cornerMultiple, nonMonsterTiles
}

func banishMonsters(image map[point]bool, tilesPerRow int) int {
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
			return count
		}
	}
	panic("No monsters found")
}

func stitchTiles(tiles tileset, order []int, rotations []orientation, tilesPerRow int) map[point]bool {
	image := map[point]bool{}
	for tileRow := 0; tileRow < len(order)/tilesPerRow; tileRow++ {
		for tileColOffset := 0; tileColOffset < tilesPerRow; tileColOffset++ {
			imX, imY := tileColOffset*8, tileRow*8
			tile := tiles[order[tileRow*tilesPerRow+tileColOffset]]
			tileRotation := rotations[tileRow*tilesPerRow+tileColOffset]
			for y := 0; y < 8; y++ {
				rowPixels := tile.row(y+1, tileRotation)
				for x := 0; x < 8; x++ {
					if rowPixels&(1<<8>>x) > 0 {
						image[point{imX + x, imY + y}] = true
					}
				}
			}
		}
	}
	return image
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

func align(tiles tileset, tilesPerRow int) ([]int, []orientation) {
	tileIDs := make([]int, 0, len(tiles))
	for id := range tiles {
		tileIDs = append(tileIDs, id)
	}
	order, rotations, valid := tryArrangement(tiles, tilesPerRow, []int{}, tileIDs, []orientation{})
	if !valid {
		panic("No solution")
	}
	return order, rotations
}

func tryArrangement(tiles tileset, tilesPerRow int, order, remaining []int, rotations []orientation) ([]int, []orientation, bool) {
	if len(order) > 1 {
		curTileIndex := len(order) - 1
		curTileX := curTileIndex % tilesPerRow
		curTileY := curTileIndex / tilesPerRow
		curTile := tiles[order[curTileIndex]]
		if curTileY > 0 {
			aboveIndex := curTileIndex - tilesPerRow
			aboveTile := tiles[order[aboveIndex]]
			if aboveTile.row(9, rotations[aboveIndex]) != curTile.row(0, rotations[curTileIndex]) {
				return nil, nil, false
			}
		}
		if curTileX > 0 {
			leftIndex := curTileIndex - 1
			leftTile := tiles[order[leftIndex]]

			if leftTile.col(9, rotations[leftIndex]) != curTile.col(0, rotations[curTileIndex]) {
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
				finalOrder, finalRotations, valid := tryArrangement(tiles, tilesPerRow, newOrder, newRemaining, nextOrientations)
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

type tir struct {
	tile, index int
	rotation    orientation
}

var knownRows = map[tir]int{}
var knownCols = map[tir]int{}

type imageTile struct {
	id     int
	pixels map[orientedPoint]bool
}

func (t imageTile) row(rowNum int, rot orientation) int {
	if edge, found := knownRows[tir{t.id, rowNum, rot}]; found {
		return edge
	}
	row := 0
	start, step := 0, 1
	if (rot & flipX) > 0 {
		start, step = 9, -1
	}
	for x := start; x < 10 && x >= 0; x += step {
		row <<= 1
		if t.pixels[orientedPoint{x, rowNum, rot % orientMax}] {
			row++
		}
	}
	knownRows[tir{t.id, rowNum, rot}] = row
	return row
}

func (t imageTile) col(colNum int, rot orientation) int {
	if edge, found := knownCols[tir{t.id, colNum, rot}]; found {
		return edge
	}
	row := 0
	useCol := colNum
	if (rot & flipX) > 0 {
		useCol = 9 - useCol
	}
	for y := 0; y < 10; y++ {
		row <<= 1
		if t.pixels[orientedPoint{useCol, y, rot % orientMax}] {
			row++
		}
	}
	knownCols[tir{t.id, colNum, rot}] = row
	return row
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
