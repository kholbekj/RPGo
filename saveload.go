package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func saveState(p Player) {
	save, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	saveToFile(save)
}

func saveToFile(jsonString []byte) {
	err := ioutil.WriteFile(".savedata.dat", jsonString, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Game Saved Succesfully!")
}

func loadState() Player {
	file := loadFromFile()
	if file == nil {
		return createCharacter()
	}

	var p Player
	err := json.Unmarshal(file, &p)
	if err != nil {
		panic(err)
	}

	return p
}

func loadFromFile() []byte {
	file, err := ioutil.ReadFile(".savedata.dat")
	if err != nil {
		// file not found, or other error,
		// load game without a savefile
		return nil
	}

	return file
}
