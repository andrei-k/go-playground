package main

import (
	"fmt"
)

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	var animals = make(map[string]Animal)
	animals["cow"] = Animal{"grass", "walk", "moo"}
	animals["bird"] = Animal{"worms", "fly", "peep"}
	animals["snake"] = Animal{"mice", "slither", "hsss"}

	var inputAnimal, inputAction string

	for {
		fmt.Print(">")
		// Example user request: "cow eat", "bird move", "snake speak"
		fmt.Scan(&inputAnimal, &inputAction)

		switch inputAction {
		case "eat":
			animals[inputAnimal].Eat()
		case "move":
			animals[inputAnimal].Move()
		case "speak":
			animals[inputAnimal].Speak()
		default:
			fmt.Println("Invalid animal or action")
		}
	}
}
