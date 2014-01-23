package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// createCharacter should be run at beginning of
// the program to get the info necessary
// to create a new player.
func createCharacter() Player {
	welcome := fmt.Sprintf("Welcome to RPGo! Since this is your first time\n" +
		"visiting, we'll need to make you a brand new\n" +
		"character! We'll start off with your name!\n" +
		"Enter name: ")

	na := dialogue(welcome, "Too long, please keep it 8 characters or shorter.", isOfLength(8))

	p := Player{Entity{na, 5, 5, 5, 5, 0}, []int{0}}

	menu := fmt.Sprintf("%[1]v, eh? Well, it's your character... \n"+
		"Anyway, how would you describe %[1]v?\n"+
		"[1] %[1]v lifts boulders for breakfast.\n"+
		"[2] %[1]v can dodge bullets.\n"+
		"[3] %[1]v can take a beating from Chuck Norris\n", na)

	selection := dialogue(menu, "Please type 1, 2 or 3...", isInRange(1, 3))

	switch selection {
	case "1":
		p.Str = 10
	case "2":
		p.Def = 10
	case "3":
		p.Vit = 10
		p.Hp = 10
	}

	fmt.Println("Great, this is what you got:")
	characterSheet(p)
	return p
}

// characterSheet takes a player, and prints said players character sheet.
func characterSheet(p Player) {
	fmt.Printf("+--------------------->\n")
	fmt.Printf("| name	|	%v\n", p.Name)
	fmt.Printf("|	|\n")
	fmt.Printf("| Str	|	%v\n", p.Str)
	fmt.Printf("| Def	|	%v\n", p.Def)
	fmt.Printf("| Vit	|	%v\n", p.Vit)
	fmt.Printf("|	|\n")
	fmt.Printf("| HP	|	%v/%v\n", p.Hp, p.Vit)
	fmt.Printf("+--------------------->\n")
}

// dialogue prompts the user with the prompts, reads a line from stdIn, and if compare
// with that string is true, it returns the string, otherwise it prints the failure message,
// and starts over.
func dialogue(prompt, failure string, compare func(s string) bool) string {
	for {
		fmt.Printf(prompt)
		in := bufio.NewReader(os.Stdin)
		input, _ := in.ReadString('\n')
		input = strings.TrimSpace(input)
		if compare(input) {
			return input
		}

	}
	return ""
}

// isInRange returns a function for checking weather a string is an integer in the range [min-max]
func isInRange(min, max int) func(string) bool {
	return func(s string) bool {
		i, _ := strconv.Atoi(s)
		return (min <= i && i <= max)
	}
}

// isOfLength returns a function for checking of a string is of equal or shorter lenght than the argument.
func isOfLength(l int) func(string) bool {
	return func(s string) bool {
		return len(s) <= l
	}

}
