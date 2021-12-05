package main

import (
	"fmt"
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
	c := loadInput("input.txt")
	return c.bestCookie()
}

func part2() int {
	c := loadInput("input.txt")
	return c.bestCookieMeal()
}

type quantities map[string]int

type cupboard struct {
	tried       map[string]bool
	ingredients map[string]ingredientStats
}

func (c cupboard) bestCookie() int {
	return c.tryMix(quantities{}, false)
}

func (c cupboard) bestCookieMeal() int {
	return c.tryMix(quantities{}, true)
}

func (c cupboard) tryMix(used quantities, limitCalories bool) int {
	remaining := 100
	remainingIngredients := []string{}
	for ingredient := range c.ingredients {
		if val, found := used[ingredient]; found {
			remaining -= val
		} else {
			remainingIngredients = append(remainingIngredients, ingredient)
		}
	}
	if len(remainingIngredients) == 0 || remaining <= 0 {
		cookie := ingredientStats{}
		for name, quantity := range used {
			cookie = cookie.plus(c.ingredients[name].times(quantity))
		}
		return cookie.total(limitCalories)
	}
	if len(remainingIngredients) == 1 {
		newMix := quantities{}
		for name, val := range used {
			newMix[name] = val
		}
		newMix[remainingIngredients[0]] = remaining
		mixHash := fmt.Sprint(newMix)
		if c.tried[mixHash] {
			return 0
		}
		c.tried[mixHash] = true
		return c.tryMix(newMix, limitCalories)
	}
	best := 0
	for _, ingredient := range remainingIngredients {
		for i := 1; i <= remaining; i++ {

			newMix := quantities{}
			for name, val := range used {
				newMix[name] = val
			}
			newMix[ingredient] = i
			mixHash := fmt.Sprint(newMix)
			if c.tried[mixHash] {
				continue
			}
			score := c.tryMix(newMix, limitCalories)
			if score > best {
				best = score
			}
			c.tried[mixHash] = true
		}
	}
	return best
}

type ingredientStats struct {
	capacity, durability, flavor, texture, calories int
}

func (i ingredientStats) times(quantity int) ingredientStats {
	return ingredientStats{
		capacity:   quantity * i.capacity,
		durability: quantity * i.durability,
		flavor:     quantity * i.flavor,
		texture:    quantity * i.texture,
		calories:   quantity * i.calories,
	}
}

func (i ingredientStats) plus(j ingredientStats) ingredientStats {
	return ingredientStats{
		capacity:   j.capacity + i.capacity,
		durability: j.durability + i.durability,
		flavor:     j.flavor + i.flavor,
		texture:    j.texture + i.texture,
		calories:   j.calories + i.calories,
	}
}

func (i ingredientStats) total(limitCalories bool) int {
	if limitCalories && i.calories != 500 {
		return 0
	}
	if i.capacity < 0 ||
		i.durability < 0 ||
		i.flavor < 0 ||
		i.texture < 0 {
		return 0
	}
	return i.capacity * i.durability * i.flavor * i.texture
}

func loadInput(filename string) cupboard {
	c := cupboard{
		ingredients: map[string]ingredientStats{},
		tried:       map[string]bool{},
	}
	for _, line := range utils.ReadInputLines(filename) {
		parts := strings.Split(line, ":")
		name := parts[0]
		statInts := utils.GetInts(parts[1])
		c.ingredients[name] = ingredientStats{
			capacity:   statInts[0],
			durability: statInts[1],
			flavor:     statInts[2],
			texture:    statInts[3],
			calories:   statInts[4],
		}

	}
	return c
}

var benchmark = false
