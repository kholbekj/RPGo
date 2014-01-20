package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CreateCharacter should be run at beginning of
// the program to get the info necessary
// to create a new player.
func createCharacter() Player {
	fmt.Printf("Welcome to PRGo! Since this is your first time\n" +
		"visiting, we'll need to make you a brand new\n" +
		"character! We'll start off with your name!\n")
	fmt.Print("Enter name: ")
	in := bufio.NewReader(os.Stdin)
	na, _ := in.ReadString('\n')
	na = strings.TrimSpace(na)

	fmt.Printf("%v, eh? Well, it's your character... \n", na)
	fmt.Printf("Anyway, how would you describe %v?\n", na)

	p := Player{}

	for chosen := false; chosen == false; {
		fmt.Printf("[1] %v lifts boulders for breakfast.\n", na)
		fmt.Printf("[2] %v can dodge bullets.\n", na)
		fmt.Printf("[3] %v can take a beating from Chuck Norris\n", na)
		selection, _ := in.ReadString('\n')
		selection = strings.TrimSpace(selection)
		switch selection {
		case "1":
			p = Player{Entity{na, 10, 5, 5, 5, 0}, []int{0}}
			chosen = true
		case "2":
			p = Player{Entity{na, 5, 5, 10, 5, 0}, []int{0}}
			chosen = true
		case "3":
			p = Player{Entity{na, 5, 10, 5, 10, 0}, []int{0}}
			chosen = true
		default:
			fmt.Println("Please chose 1, 2 or 3..")
		}
	}
	fmt.Println("Great, this is what you got:")
	characterSheet(p)
	return p
}

func characterSheet(p Player) {
	fmt.Printf("+-------------------------------+\n")
	fmt.Printf("| name |	%v\n", p.name)
	fmt.Printf("|\n")
	fmt.Printf("| Str	|   %v\n", p.str)
	fmt.Printf("| Def	|\n")
	fmt.Printf("| Vit	|   %v\n", p.vit)
	fmt.Printf("|\n")
	fmt.Printf("| HP	|	%v/%v\n", p.hp, p.vit)
	fmt.Printf("+-------------------------------+\n")
}
