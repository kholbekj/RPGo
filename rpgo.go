// Tiny RPG to try out structs - Licensed with the cruelest
// of patents known to mankind. If you copy this original and
// beautiful program, a starving, rabies-infected pdigeon
// will be released upon your local community.
package main

import (
	"bufio"
	"fmt"
	//"time"
	//"math/rand"
	"os"
	"strings"
)

func main() {

	fmt.Println("You're fighting a monster!")
	goblin1 := Monster{"GOBLIN", 3, 3, 3, 3, 1}
	fmt.Printf("It's a %v!\n", goblin1.name)
	if fight(&vip, goblin1) {
		fmt.Println("You won!")
		fmt.Printf("Winning nested you %v XP! You lucky bugger!\n", goblin1.xp)
	} else {
		fmt.Println("You lost!")
	}
}

type Player struct {
	name              string
	str, vit, def, hp int
}

func (p Player) alive() bool {
	return p.hp > 0
}

type Monster struct {
	name                  string
	str, vit, def, hp, xp int
}

func (m Monster) alive() bool {
	return m.hp > 0
}

func fight(vip *Player, enemy Monster) bool {
	for vip.alive() {
		enemy.hp--
		if !enemy.alive() {
			return true
		}
	}
	return false
}
