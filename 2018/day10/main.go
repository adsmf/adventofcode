package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode2018/utils"
)

func main() {
	lines := utils.ReadInputLines("input.txt")
	stars := parseLines(lines)
	bestArea := -1
	bestAreaTick := 0
	var bestPos []vector
	for tick := 0; tick < 100000; tick++ {
		pos, width, height := getStarPos(stars, tick)
		area := width * height
		if bestArea == -1 || area < bestArea {
			bestArea = area
			bestPos = pos
			bestAreaTick = tick
		}
	}
	showStars(bestPos)
	fmt.Printf("Would take: %d\n", bestAreaTick)
}

func getStarPos(stars []star, tick int) ([]vector, int, int) {
	curPositions := []vector{}
	for _, star := range stars {
		curPos := vector{
			x: star.position.x + tick*star.velocity.x,
			y: star.position.y + tick*star.velocity.y,
		}
		curPositions = append(curPositions, curPos)
	}
	minX := curPositions[0].x
	maxX := curPositions[0].x
	minY := curPositions[0].y
	maxY := curPositions[0].y
	for _, star := range curPositions {
		if star.x < minX {
			minX = star.x
		}
		if star.y < minY {
			minY = star.y
		}
		if star.x > maxX {
			maxX = star.x
		}
		if star.y > maxY {
			maxY = star.y
		}
	}
	width := maxX - minX + 1
	height := maxY - minY + 1
	return curPositions, width, height
}

func showStars(positions []vector) {
	maxX := 0
	maxY := 0
	for _, pos := range positions {
		if pos.x > maxX {
			maxX = pos.x
		}
		if pos.y > maxY {
			maxY = pos.y
		}
	}
	fmt.Printf("Showing stars (%d x %d)\n", maxX, maxY)
	grid := make([][]string, maxY+1)
	for row := 0; row <= maxY; row++ {
		grid[row] = make([]string, maxX+1)
		for col := 0; col <= maxX; col++ {
			grid[row][col] = " "
		}
	}
	for _, star := range positions {
		if star.x >= 0 && star.x <= maxX {
			if star.y >= 0 && star.y <= maxY {
				grid[star.y][star.x] = "#"
			}
		}
	}

	for row := 0; row <= maxY; row++ {
		fmt.Println(strings.Join(grid[row], ""))
	}

}

func parseLines(lines []string) []star {
	stars := []star{}
	re := regexp.MustCompile("position=<[ ]*(-?\\d+),[ ]*(-?\\d+)> velocity=<[ ]*(-?\\d+),[ ]*(-?\\d+)>")
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		posX, _ := strconv.Atoi(match[1])
		posY, _ := strconv.Atoi(match[2])
		velX, _ := strconv.Atoi(match[3])
		velY, _ := strconv.Atoi(match[4])
		newStar := star{
			position: vector{posX, posY},
			velocity: vector{velX, velY},
		}
		stars = append(stars, newStar)

	}
	return stars
}

type star struct {
	position vector
	velocity vector
}

type vector struct {
	x int
	y int
}
