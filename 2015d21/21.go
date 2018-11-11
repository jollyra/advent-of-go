package main

import (
	"fmt"
)

type weapon struct {
	Name   string
	Cost   int
	Damage int
}

type armor struct {
	Name  string
	Cost  int
	Armor int
}

type ring struct {
	Name   string
	Cost   int
	Damage int
	Armor  int
}

type entity interface {
	Damage() int
	Armor() int
	Hp() int
	SetHp(int)
}

type hero struct {
	hp     int
	Weapon weapon
	armor  armor
	Ring1  ring
	Ring2  ring
}

func (h *hero) Damage() int  { return h.Weapon.Damage + h.Ring1.Damage + h.Ring2.Damage }
func (h *hero) Armor() int   { return h.armor.Armor + h.Ring1.Armor + h.Ring2.Armor }
func (h *hero) Hp() int      { return h.hp }
func (h *hero) SetHp(hp int) { h.hp = hp }

func (h *hero) Attack(e entity) bool {
	heroDamage := h.Damage()
	entityArmor := e.Armor()
	if damage := heroDamage - entityArmor; damage > 0 {
		e.SetHp(e.Hp() - damage)
	} else {
		e.SetHp(e.Hp() - 1)
	}
	return e.Hp() <= 0
}

var dagger = weapon{"Dagger", 8, 4}
var shortsword = weapon{"Shortsword", 10, 5}
var warhammer = weapon{"Warhammer", 25, 6}
var longsword = weapon{"Longsword", 40, 7}
var greataxe = weapon{"Greataxe", 74, 8}

var weapons = []weapon{dagger, shortsword, warhammer, longsword, greataxe}

var leather = armor{"Leather", 13, 1}
var chainmail = armor{"Chainmail", 31, 2}
var splintmail = armor{"Splintmail", 53, 3}
var bandedmail = armor{"Bandedmail", 75, 4}
var platemail = armor{"Platemail", 102, 5}
var naked = armor{"naked", 0, 0}

var armors = []armor{leather, chainmail, splintmail, bandedmail, platemail, naked}

var w1 = ring{"Damage + 1", 25, 1, 0}
var w2 = ring{"Damage + 2", 50, 2, 0}
var w3 = ring{"Damage + 3", 100, 3, 0}
var a1 = ring{"Defense + 1", 20, 0, 1}
var a2 = ring{"Defense + 2", 40, 0, 2}
var a3 = ring{"Defense + 3", 80, 0, 3}
var n0 = ring{"none 0", 0, 0, 0}
var n1 = ring{"none 1", 0, 0, 0}

var rings = []ring{w1, w2, w3, a1, a2, a3, n0, n1}

// Returns true if the hero won and false otherwise
func runScenario(hero *hero, boss *hero) bool {
	for {
		if kill := hero.Attack(boss); kill {
			return true
		}
		if kill := boss.Attack(hero); kill {
			return false
		}
	}
}

func main() {
	minCost := 1<<32 - 1
	maxCost := 0
	for _, w := range weapons {
		for _, a := range armors {
			for _, r1 := range rings {
				for _, r2 := range rings {
					h := &hero{
						hp:     100,
						Weapon: w,
						armor:  a,
						Ring1:  r1,
						Ring2:  r2,
					}

					boss := &hero{
						hp:     100,
						Weapon: weapon{"boss weapon", 0, 8},
						armor:  armor{"boss armor", 0, 2},
						Ring1:  ring{"boss ring", 0, 0, 0},
						Ring2:  ring{"boss ring", 0, 0, 0},
					}

					cost := w.Cost + a.Cost + r1.Cost + r2.Cost
					if win := runScenario(h, boss); win {
						if cost < minCost {
							minCost = cost
						}
					} else {
						if cost > maxCost {
							maxCost = cost
						}
					}
				}
			}
		}
	}

	fmt.Println("The minimum cost to beat the boss is", minCost)
	fmt.Println("The max gold spent to lose is", maxCost)
}
