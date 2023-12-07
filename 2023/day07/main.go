package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1 := calcWinnings(false)
	p2 := calcWinnings(true)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func calcWinnings(jWild bool) int {
	hands := []handInfo{}
	utils.EachLine(input, func(index int, line string) (done bool) {
		bid, _ := strconv.Atoi(line[6:])
		hand := handInfo{
			cards: line[:5],
			bid:   bid,
		}
		hands = append(hands, hand)
		return false
	})
	localRanks := cardRanks
	if jWild {
		localRanks['J'] = -1
	}
	sort.Slice(hands, func(i, j int) bool {
		s1, s2 := hands[i].score(jWild), hands[j].score(jWild)
		if s1 != s2 {
			return s1 < s2
		}
		for k := 0; k < 5; k++ {
			c1, c2 := hands[i].cards[k], hands[j].cards[k]
			if c1 != c2 {
				return localRanks[c1] < localRanks[c2]
			}

		}
		return false
	})
	winnings := 0
	for i := 0; i < len(hands); i++ {
		winnings += (i + 1) * hands[i].bid
	}
	return winnings
}

type handInfo struct {
	cards  string
	bid    int
	cached handType
}

func (h *handInfo) score(jWild bool) handType {
	if h.cached != handUnscored {
		return h.cached
	}
	cardCounts := [14]int{}
	jokers := 0
	for _, card := range h.cards {
		cardCounts[cardRanks[byte(card)]]++
	}
	countCards := [6][]int{}
	for card, count := range cardCounts {
		if card == 9 && jWild {
			jokers += count
		} else {
			countCards[count] = append(countCards[count], card)
		}
	}
	for j := 0; j < jokers; j++ {
		for i := len(countCards) - 1; i >= 0; i-- {
			if len(countCards[i]) == 0 {
				continue
			}
			countCards[i], countCards[i+1] = countCards[i][1:], []int{countCards[i][0]}
			break
		}
	}
	rank := handUnscored
	switch {
	case len(countCards[5]) == 1:
		rank = handFiveOfKind
	case len(countCards[4]) == 1:
		rank = handFourOfKind
	case len(countCards[3]) == 1 && len(countCards[2]) == 1:
		rank = handFullHouse
	case len(countCards[3]) == 1:
		rank = handThreeOfKind
	case len(countCards[2]) == 2:
		rank = handTwoPair
	case len(countCards[2]) == 1:
		rank = handOnePair
	default:
		rank = handHighCard
	}
	h.cached = rank
	return rank
}

var cardRanks = [...]int{'2': 0, '3': 1, '4': 2, '5': 3, '6': 4, '7': 5, '8': 6, '9': 7, 'T': 8, 'J': 9, 'Q': 10, 'K': 11, 'A': 12}

type handType int

const (
	handUnscored handType = iota

	handHighCard
	handOnePair
	handTwoPair
	handThreeOfKind
	handFullHouse
	handFourOfKind
	handFiveOfKind
)

var benchmark = false
