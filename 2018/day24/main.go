package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
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
	body := load("input.txt")
	_, _, remaining := body.war()
	return remaining
}

func part2() int {
	boostMin, boostMax := 0, 1
	boost := 1
	haveMax := false
	for iter := 0; iter <= 30; iter++ {
		body := load("input.txt")
		for id := range body.immuneSystem {
			body.immuneSystem[id].damage += boost
		}
		stalled, survive, remaining := body.war()
		if stalled {
			boost++
			continue
		}
		if !survive {
			if !haveMax {
				haveMax = true
				boostMax = boost
				boostMin = boost / 2
			} else {
				if boostMax == boostMin+1 {
					return remaining
				}
				boostMax = boost
			}
		} else {
			if !haveMax {
				boost *= 2
				continue
			}
			boostMin = boost + 1
		}
		boost = boostMin + (boostMax-boostMin)/2
	}
	return -1
}

func (r reindeer) war() (bool, bool, int) {
	for {
		stalled := r.fight()
		if stalled {
			return true, false, -1
		}

		infHasUnits := false
		immHasUnits := false

		for _, g := range r.infection {
			if g.units > 0 {
				infHasUnits = true
				break
			}
		}
		for _, g := range r.immuneSystem {
			if g.units > 0 {
				immHasUnits = true
				break
			}
		}

		if !infHasUnits {
			sum := 0
			for _, g := range r.immuneSystem {
				if g.units > 0 {
					sum += g.units
				}
			}
			return false, false, sum
		}

		if !immHasUnits {
			sum := 0
			for _, g := range r.infection {
				if g.units > 0 {
					sum += g.units
				}
			}
			return false, true, sum
		}
	}
}

type reindeer struct {
	infection    army
	immuneSystem army
}

func (r reindeer) fight() bool {
	groups := &attackGroups{}
	immuneSystemTargets := army{}
	infectionTargets := army{}
	heap.Init(groups)
	for _, g := range r.immuneSystem {
		if g.units <= 0 {
			continue
		}
		heap.Push(groups, attackGroup{g, false})
		immuneSystemTargets = append(immuneSystemTargets, g)
	}
	for _, g := range r.infection {
		if g.units <= 0 {
			continue
		}
		heap.Push(groups, attackGroup{g, true})
		infectionTargets = append(infectionTargets, g)
	}
	attackList := &attackTargets{}
	heap.Init(attackList)
	for groups.Len() > 0 {
		attacker := heap.Pop(groups).(attackGroup)
		var targetList army
		if attacker.isInfection {
			targetList = immuneSystemTargets
		} else {
			targetList = infectionTargets
		}
		var bestTarget *attackChoice
		for _, target := range targetList {
			thisTarget := attackChoice{
				id:         target.id,
				damage:     target.damageFrom(attacker.group.effectivePower(), attacker.group.damageType),
				ep:         target.effectivePower(),
				initiative: target.initiative,
			}
			if thisTarget.damage == 0 {
				continue
			}
			if bestTarget == nil {
				bestTarget = &thisTarget
				continue
			}
			if thisTarget.damage < bestTarget.damage {
				continue
			}
			if thisTarget.damage > bestTarget.damage {
				bestTarget = &thisTarget
				continue
			}
			if thisTarget.ep < bestTarget.ep {
				continue
			}
			if thisTarget.ep > bestTarget.ep {
				bestTarget = &thisTarget
				continue
			}
			if thisTarget.initiative > bestTarget.initiative {
				bestTarget = &thisTarget
			}
		}

		if bestTarget != nil {
			attack := attackTarget{
				attackerID:  attacker.group.id,
				defenderID:  bestTarget.id,
				isInfection: attacker.isInfection,
				initiative:  attacker.group.initiative,
			}
			heap.Push(attackList, attack)
			if attacker.isInfection {
				newTargets := make(army, 0, len(immuneSystemTargets)-1)
				for _, g := range immuneSystemTargets {
					if g.id == bestTarget.id {
						continue
					}
					newTargets = append(newTargets, g)
				}
				immuneSystemTargets = newTargets
			} else {
				newTargets := make(army, 0, len(infectionTargets)-1)
				for _, g := range infectionTargets {
					if g.id == bestTarget.id {
						continue
					}
					newTargets = append(newTargets, g)
				}
				infectionTargets = newTargets
			}
		}
	}
	stalled := true
	for attackList.Len() > 0 {
		attack := heap.Pop(attackList).(attackTarget)
		var attArmy, defArmy army
		if attack.isInfection {
			attArmy, defArmy = r.infection, r.immuneSystem
		} else {
			attArmy, defArmy = r.immuneSystem, r.infection
		}
		attInfo := attArmy[attack.attackerID-1]
		defInfo := defArmy[attack.defenderID-1]
		attUnits := attInfo.units
		if (attUnits) <= 0 {
			continue
		}
		attDamage := defInfo.damageFrom(attInfo.effectivePower(), attInfo.damageType)
		defUnitsKilled := attDamage / defInfo.hp
		defArmy[attack.defenderID-1].units -= defUnitsKilled
		if defUnitsKilled > 0 {
			stalled = false
		}
	}
	return stalled
}

type attackTarget struct {
	attackerID  int
	defenderID  int
	isInfection bool
	initiative  int
}
type attackTargets []attackTarget

func (h attackTargets) Len() int      { return len(h) }
func (h attackTargets) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h attackTargets) Less(i, j int) bool {
	return h[i].initiative > h[j].initiative
}
func (h *attackTargets) Push(x interface{}) { *h = append(*h, x.(attackTarget)) }
func (h *attackTargets) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type attackGroup struct {
	group       group
	isInfection bool
}

type attackGroups []attackGroup

func (h attackGroups) Len() int      { return len(h) }
func (h attackGroups) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h attackGroups) Less(i, j int) bool {
	epI, epJ := h[i].group.effectivePower(), h[j].group.effectivePower()
	if epI != epJ {
		return epI > epJ
	}
	return h[i].group.initiative > h[j].group.initiative
}
func (h *attackGroups) Push(x interface{}) { *h = append(*h, x.(attackGroup)) }
func (h *attackGroups) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type attackChoice struct {
	id         int
	damage     int
	ep         int
	initiative int
}

type army []group

type group struct {
	id         int
	units      int
	hp         int
	initiative int
	damage     int
	damageType string
	weaknesses map[string]bool
	immunities map[string]bool
}

func (g group) damageFrom(amount int, damageType string) int {
	if g.immunities[damageType] {
		return 0
	}
	if g.weaknesses[damageType] {
		return amount * 2
	}
	return amount
}

func (g group) effectivePower() int {
	return g.units * g.damage
}

func (g group) String() string {
	sb := &strings.Builder{}

	sb.WriteByte('#')
	sb.WriteString(strconv.Itoa(g.id))

	sb.WriteString("\tunits: ")
	sb.WriteString(strconv.Itoa(g.units))
	sb.WriteString("\thp: ")
	sb.WriteString(strconv.Itoa(g.hp))

	sb.WriteString("\tdamage: ")
	sb.WriteString(strconv.Itoa(g.damage))
	sb.WriteString("-")
	sb.WriteString(g.damageType)

	sb.WriteString("\t\tweaknesses: ")
	if len(g.weaknesses) == 0 {
		sb.WriteString("<none>")
	} else {
		weaknesses := make([]string, 0, len(g.weaknesses))
		for weakness := range g.weaknesses {
			weaknesses = append(weaknesses, weakness)
		}
		sort.Strings(weaknesses)
		sb.WriteString(strings.Join(weaknesses, ","))
	}

	sb.WriteString("\t\timmunities: ")
	if len(g.immunities) == 0 {
		sb.WriteString("<none>")
	} else {
		immunities := make([]string, 0, len(g.immunities))
		for immunity := range g.immunities {
			immunities = append(immunities, immunity)
		}
		sort.Strings(immunities)
		sb.WriteString(strings.Join(immunities, ","))
	}

	return sb.String()
}

func load(filename string) reindeer {
	immuneSystem := army{}
	infection := army{}

	inputBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	sections := strings.Split(string(inputBytes), "\n\n")
	for idx, line := range strings.Split(strings.TrimSpace(sections[0]), "\n")[1:] {
		immuneSystem = append(immuneSystem, groupFromLine(idx+1, line))
	}
	for idx, line := range strings.Split(strings.TrimSpace(sections[1]), "\n")[1:] {
		infection = append(infection, groupFromLine(idx+1, line))
	}

	body := reindeer{
		infection:    infection,
		immuneSystem: immuneSystem,
	}
	return body
}

func groupFromLine(id int, line string) group {
	matcher := regexp.MustCompile(`^(\d+) units each with (\d+) hit points(:? \(([a-z,; ]+)\))? with an attack that does (\d+) (\w+) damage at initiative (\d+)$`)
	parts := matcher.FindStringSubmatch(line)
	g := group{
		id:         id,
		units:      utils.MustInt[int](parts[1]),
		hp:         utils.MustInt[int](parts[2]),
		damage:     utils.MustInt[int](parts[5]),
		damageType: parts[6],
		initiative: utils.MustInt[int](parts[7]),
		weaknesses: map[string]bool{},
		immunities: map[string]bool{},
	}
	if parts[4] != "" {
		for _, attr := range strings.Split(parts[4], "; ") {
			if strings.HasPrefix(attr, "weak to ") {
				weaknesses := strings.Split(strings.TrimPrefix(attr, "weak to "), ", ")
				for _, weakness := range weaknesses {
					g.weaknesses[weakness] = true
				}
			} else if strings.HasPrefix(attr, "immune to ") {
				immunities := strings.Split(strings.TrimPrefix(attr, "immune to "), ", ")
				for _, immunity := range immunities {
					g.immunities[immunity] = true
				}
			}
		}
	}
	return g
}

var benchmark = false
