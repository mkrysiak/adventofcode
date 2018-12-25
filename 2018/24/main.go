package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type group struct {
	id         int
	unitType   string
	units      int
	hitPoints  int
	damage     int
	initiative int
	attackType string
	weaknesses map[string]struct{}
	immunities map[string]struct{}
	boost      int
}

type groups []group

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	grps := parseInput(input())

	for !grps.isArmyDefeated() {

		// Target selection - choose one target

		// In decreasing order of effective power, groups choose their targets.
		// In a tie, the group with higher initiative chooses first
		grps.sortByEffectivePower()

		// The attacking group choose to target the group to which it would deal the most damange
		//      after accounting for weaknesses and immunities.  Tiebreaker: enemy with largest effective power, if still tied, highest initiative
		//      If it cannot deal any defending groups damage, it does not choose a target
		attackingGroup := map[int]enemy{}
		chosen := map[int]struct{}{}
		for _, g := range grps {
			if g.units <= 0 {
				continue
			}
			bestEn := g.bestEnemy(grps, chosen)
			if bestEn.damage != -1 {
				attackingGroup[g.id] = bestEn
			}
		}
		// Attacking
		//	Groups attack in decreasing order of initiative, regardless of the army type

		attackers := groups{}
		for i := range attackingGroup {
			for _, vv := range grps {
				if i == vv.id {
					attackers = append(attackers, vv)
				}
			}
		}
		attackers.sortByInitiative()

		for _, attacker := range attackers {
			attkr := grps[grps.indexOf(attacker)]
			if attkr.units <= 0 {
				continue
			}
			enemy := grps[grps.indexOf(attackingGroup[attacker.id].group)]
			newTotalHitPoints := (enemy.units * enemy.hitPoints) - attkr.damageInflicted(enemy)
			if newTotalHitPoints < 0 {
				newTotalHitPoints = 0
			}
			remainingUnitsEnemy := newTotalHitPoints / enemy.hitPoints
			if newTotalHitPoints > 0 && (newTotalHitPoints%enemy.hitPoints > 0) {
				remainingUnitsEnemy++
			}
			for i, v := range grps {
				if v.id == enemy.id {
					grps[i].units = remainingUnitsEnemy
				}
			}
		}
	}

	totalUnits := 0
	for _, v := range grps {
		if v.units > 0 {
			totalUnits += v.units
		}
	}
	return totalUnits
}

func part2() int {
	grps := parseInput(input())
	boost := 0
	for grps.isInfected() {
		grps := parseInput(input())
		for k := range grps {
			if grps[k].unitType == "immune" {
				grps[k].damage += boost
			}
		}

		for !grps.isArmyDefeated() {
			// Target selection - choose one target

			// In decreasing order of effective power, groups choose their targets.
			// In a tie, the group with higher initiative chooses first
			grps.sortByEffectivePower()

			// The attacking group choose to target the group to which it would deal the most damange
			//      after accounting for weaknesses and immunities.  Tiebreaker: enemy with largest effective power, if still tied, highest initiative
			//      If it cannot deal any defending groups damage, it does not choose a target
			attackingGroup := map[int]enemy{}
			chosen := map[int]struct{}{}
			for _, g := range grps {
				if g.units <= 0 {
					continue
				}
				bestEn := g.bestEnemy(grps, chosen)
				if bestEn.damage != -1 {
					attackingGroup[g.id] = bestEn
				}
			}
			// Attacking
			//	Groups attack in decreasing order of initiative, regardless of the army type
			attackers := groups{}
			for i := range attackingGroup {
				for _, vv := range grps {
					if i == vv.id {
						attackers = append(attackers, vv)
					}
				}
			}
			attackers.sortByInitiative()
			if len(attackers) == 0 {
				break
			}
			for _, attacker := range attackers {
				attkr := grps[grps.indexOf(attacker)]
				if attkr.units <= 0 {
					continue
				}
				enemy := grps[grps.indexOf(attackingGroup[attacker.id].group)]
				newTotalHitPoints := (enemy.units * enemy.hitPoints) - attkr.damageInflicted(enemy)
				if newTotalHitPoints < 0 {
					newTotalHitPoints = 0
				}
				remainingUnitsEnemy := newTotalHitPoints / enemy.hitPoints
				if newTotalHitPoints > 0 && (newTotalHitPoints%enemy.hitPoints > 0) {
					remainingUnitsEnemy++
				}
				for i, v := range grps {
					if v.id == enemy.id {
						grps[i].units = remainingUnitsEnemy
					}
				}
			}
		}
		if !grps.isInfected() {
			totalUnits := 0
			for _, v := range grps {
				if v.units > 0 {
					totalUnits += v.units
				}
			}
			return totalUnits
		}
		boost++
	}
	return 0
}

func (g groups) indexOf(g2 group) int {
	for k := range g {
		if g[k].id == g2.id {
			return k
		}
	}
	return -1
}

func (g groups) isArmyDefeated() bool {
	immune, infection := false, false
	for _, v := range g {
		if v.unitType == "immune" && v.units > 0 {
			immune = true
		} else if v.unitType == "infection" && v.units > 0 {
			infection = true
		}
	}
	if (immune && !infection) || (!immune && infection) {
		return true
	}
	return false
}

func (g groups) isInfected() bool {
	for _, v := range g {
		if v.unitType == "infection" && v.units > 0 {
			return true
		}
	}
	return false
}

type enemy struct {
	group  group
	damage int
}
type enemies []enemy

func (g group) bestEnemy(e groups, chosen map[int]struct{}) enemy {
	best := enemy{group{}, -1}
	en := enemies{}
	for _, v := range e {
		if g.unitType == v.unitType || v.units <= 0 {
			continue
		}
		if _, ok := chosen[v.id]; !ok {
			damage := g.damageInflicted(v)
			//  If it cannot deal any defending groups damage, it does not choose a target.
			if damage > 0 {
				en = append(en, enemy{v, damage})
			}
		}
	}
	if len(en) > 0 {
		en.sortByDamage()
		chosen[en[0].group.id] = struct{}{}
		return en[0]
	}
	return best
}

func (g group) damageInflicted(e group) int {
	damage := g.effectivePower()
	if _, ok := e.weaknesses[g.attackType]; ok {
		damage *= 2
	} else if _, ok := e.immunities[g.attackType]; ok {
		damage = 0
	}
	return damage
}

func (e enemies) sortByDamage() {
	sort.Slice(e, func(i, j int) bool {
		if e[i].damage == e[j].damage {
			if e[i].group.effectivePower() == e[j].group.effectivePower() {
				return e[i].group.initiative > e[j].group.initiative
			}
			return e[i].group.effectivePower() > e[j].group.effectivePower()
		}
		return e[i].damage > e[j].damage
	})
}

func (g groups) sortByInitiative() {
	sort.Slice(g, func(i, j int) bool {
		return g[i].initiative > g[j].initiative
	})
}

func (g groups) sortByEffectivePower() {
	sort.Slice(g, func(i, j int) bool {
		if g[i].effectivePower() == g[j].effectivePower() {
			return g[i].initiative > g[j].initiative
		}
		return g[i].effectivePower() > g[j].effectivePower()
	})
}

func (g group) effectivePower() int {
	return g.units * (g.damage + g.boost)
}

func (g groups) print() {
	for _, v := range g {
		if v.units > 0 {
			fmt.Printf("%d: Units: %d HP: %d Damage: %d Initiative: %d AttackType: %s Weaknesses: %v Immunities %v EffectivePower: %d\n",
				v.id, v.units, v.hitPoints, v.damage, v.initiative, v.attackType, v.weaknesses, v.immunities, v.effectivePower())
		}
	}
}

func parseInput(in string) groups {
	grps := groups{}
	lines := strings.Split(in, "\n")
	immune, infection := false, false
	id := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		if line == "Immune System:" {
			immune = true
			continue
		}
		if line == "Infection:" {
			immune = false
			infection = true
			continue
		}
		// (?:immune\sto\s([\w,\s]+))?;?\s?(?:weak\sto\s([\w,\s]+))?
		m := inputRE.FindStringSubmatch(line)

		if m != nil {
			im := immuneRE.FindStringSubmatch(line)
			we := weakRE.FindStringSubmatch(line)
			weaknesses, immunities := map[string]struct{}{}, map[string]struct{}{}
			if im != nil {
				for _, v := range strings.Split(im[1], ", ") {
					immunities[v] = struct{}{}
				}
			}
			if we != nil {
				for _, v := range strings.Split(we[1], ", ") {
					weaknesses[v] = struct{}{}
				}
			}
			units, _ := strconv.Atoi(m[1])
			hitPoints, _ := strconv.Atoi(m[2])
			damage, _ := strconv.Atoi(m[4])
			initiative, _ := strconv.Atoi(m[6])

			id++
			if immune {
				g := group{
					id:         id,
					unitType:   "immune",
					units:      units,
					hitPoints:  hitPoints,
					damage:     damage,
					attackType: m[5],
					initiative: initiative,
					weaknesses: weaknesses,
					immunities: immunities,
					boost:      0,
				}
				grps = append(grps, g)
			}
			if infection {
				g := group{
					id:         id,
					unitType:   "infection",
					units:      units,
					hitPoints:  hitPoints,
					damage:     damage,
					attackType: m[5],
					initiative: initiative,
					weaknesses: weaknesses,
					immunities: immunities,
					boost:      0,
				}
				grps = append(grps, g)
			}
		}
	}

	return grps
}

var inputRE = regexp.MustCompile(`^(\d+)\sunits\seach\swith\s(\d+)\shit\spoints\s(?:\((.*)\)\s|)with\san\sattack\sthat\sdoes\s(\d+)\s(\w+)\sdamage\sat\sinitiative\s(\d+)$`)
var immuneRE = regexp.MustCompile(`immune\sto\s([\w\s,]+)+`)
var weakRE = regexp.MustCompile(`weak\sto\s([\w\s,]+)+`)

func testInput() string {
	return `Immune System:
17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2
989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3

Infection:
801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1
4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4`
}

func input() string {
	return `Immune System:
3400 units each with 1430 hit points (immune to fire, radiation, slashing) with an attack that does 4 radiation damage at initiative 4
138 units each with 8650 hit points (weak to bludgeoning; immune to slashing, cold, radiation) with an attack that does 576 slashing damage at initiative 16
255 units each with 9469 hit points (weak to radiation, fire) with an attack that does 351 bludgeoning damage at initiative 8
4145 units each with 2591 hit points (immune to cold; weak to slashing) with an attack that does 6 fire damage at initiative 12
3605 units each with 10989 hit points with an attack that does 26 fire damage at initiative 19
865 units each with 11201 hit points with an attack that does 102 slashing damage at initiative 10
633 units each with 10092 hit points (weak to slashing, radiation) with an attack that does 150 slashing damage at initiative 11
2347 units each with 3322 hit points with an attack that does 12 cold damage at initiative 2
7045 units each with 3877 hit points (weak to radiation) with an attack that does 5 bludgeoning damage at initiative 5
1086 units each with 8626 hit points (weak to radiation) with an attack that does 69 slashing damage at initiative 13

Infection:
2152 units each with 12657 hit points (weak to fire, cold) with an attack that does 11 fire damage at initiative 18
40 units each with 39458 hit points (immune to radiation, fire, slashing; weak to bludgeoning) with an attack that does 1519 slashing damage at initiative 7
59 units each with 35138 hit points (immune to radiation; weak to fire) with an attack that does 1105 fire damage at initiative 15
1569 units each with 51364 hit points (weak to radiation) with an attack that does 55 radiation damage at initiative 17
929 units each with 23887 hit points (weak to bludgeoning) with an attack that does 48 cold damage at initiative 14
5264 units each with 14842 hit points (immune to cold, fire; weak to slashing, bludgeoning) with an attack that does 4 bludgeoning damage at initiative 9
1570 units each with 30419 hit points (weak to radiation, cold; immune to fire) with an attack that does 35 slashing damage at initiative 1
1428 units each with 21393 hit points (weak to radiation) with an attack that does 29 cold damage at initiative 6
1014 units each with 25717 hit points (weak to fire) with an attack that does 47 fire damage at initiative 3
7933 units each with 29900 hit points (immune to bludgeoning, radiation, slashing) with an attack that does 5 slashing damage at initiative 20`
}
