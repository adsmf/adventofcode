package main

import (
	"bytes"
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
	winner := playGame(hands, false)
	return scoreHand(hands[winner])
}

func part2() int {
	hands := load("input.txt")
	winner := playGame(hands, true)
	return scoreHand(hands[winner])
}

func playGame(hands gameState, recursive bool) int {
	resursiveStates := map[stateHashString]struct{}{}
	var p1card, p2card byte
	for {
		hash := hands.hashCombine()
		if _, found := resursiveStates[hash]; found {
			return 0
		}
		resursiveStates[hash] = struct{}{}
		if len(hands[0]) == 0 {
			return 1
		}
		if len(hands[1]) == 0 {
			return 0
		}
		p1card, hands[0] = hands[0][0], hands[0][1:]
		p2card, hands[1] = hands[1][0], hands[1][1:]

		if recursive && int(p1card) <= len(hands[0]) && int(p2card) <= len(hands[1]) {
			subDeck := gameState{
				append(playerHand{}, hands[0][:p1card]...),
				append(playerHand{}, hands[1][:p2card]...),
			}

			winner := playGame(subDeck, true)
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

func scoreHand(hand playerHand) int {
	score, value := 0, len(hand)
	for _, card := range hand {
		score += value * int(card)
		value--
	}
	return score
}

type gameState []playerHand

type stateHashInt struct {
	p1score, p2score int
}

type stateHashString string

func (g gameState) hashScore() stateHashInt {
	return stateHashInt{scoreHand(g[0]), scoreHand(g[1])}
}

func (g gameState) hashSprint() stateHashString {
	return stateHashString(fmt.Sprint(g))
}

func (g gameState) hashCombine() stateHashString {
	return stateHashString(g[0]) + "|" + stateHashString(g[1])
}

func (g gameState) hashFromByteSlice() stateHashString {
	hashBytes := make([]byte, 0, 1+len(g[0])+len(g[1]))
	for _, card := range g[0] {
		hashBytes = append(hashBytes, byte(card))
	}
	hashBytes = append(hashBytes, byte(0))
	for _, card := range g[1] {
		hashBytes = append(hashBytes, byte(card))
	}
	return stateHashString(hashBytes)
}

func (g gameState) hashByteBuffer() stateHashString {
	var buf bytes.Buffer
	for _, card := range g[0] {
		buf.WriteByte(byte(card))
	}
	buf.WriteByte(0)
	for _, card := range g[1] {
		buf.WriteByte(byte(card))
	}
	return stateHashString(buf.String())
}

type playerHand []byte

func load(filename string) gameState {
	inputBytes, _ := ioutil.ReadFile(filename)
	hands := make(gameState, 2)
	for player, block := range strings.Split(string(inputBytes), "\n\n") {
		cards := utils.GetInts(block)[1:]
		cardBytes := make(playerHand, len(cards))
		for index, card := range cards {
			cardBytes[index] = byte(card)
		}
		hands[player] = cardBytes
	}
	return hands
}

var benchmark = false
