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
	player := loadState()

	fmt.Printf("It's a %v!\n", goblin1.Name)
	if fight(&player, goblin1) {
		fmt.Println("You won!")
		fmt.Printf("Winning nested you %v XP! You lucky bugger!\n", goblin1.Xp)
		player.addXp(goblin1.Xp)
		player.addItem(3)
		saveState(player)
	} else {
		fmt.Println("You lost!")
	}
	//Test gameboard
	gb := gameBoard{}
	gb.rooms = append(gb.rooms, room{0, "Haunted Marketplace", "This place is emptier than your wallet. And that's saying something!", []Monster{goblin1}, map[string]int{"Herp": 23}})
	gb.currentRoom = 0
	action := dialogue(gb.rooms[gb.currentRoom].getMenu(), "Please enter one of the numbers in the menu..", isInRange(1, len(gb.rooms[gb.currentRoom].enemies)+len(gb.rooms[gb.currentRoom].paths)))
	fmt.Println("You picked number ", action, "!")

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
