package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	lines := utils.ReadInputLines("input.txt")
	claims := genClaims(lines)
	fab := makeFabric()
	fab = landgrab(fab, claims)

	part1(fab)
	part2(claims, fab)
}

type claim struct {
	ID       int
	Position position
	Size     size
}
type position struct {
	X int
	Y int
}
type size struct {
	Height int
	Width  int
}
type fabric [][][]claim

func genClaims(lines []string) []claim {
	claims := []claim{}
	for _, line := range lines {
		lineParts := strings.Split(line, " ")
		id, _ := strconv.Atoi(strings.TrimPrefix(lineParts[0], "#"))
		posParts := strings.Split(strings.TrimSuffix(lineParts[2], ":"), ",")
		x, _ := strconv.Atoi(posParts[0])
		y, _ := strconv.Atoi(posParts[1])

		sizeParts := strings.Split(lineParts[3], "x")
		width, _ := strconv.Atoi(sizeParts[0])
		height, _ := strconv.Atoi(sizeParts[1])

		claims = append(claims, claim{
			ID: id,
			Position: position{
				X: x,
				Y: y,
			},
			Size: size{
				Width:  width,
				Height: height,
			},
		})
	}
	return claims
}

func makeFabric() fabric {
	fab := make([][][]claim, 1000)
	for i := range fab {
		fab[i] = make([][]claim, 1000)
	}
	return fab
}

func landgrab(fab fabric, claims []claim) fabric {
	for _, claim := range claims {
		for col := claim.Position.X; col < claim.Position.X+claim.Size.Width; col++ {
			for row := claim.Position.Y; row < claim.Position.Y+claim.Size.Height; row++ {
				fab[col][row] = append(fab[col][row], claim)
			}
		}
	}
	return fab
}

func part1(fab fabric) {
	contested := 0
	for _, fabricRow := range fab {
		for _, claims := range fabricRow {
			if len(claims) > 1 {
				contested++
			}
		}
	}
	fmt.Printf("Contested squares: %d\n", contested)
}

func part2(claims []claim, fab fabric) {
	for _, claim := range claims {
		undisputed := true
		for col := claim.Position.X; col < claim.Position.X+claim.Size.Width; col++ {
			for row := claim.Position.Y; row < claim.Position.Y+claim.Size.Height; row++ {
				if len(fab[col][row]) != 1 {
					undisputed = false
					break
				}
			}
			if !!!undisputed {
				break
			}
		}
		if undisputed {
			fmt.Printf("Claim undisputed: %+v\n", claim)
			break
		}
	}
}
