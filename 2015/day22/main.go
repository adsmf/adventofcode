package main

import (
	"fmt"
	"io/ioutil"

	"github.com/adsmf/adventofcode/utils"
)

var benchmark = false
var debug = false
var debugCounter = 0
var printEvery = 1000

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	game := load("input.txt")
	return game.optimana()
}

func part2() int {
	game := load("input.txt")
	game.player.hploss = 1
	return game.optimana()
}

func load(filename string) gameData {
	inputBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	ints := utils.GetInts(string(inputBytes))

	g := gameData{
		playerNext: true,
		boss: entity{
			hitpoints:  ints[0],
			baseDamage: ints[1],
		},
		player: entity{
			hitpoints: 50,
			mana:      500,
		},
	}
	return g
}

type gameData struct {
	won         bool
	playerNext  bool
	player      entity
	boss        entity
	actions     string
	description string
}

func (g gameData) optimana() int {
	openRoutes := map[string]gameData{"": g}
	bestCost := utils.MaxInt

	iter := 0
	for len(openRoutes) > 0 {
		iter++
		nextOpenRoutes := map[string]gameData{}
		for prefix, seedGame := range openRoutes {
			startGame := seedGame.copy()
			startGame.applyEffects()
			if startGame.boss.hitpoints <= 0 {
				if startGame.player.manaSpent < bestCost {
					bestCost = startGame.player.manaSpent
				}
				continue
			}
			if startGame.playerNext {
				for _, option := range "mdspr" {
					route := fmt.Sprintf("%s%c", prefix, option)
					next, _, valid := startGame.playerTurn(option)
					if !valid {
						continue
					}
					if next.player.hitpoints <= 0 {
						continue
					}
					if next.player.manaSpent > bestCost {
						continue
					}
					nextOpenRoutes[route] = next
				}
			} else {
				route := fmt.Sprintf("%sb", prefix)
				next := startGame.bossTurn()
				if next.player.hitpoints > 0 {
					nextOpenRoutes[route] = next
				}
			}
		}
		openRoutes = nextOpenRoutes
	}

	return bestCost

}

func (g gameData) bossTurn() gameData {
	next := g.copy()
	next.actions += "b"
	bossDamage := g.boss.baseDamage - g.player.armour
	next.player.hitpoints -= bossDamage
	next.playerNext = true
	return next
}
func (g gameData) playerTurn(choice rune) (gameData, int, bool) {
	next := g.copy()
	next.player.hitpoints -= next.player.hploss
	cost := 0
	next.actions += string(choice)
	if debug && debugCounter%printEvery == 0 {
		fmt.Println(debugCounter, next.player.mana, next.player.manaSpent, next.actions)
	}
	debugCounter++
	valid := true
	switch choice {
	case 'm':
		cost = 53
		next.boss.hitpoints -= 4
	case 'd':
		cost = 73
		next.boss.hitpoints -= 2
		next.player.hitpoints += 2
	case 's':
		if next.player.shieldedFor > 0 {
			valid = false
			break
		}
		cost = 113
		next.player.shieldedFor = 6
	case 'p':
		if next.boss.poisonedFor > 0 {
			valid = false
			break
		}
		cost = 173
		next.boss.poisonedFor = 6
	case 'r':
		if next.player.rechargeFor > 0 {
			valid = false
			break
		}
		cost = 229
		next.player.rechargeFor = 5
	}
	if !valid || cost > next.player.mana {
		return g, 99999, false
	}
	next.player.manaSpent += cost
	next.player.mana -= cost
	next.playerNext = false
	return next, cost, true
}
func (g *gameData) applyEffects() {
	g.description = ""
	if g.player.shieldedFor > 0 {
		g.player.shieldedFor--
		g.description += fmt.Sprintf(" shielded (%d)", g.player.shieldedFor)
		g.player.armour = 7
	} else {
		g.player.armour = 0
	}
	if g.player.rechargeFor > 0 {
		g.player.rechargeFor--
		g.player.mana += 101
		g.description += fmt.Sprintf(" recharge (%d)", g.player.rechargeFor)
	}
	if g.boss.poisonedFor > 0 {
		g.boss.poisonedFor--
		g.boss.hitpoints -= 3
		g.description += fmt.Sprintf(" poisoned (%d)", g.boss.poisonedFor)
	}
}
func (g gameData) copy() gameData {
	copy := gameData{
		won:         g.won,
		playerNext:  g.playerNext,
		player:      g.player,
		boss:        g.boss,
		actions:     g.actions,
		description: g.description,
	}
	return copy
}

type entity struct {
	hploss      int
	hitpoints   int
	mana        int
	armour      int
	manaSpent   int
	baseDamage  int
	shieldedFor int
	poisonedFor int
	rechargeFor int
}
