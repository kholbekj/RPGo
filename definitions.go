package main

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

type item struct {
	name string
}
type asset struct {
	id                int
	name, description string
	//trigger func()
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

type room struct {
	id int
	name, description string
	interactives      []asset
	paths             map[string]int
}

//func (args ...string) func (g gameBoard)(err)
