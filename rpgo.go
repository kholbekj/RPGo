// Tiny RPG to try out structs - Licensed with the cruelest
// of patents known to mankind. If you copy this original and
// beautiful program, a starving, rabies-infected pdigeon
// will be released upon your local community.
package main

import (
	"fmt"
)

func main() {
	fmt.Println("You're fighting a monster!")

	goblin1 := createMonster("Goblin", monster_easy)
	player := createCharacter()

	fmt.Printf("It's a %v!\n", goblin1.Name)
	if fight(&player, goblin1) {
		fmt.Println("You won!")
		fmt.Printf("Winning nested you %v XP! You lucky bugger!\n", goblin1.Xp)
		saveState(player)
	} else {
		fmt.Println("You lost!")
	}
}

func fight(vip *Player, enemy Monster) bool {
	for vip.alive() {
		enemy.Hp--
		if !enemy.alive() {
			return true
		}
	}
	return false
}

type Entity struct {
	Name                  string
	Str, Vit, Def, Hp, Xp int
}

func (e Entity) alive() bool {
	return e.Hp > 0
}

type Player struct {
	Entity
	Inventory []int
}

type Monster struct {
	Entity
	lootTable []int
}

// Monster types					{str, vit, def, hp, xp}
var (
	monster_weak      = [...]int{1, 1, 1, 3, 1}
	monster_easy      = [...]int{2, 1, 1, 4, 2}
	monster_medium    = [...]int{3, 2, 2, 5, 5}
	monster_difficult = [...]int{5, 3, 3, 8, 10}
	monster_hard      = [...]int{8, 5, 5, 12, 20}
)

func createMonster(name string, mtype [len(monster_weak)]int) Monster {
	return Monster{Entity{name, mtype[0], mtype[1], mtype[2], mtype[3], mtype[4]}, []int{0}}
}
