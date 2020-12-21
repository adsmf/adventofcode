package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1, p2 := foodInspector()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %s\n", p2)
	}
}

func foodInspector() (int, string) {
	recipes := load("input.txt")

	possibleAllergens := map[string]map[string]bool{}
	allIngredients := map[string]int{}
	for _, recipe := range recipes {
		for ing := range recipe.ingredients {
			allIngredients[ing]++
		}
		for allergen := range recipe.allergens {
			if possibleAllergens[allergen] == nil {
				possibleAllergens[allergen] = make(map[string]bool, len(recipe.ingredients))
				for ing := range recipe.ingredients {
					possibleAllergens[allergen][ing] = true
				}
			} else {
				for ing := range possibleAllergens[allergen] {
					if _, found := recipe.ingredients[ing]; !found {
						delete(possibleAllergens[allergen], ing)
					}
				}
			}
		}
	}

	knownAllergens := map[string]string{}
	knownBadIngredients := map[string]string{}
	for len(possibleAllergens) > 0 {
		for allergen, possible := range possibleAllergens {
			if len(possible) == 1 {
				ingredient := ""
				for ing := range possible {
					ingredient = ing
				}
				knownAllergens[allergen] = ingredient
				knownBadIngredients[ingredient] = allergen
				for otherAllergen := range possibleAllergens {
					delete(possibleAllergens[otherAllergen], ingredient)
				}
				delete(possibleAllergens, allergen)
			}
		}
	}

	part1 := 0
	for ing, count := range allIngredients {
		if _, found := knownBadIngredients[ing]; !found {
			part1 += count
		}
	}
	allergens := []string{}
	for al := range knownAllergens {
		allergens = append(allergens, al)
	}
	sort.Strings(allergens)
	badIngredients := []string{}
	for _, al := range allergens {
		badIngredients = append(badIngredients, knownAllergens[al])
	}
	return part1, strings.Join(badIngredients, ",")
}

func load(filename string) []recipe {
	recipes := []recipe{}
	for _, line := range utils.ReadInputLines(filename) {
		lineIngredients := map[string]bool{}
		lineAllergens := map[string]bool{}
		parts := strings.Split(line[:len(line)-1], " (contains ")
		for _, ingredient := range strings.Split(parts[0], " ") {
			lineIngredients[ingredient] = true
		}
		for _, allergen := range strings.Split(parts[1], ", ") {
			lineAllergens[allergen] = true
		}
		recipes = append(recipes, recipe{lineIngredients, lineAllergens})
	}
	return recipes
}

type recipe struct {
	ingredients map[string]bool
	allergens   map[string]bool
}

var benchmark = false
