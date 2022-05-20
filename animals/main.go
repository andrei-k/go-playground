package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Animal struct {
	Food, Locomotion, Sound string
}

func (a Animal) Eat() {
	fmt.Println(a.Food)
}

func (a Animal) Move() {
	fmt.Println(a.Locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.Sound)
}

type MyAnimal struct {
	Name, Animal string
}

func main() {
	var animals = make(map[string]Animal)
	animals["cow"] = Animal{"grass", "walk", "moo"}
	animals["bird"] = Animal{"worms", "fly", "peep"}
	animals["snake"] = Animal{"mice", "slither", "hsss"}

	fmt.Println(animals)

	var myAnimals []MyAnimal

	for {
		var command, name, animal string
		fmt.Print("> ")
		fmt.Scan(&command, &name, &animal)
		fmt.Println("You typed:", command, name, animal)

		// Each “newanimal” command must be a single line containing three strings.
		// The first string is “newanimal”
		// The second string is an arbitrary string which will be the name of the new animal
		// The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
		// Your program should process each newanimal command by creating the new animal.

		if command == "new" {
			myAnimals = append(myAnimals, MyAnimal{name, animal})
			fmt.Println("Created it!")
		}

		// Each “query” command must be a single line containing 3 strings
		// The first string is “query”
		// The second string is the name of the animal
		// The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”
		// Your program should process each query command by printing out the requested data
		if command == "query" {
			fmt.Println(myAnimals)

			idx := slices.IndexFunc(myAnimals, func(m MyAnimal) bool { return m.Name == name })
			if idx > -1 {
				a := myAnimals[idx].Animal
				fmt.Println(name, "is a", a)
				animals[a].Speak()
			}
		}
	}

}
