package main

import (
	"fmt"
	"strconv"
	"strings"
)

var logger func(string, ...interface{})

func main() {
	input := 290431
	part1(setup(), input)
	part2(setup(), input)
}

func setup() cookState {
	elves := workers{}
	board := scoreboard{3, 7}

	elves = append(
		elves,
		0,
		1,
	)
	state := cookState{
		elves: elves,
		board: board,
	}
	return state
}

func part1(state cookState, input int) {
	fmt.Println(nextTen(state, 290431))
}

func part2(state cookState, input int) {
	target := fmt.Sprint(input)
	fmt.Println(firstAppears(state, target))
}

func firstAppears(state cookState, target string) int {
	curTick := 0
	for {
		state = tick(state)
		curTick++
		if len(state.board) < len(target)+1 {
			continue
		}
		lastXPlus1 := state.board.LastScores(len(target) + 1)
		if lastXPlus1[1:] == target {
			return len(state.board) - len(target)
		} else if lastXPlus1[:len(target)] == target {
			return len(state.board) - len(target) - 1
		}
	}
}

func nextTen(state cookState, startAt int) string {
	for len(state.board) < startAt+10 {
		state = tick(state)
	}
	return state.board.LastScores(10)
}

func tick(state cookState) cookState {
	newBoard := append(state.board, genRecipes(state)...)

	if logger != nil {
		logger("Old state: %s", state.AsString())
	}

	state.board = newBoard

	newElves := moveElves(state)
	state.elves = newElves

	if logger != nil {
		logger("New state:%s\n%s\n%s", state.AsString(), newBoard.AsString(), newElves.AsString())
	}
	return state
}

func genRecipes(state cookState) []recipe {
	nextScore := 0
	for _, elf := range state.elves {
		nextScore += int(state.board[elf])
	}
	scoreStr := strconv.Itoa(nextScore)
	newRecipes := []recipe{}
	for _, digit := range scoreStr {
		newRecipe, _ := strconv.Atoi(string(digit))
		newRecipes = append(newRecipes, recipe(newRecipe))
	}

	return newRecipes
}

func moveElves(state cookState) workers {
	newElves := workers{}
	for _, curElf := range state.elves {
		if logger != nil {
			logger(
				"(%d + %d) %% %d = %d",
				curElf,
				elf(1+int(state.board[curElf])),
				len(state.board),
				(int(curElf)+1+int(state.board[curElf]))%len(state.board),
			)
		}
		nextPos := (int(curElf) + 1 + int(state.board[curElf])) % len(state.board)
		if logger != nil {
			logger("nextPos: %d", nextPos)
		}
		newElves = append(
			newElves,
			elf(nextPos),
		)
	}
	return newElves
}

type cookState struct {
	board scoreboard
	elves workers
}

func (c *cookState) AsString() string {
	retString := ""
	for recipieID, score := range c.board {
		curTask := false
		for elfID, elf := range c.elves {
			if int(elf) == recipieID {
				curTask = true
				switch elfID {
				case 0:
					retString += fmt.Sprintf("(%d)", score)
				case 1:
					retString += fmt.Sprintf("[%d]", score)
				default:
					retString += fmt.Sprintf("?%d?", score)
				}
			}
		}
		if !!!curTask {
			retString += fmt.Sprintf(" %d ", score)
		}
	}
	return retString
}

type recipe int
type scoreboard []recipe

func (s scoreboard) AsString() string {
	scoreStrings := []string{}
	for _, score := range s {
		scoreStrings = append(scoreStrings, fmt.Sprint(score))
	}
	scores := strings.Join(scoreStrings, ",")
	return fmt.Sprint(len(s), " recipies: ", scores)
}
func (s scoreboard) LastScores(numScores int) string {
	lastX := s[len(s)-numScores:]
	result := ""
	for _, score := range lastX {
		result = fmt.Sprint(result, score)
	}
	return result
}

type elf int
type workers []elf

func (w workers) AsString() string {
	curRecipies := []string{}
	for _, cur := range w {
		curRecipies = append(curRecipies, fmt.Sprint(cur))
	}
	tasks := strings.Join(curRecipies, ",")
	return fmt.Sprint(len(w), " elves working on ", tasks)
}
