package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/adsmf/adventofcode2019/utils"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	v := loadMap("input.txt")
	v.reduce()
	return v.collectKeys()
}

func part2() int {
	v := loadMap("input2.txt")
	v.reduce()
	return v.collectKeys()
}

type vault struct {
	vault      map[point]tile
	cache      map[string]int
	keys       map[int]bool
	minX, minY int
	maxX, maxY int
}

func (v *vault) collectKeys() int {
	entrances := []point{}

	for pos, t := range v.vault {
		if t.tileType == tileTypeStart {
			entrances = append(entrances, pos)
		}
	}

	keys := keyring("")

	crawlSteps := v.crawl(entrances, keys)
	return crawlSteps
}

func (v *vault) crawl(entrances []point, keys keyring) int {
	cacheHash := fmt.Sprintf("%v/%v", keys, entrances)

	if val, found := v.cache[cacheHash]; found {
		return val
	}
	newKeys := v.findNextAll(entrances, keys)
	if len(newKeys) == 0 {
		v.cache[cacheHash] = 0
		return 0
	}

	try := []int{}
	for sym, info := range newKeys {
		newKeyring := keys.with(sym)
		var newPoints []point
		for entIdx, entPos := range entrances {
			if entIdx == info.entIndex {
				newPoints = append(newPoints, info.pos)
			} else {
				newPoints = append(newPoints, entPos)
			}
		}
		crawlSteps := v.crawl(newPoints, newKeyring)
		try = append(try, info.steps+crawlSteps)
	}
	sort.Ints(try)
	v.cache[cacheHash] = try[0]

	return try[0]
}

func (v *vault) findNextAll(entrances []point, keys keyring) keyMap {
	foundKeys := keyMap{}
	for idx, entrance := range entrances {
		for sym, info := range v.findNext(entrance, keys) {
			foundKeys[sym] = keyInfo{
				steps:    info.steps,
				pos:      info.pos,
				entIndex: idx,
			}
		}
	}
	return foundKeys
}

func (v *vault) findNext(entrance point, keys keyring) keyMap {
	foundKeys := keyMap{}

	searchArea := []point{entrance}
	steps := map[point]int{
		entrance: 0,
	}

	for len(searchArea) > 0 {
		next := searchArea[0]
		searchArea = searchArea[1:]
		for _, pos := range next.neighbours() {
			t := v.vault[pos]
			if t.tileType == tileTypeUnkown ||
				t.tileType == tileTypeWall {
				continue
			}
			if _, found := steps[pos]; found {
				continue
			}
			steps[pos] = steps[next] + 1

			if t.tileType == tileTypeDoor && !keys.has(t.id) {
				continue
			}

			if t.tileType == tileTypeKey {
				if keys.has(t.id) {
					searchArea = append(searchArea, pos)
				} else {
					foundKeys[t.id] = keyInfo{
						steps: steps[pos],
						pos:   pos,
					}
				}
				continue
			}
			searchArea = append(searchArea, pos)
		}
	}

	return foundKeys
}

type keyMap map[int]keyInfo

type keyInfo struct {
	steps    int
	pos      point
	entIndex int
}

func (v *vault) set(pos point, val tile) {
	val.pos = pos
	val.vault = v
	v.vault[pos] = val
	if v.minX > pos.x {
		v.minX = pos.x
	}
	if v.maxX < pos.x {
		v.maxX = pos.x
	}
	if v.minY > pos.y {
		v.minY = pos.y
	}
	if v.maxY < pos.y {
		v.maxY = pos.y
	}
}

func (v vault) String() string {
	newString := ""
	for y := v.minY; y <= v.maxY; y++ {
		for x := v.minX; x <= v.maxX; x++ {
			newString += fmt.Sprintf("%v", v.vault[point{x, y}])
		}
		newString += fmt.Sprintln()
	}
	return newString
}

func (v *vault) reduce() {
	for {
		simplifyTiles := []point{}
		for pos, t := range v.vault {
			if t.tileType != tileTypeEmpty {
				continue
			}
			count := 0
			for _, n := range pos.neighbours() {
				if v.vault[n].tileType == tileTypeWall {
					count++
				}
			}
			if count >= 3 {
				simplifyTiles = append(simplifyTiles, pos)
			}
		}
		if len(simplifyTiles) == 0 {
			return
		}
		for _, pos := range simplifyTiles {
			existing := v.vault[pos]
			existing.tileType = tileTypeWall
			v.set(pos, existing)
		}
	}
}

func loadMap(filename string) vault {
	newVault := vault{
		vault: map[point]tile{},
		cache: map[string]int{},
		keys:  map[int]bool{},
	}

	lines := utils.ReadInputLines(filename)
	for y, line := range lines {
		for x, char := range line {
			pos := point{x, y}
			switch {
			case char == '@':
				newVault.set(pos, tile{tileType: tileTypeStart})
			case char == '.':
				newVault.set(pos, tile{tileType: tileTypeEmpty})
			case char == '#':
				newVault.set(pos, tile{tileType: tileTypeWall})
			case 'a' <= char && char <= 'z':
				id := int(char)
				newVault.set(pos, tile{tileType: tileTypeKey, id: id})
				newVault.keys[id] = true
			case 'A' <= char && char <= 'Z':
				id := int(char - 'A' + 'a')
				newVault.set(pos, tile{tileType: tileTypeDoor, id: id})
			}
		}
	}
	return newVault
}

type tile struct {
	tileType tileType
	id       int
	pos      point
	vault    *vault
}

func (t tile) String() string {
	switch t.tileType {
	case tileTypeEmpty:
		return " "
	case tileTypeWall:
		return "█"
	case tileTypeDoor:
		return string(t.id - 'a' + 'A')
	case tileTypeKey:
		return string(t.id)
	case tileTypeStart:
		return "s"
	case tileTypeUnkown:
		return "?"
	default:
		return "!"
	}
}

type tileType int

const (
	tileTypeUnkown tileType = iota
	tileTypeEmpty
	tileTypeWall
	tileTypeDoor
	tileTypeKey
	tileTypeStart
)

type point struct {
	x, y int
}

func (p point) neighbours() []point {
	return []point{
		point{p.x - 1, p.y},
		point{p.x + 1, p.y},
		point{p.x, p.y - 1},
		point{p.x, p.y + 1},
	}
}

type keyring string

func (k keyring) has(key int) bool {
	return strings.Contains(
		string(k),
		fmt.Sprintf("%c", key),
	)
}

func (k keyring) with(key int) keyring {
	keyStr := fmt.Sprintf("%c", key)
	if strings.Contains(string(k), keyStr) {
		return k
	}
	newRing := k + keyring(keyStr)
	newRune := keyRune(newRing)

	sort.Sort(newRune)

	return keyring(newRune)
}

type keyRune []rune

func (k keyRune) Len() int           { return len(k) }
func (k keyRune) Less(i, j int) bool { return k[i] < k[j] }
func (k keyRune) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }
