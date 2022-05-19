package main

import "fmt"

// Interface type
type Animal interface {
	// List all the functions this interface must have in order to be that interface.
	// For example, Dog must have these two functions to satisfy the requirements of the Animal interface.
	Says() string
	HowManyLegs() int
}

type Dog struct {
	Name      string
	Sound     string
	NumOfLegs int
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumOfLegs
}

type Cat struct {
	Name      string
	Sound     string
	NumOfLegs int
	HasTail   bool
}

func (c *Cat) Says() string {
	return c.Sound
}

func (c *Cat) HowManyLegs() int {
	return c.NumOfLegs
}

func main() {
	dog := Dog{
		Name:      "dog",
		Sound:     "woof",
		NumOfLegs: 4,
	}

	Riddle(&dog)

	var cat Cat
	cat.Name = "cat"
	cat.NumOfLegs = 4
	cat.Sound = "meow"
	cat.HasTail = true

	Riddle(&cat)
}

func Riddle(a Animal) {
	riddle := fmt.Sprintf(`This animal says %s and has %d legs. What animal is it?`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}
