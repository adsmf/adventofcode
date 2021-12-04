package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	draws, boards := loadInput()
	p1, p2 := getScores(draws, boards)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func getScores(draws []int, boards []bingoBoard) (int, int) {
	allWinSets := generateWinSets(boards)
	winningBoard, lastBoard := bingoBoard{}, bingoBoard{}
	winningDrawNum, lastDrawNum := -1, -1
	winners := make(set, len(boards))
	for draw, num := range draws {
		for boardNum, winSets := range allWinSets {
			for _, winSet := range winSets {
				if winSet[num] {
					delete(winSet, num)
				}
				if len(winSet) == 0 {
					winners[boardNum] = true
					if len(winners) == 1 {
						winningBoard = boards[boardNum]
						winningDrawNum = draw
					} else if len(winners) == len(boards) {
						lastBoard = boards[boardNum]
						lastDrawNum = draw
					}
					allWinSets[boardNum] = []set{}
					break
				}
			}
			if len(winners) == len(boards) {
				break
			}
		}
		if len(winners) == len(boards) {
			break
		}
	}
	winnerBoardScore := winningBoard.unmarkedAfter(draws[:winningDrawNum+1])
	lastBoardScore := lastBoard.unmarkedAfter(draws[:lastDrawNum+1])
	return winnerBoardScore * draws[winningDrawNum], lastBoardScore * draws[lastDrawNum]
}

func generateWinSets(boards []bingoBoard) [][]set {
	sets := make([][]set, len(boards))
	for boardID, board := range boards {
		sets[boardID] = make([]set, 10)
		for i := 0; i < 10; i++ {
			sets[boardID][i] = make(set, 5)
		}
		for x := 0; x < 5; x++ {
			for y := 0; y < 5; y++ {
				num := board[point(x, y)]
				sets[boardID][x][num] = true
				sets[boardID][5+y][num] = true
			}
		}
	}
	return sets
}

func loadInput() ([]int, []bingoBoard) {
	blocks := strings.Split(input, "\n\n")

	drawString := blocks[0]
	boardStrings := blocks[1:]

	draws := utils.GetInts(drawString)

	boards := make([]bingoBoard, 0, len(boardStrings))
	for _, boardString := range boardStrings {
		board := make(bingoBoard, 25)
		for row, line := range strings.Split(boardString, "\n") {
			for col, number := range utils.GetInts(line) {
				board[point(row, col)] = number
			}
		}
		boards = append(boards, board)
	}

	return draws, boards
}

func point(x, y int) int {
	return int(x*5 + y)
}

type set map[int]bool

type bingoBoard []int

func (b bingoBoard) unmarkedAfter(draws []int) int {
	drawn := make(set, len(draws))
	for _, draw := range draws {
		drawn[draw] = true
	}
	total := 0
	for _, num := range b {
		if !drawn[num] {
			total += num
		}
	}
	return total
}

var benchmark = false
