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

	fmt.Println("Welcome to tinyRPG - the tiny RPG for big people or whatever.")
	fmt.Println("First things first - what would you like to be called?")

	in := bufio.NewReader(os.Stdin)
	name, _ := in.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Alright, %v! Please select one:\n", name)
KIND_SELECT:
	fmt.Println("1: I want to be strong!")
	fmt.Println("2: I want to be tough as a rock!")
	fmt.Println("3: I want to be dodgy!")

	kind, _ := in.ReadString('\n')
	kind = strings.TrimSpace(kind)
	var vip Player

	switch kind {
	case "1":
		vip = Player{name, 10, 5, 5, 5}
		break
	case "2":
		vip = Player{name, 5, 10, 5, 10}
		break
	case "3":
		vip = Player{name, 5, 5, 10, 5}
		break
	default:
		fmt.Println("Just type 1, 2 or 3, please..")
		goto KIND_SELECT
	}
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
