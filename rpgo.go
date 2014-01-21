// Tiny RPG to try out structs - Licensed with the cruelest
// of patents known to mankind. If you copy this original and
// beautiful program, a starving, rabies-infected pdigeon
// will be released upon your local community.
package main

import (
	"fmt"
)

func init() {

}

func main() {
	fmt.Println("You're fighting a monster!")

	goblin1 := createMonster("Goblin", monster_easy)
	player := createCharacter()

	fmt.Printf("It's a %v!\n", goblin1.Name)
	if fight(&player, goblin1) {
		fmt.Println("You won!")
		fmt.Printf("Winning nested you %v XP! You lucky bugger!\n", goblin1.Xp)
		player.addXp(goblin1.Xp)
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
