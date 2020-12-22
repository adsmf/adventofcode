package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	hands := load("input.txt")
	for len(hands[0]) > 0 && len(hands[1]) > 0 {
		playRound(hands)
	}
	var winningHand playerHand
	if len(hands[0]) > 0 {
		winningHand = hands[0]
	} else {
		winningHand = hands[1]
	}
	return scoreHand(winningHand)
}

func scoreHand(hand playerHand) int {
	score := 0
	for value, i := len(hand), 0; i < len(hand); value, i = value-1, i+1 {
		score += value * hand[i]
	}
	return score
}

func playRound(hands gameState) {
	if hands[0][0] > hands[1][0] {
		hands[0], hands[1] = append(hands[0][1:], hands[0][0], hands[1][0]), hands[1][1:]
	} else {
		hands[1], hands[0] = append(hands[1][1:], hands[1][0], hands[0][0]), hands[0][1:]
	}
}

func part2() int {
	hands := load("input.txt")

	winner := playRecursiveGame(hands)

	return scoreHand(hands[winner])
}

func playRecursiveGame(hands gameState) int {
	resursiveStates := map[string]bool{}
	var p1card, p2card int
	for {
		if _, found := resursiveStates[hands.hash()]; found {
			return 0
		}
		resursiveStates[hands.hash()] = true
		if len(hands[0]) == 0 {
			return 1
		}
		if len(hands[1]) == 0 {
			return 0
		}
		p1card, hands[0] = hands[0][0], hands[0][1:]
		p2card, hands[1] = hands[1][0], hands[1][1:]

		if p1card <= len(hands[0]) && p2card <= len(hands[1]) {
			subDeck := gameState{
				append(playerHand{}, hands[0][:p1card]...),
				append(playerHand{}, hands[1][:p2card]...),
			}

			winner := playRecursiveGame(subDeck)
			if winner == 0 {
				hands[0] = append(hands[0], p1card, p2card)
			} else {
				hands[1] = append(hands[1], p2card, p1card)
			}
		} else {
			if p1card > p2card {
				hands[0] = append(hands[0], p1card, p2card)
			} else {
				hands[1] = append(hands[1], p2card, p1card)
			}
		}
	}
}

type gameState []playerHand

func (g gameState) hash() string { return fmt.Sprintf("%v", g) }

type playerHand []int

func load(filename string) gameState {
	inputBytes, _ := ioutil.ReadFile(filename)
	hands := make(gameState, 2)
	for player, block := range strings.Split(string(inputBytes), "\n\n") {
		hands[player] = utils.GetInts(block)[1:]
	}
	return hands
}

var benchmark = false
