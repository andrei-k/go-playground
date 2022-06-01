package main

import (
	"fmt"
)

// Define an interface type which describes the methods of an animal
type animal interface {
	eat()
	move()
	speak()
	getName() string
}

// Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.
type cow struct {
	name string
}

func (a cow) eat() {
	fmt.Println("grass")
}
func (a cow) move() {
	fmt.Println("walk")
}
func (a cow) speak() {
	fmt.Println("moo")
}
func (a cow) getName() string {
	return a.name
}

type bird struct {
	name string
}

func (a bird) eat() {
	fmt.Println("worms")
}
func (a bird) move() {
	fmt.Println("fly")
}
func (a bird) speak() {
	fmt.Println("peep")
}
func (a bird) getName() string {
	return a.name
}

type snake struct {
	name string
}

func (a snake) eat() {
	fmt.Println("mice")
}
func (a snake) move() {
	fmt.Println("slither")
}
func (a snake) speak() {
	fmt.Println("hsss")
}
func (a snake) getName() string {
	return a.name
}

func main() {
	var animals []animal

	for {
		var command, name, input string
		fmt.Print("> ")
		fmt.Scan(&command, &name, &input)

		// Create an object of the appropriate type
		if command == "new" {
			switch input {
			case "cow":
				cow := cow{name: name}
				animals = append(animals, cow)
				fmt.Println(cow.name)
			case "bird":
				bird := bird{name: name}
				animals = append(animals, bird)
			case "snake":
				snake := snake{name: name}
				animals = append(animals, snake)
			default:
				fmt.Println("Please enter 'cow', 'bird', or 'snake'")
				continue
			}
			fmt.Println("Created it!")

		} else if command == "query" {
			found := false
			for _, a := range animals {
				if a.getName() == name {
					switch input {
					case "eat":
						a.eat()
					case "move":
						a.move()
					case "speak":
						a.speak()
					default:
						fmt.Println("Please enter 'eat', 'move', or 'speak'")
					}
					found = true
				}
			}
			if !found {
				fmt.Println("No animal found by this name!")
			}

		} else {
			fmt.Println("Please start the command with 'newanimal' or 'query'")
		}
	}
}
