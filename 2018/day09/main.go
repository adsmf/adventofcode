package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Printf("Part 1: %d\n", playMarbles(486, 70833))
	fmt.Printf("Part 2: %d\n", playMarbles(486, 7083300))
}

func playMarbles(players, marbles int) int {
	circle := list.New()
	circle.PushFront(0)
	circleOffset := circle.Front()

	playerScores := make([]int, players)
	curPlayer := 0

	for nextMarble := 1; nextMarble <= marbles; nextMarble++ {
		curPlayer = (curPlayer + 1) % players
		if (nextMarble % 23) == 0 {
			playerScores[curPlayer] += nextMarble
			removeOffset := getOffsetCCW(circle, 7, circleOffset)
			circleOffset = removeOffset.Next()
			if circleOffset == nil {
				circleOffset = circle.Front()
			}
			playerScores[curPlayer] += (removeOffset.Value).(int)
			circle.Remove(removeOffset)
			continue
		}
		insertPos := getOffsetCW(circle, 1, circleOffset)
		circle.InsertAfter(nextMarble, insertPos)
		circleOffset = insertPos.Next()
	}
	bestScore := 0
	for _, score := range playerScores {
		if score > bestScore {
			bestScore = score
		}
	}
	return bestScore
}

func getOffsetCW(circle *list.List, targetOffset int, curPos *list.Element) *list.Element {
	for curOffset := 0; curOffset < targetOffset; curOffset++ {
		curPos = curPos.Next()
		if curPos == nil {
			curPos = circle.Front()
		}
	}
	return curPos
}

func getOffsetCCW(circle *list.List, targetOffset int, curPos *list.Element) *list.Element {
	for curOffset := 0; curOffset < targetOffset; curOffset++ {
		curPos = curPos.Prev()
		if curPos == nil {
			curPos = circle.Back()
		}
	}
	return curPos
}
