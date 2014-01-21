package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func saveState(p Player) bool {
	save, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if saveToFile(save) {
		return true
	} else {
		return false
	}
}

func saveToFile(jsonString []byte) bool {
	err := ioutil.WriteFile(".savedata.dat", jsonString, 0644)
	if err != nil {
		panic(err)
		return false
	}

	return true
}

func loadState() {
	// nyi
}
