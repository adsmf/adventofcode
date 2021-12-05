package main

import (
	"fmt"
	"math"
)

var bossStats = entity{hitPoints: 109, damage: 8, armor: 2}

var weapons = []item{
	item{"Dagger", 8, 4, 0},
	item{"Shortsword", 10, 5, 0},
	item{"Warhammer", 25, 6, 0},
	item{"Longsword", 40, 7, 0},
	item{"Greataxe", 74, 8, 0},
}

var armours = []item{
	item{"None", 0, 0, 0},
	item{"Leather", 13, 0, 1},
	item{"Chainmail", 31, 0, 2},
	item{"Splintmail", 53, 0, 3},
	item{"Bandedmail", 75, 0, 4},
	item{"Platemail", 102, 0, 5},
}

var rings = []item{
	item{"None 1", 0, 0, 0},
	item{"None 2", 0, 0, 0},
	item{"Damage +1", 25, 1, 0},
	item{"Damage +2", 50, 2, 0},
	item{"Damage +3", 100, 3, 0},
	item{"Defense +1", 20, 0, 1},
	item{"Defense +2", 40, 0, 2},
	item{"Defense +3", 80, 0, 3},
}

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	return cheapWin()
}

func part2() int {
	return expensiveLoss()
}

func cheapWin() int {
	bestCost := -1
	for _, weapon := range weapons {
		for _, armour := range armours {
			for _, ring1 := range rings {
				for _, ring2 := range rings {
					if ring1 == ring2 {
						continue
					}
					player := entity{hitPoints: 100}
					player.damage += weapon.damage + ring1.damage + ring2.damage
					player.armor += armour.armor + ring1.armor + ring2.armor
					if playerWinsBattle(player, bossStats) {
						cost := weapon.cost + armour.cost + ring1.cost + ring2.cost
						if bestCost == -1 || cost < bestCost {
							bestCost = cost
						}
					}
				}
			}
		}
	}
	return bestCost
}

func expensiveLoss() int {
	worstCost := 0
	for _, weapon := range weapons {
		for _, armour := range armours {
			for _, ring1 := range rings {
				for _, ring2 := range rings {
					if ring1 == ring2 {
						continue
					}
					player := entity{hitPoints: 100}
					player.damage += weapon.damage + ring1.damage + ring2.damage
					player.armor += armour.armor + ring1.armor + ring2.armor
					if !playerWinsBattle(player, bossStats) {
						cost := weapon.cost + armour.cost + ring1.cost + ring2.cost
						if cost > worstCost {
							worstCost = cost
						}
					}
				}
			}
		}
	}
	return worstCost
}

func playerWinsBattle(player, boss entity) bool {
	perPlayerHit := player.damage - boss.armor
	if perPlayerHit < 1 {
		perPlayerHit = 1
	}
	perBosttHit := boss.damage - player.armor
	if perBosttHit < 1 {
		perBosttHit = 1
	}
	bossHits := math.Ceil(float64(boss.hitPoints) / float64(perPlayerHit))
	playerHits := math.Ceil(float64(player.hitPoints) / float64(perBosttHit))
	return bossHits <= playerHits
}

type entity struct {
	hitPoints int
	damage    int
	armor     int
}

type item struct {
	name   string
	cost   int
	damage int
	armor  int
}

var benchmark = false
