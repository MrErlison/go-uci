package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food, locomotion, sound string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.sound)
}

type Bird struct {
	food, locomotion, sound string
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.sound)
}

type Snake struct {
	food, locomotion, sound string
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.sound)
}

func main() {
	animals := make(map[string]Animal)
	animals["cow"] = Cow{"grass", "walk", "moo"}
	animals["snake"] = Snake{"mice", "slither", "hsss"}
	animals["bird"] = Bird{"worms", "fly", "peep"}

	for {
		var command, name, action string

		fmt.Print("Enter de command\n> ")
		fmt.Scan(&command, &name, &action)

		switch command {
		case "newanimal":
			if _, ok := animals[action]; !ok {
				fmt.Println("Invalid animal type!")
				continue
			}
			animals[name] = animals[action]
			fmt.Println("Created it!")

		case "query":
			if _, ok := animals[name]; !ok {
				fmt.Println("Invalid animal!")
				continue
			}
			switch action {
			case "eat":
				animals[name].Eat()
			case "move":
				animals[name].Move()
			case "speak":
				animals[name].Speak()
			default:
				fmt.Println("Invalid action! Choose eat, move or speak")
			}
		default:
			fmt.Println("Invalid command! Choose newanimal or query")
		}
	}
}

/*
Define an interface type called Animal which describes the methods of an animal.

Specifically, the Animal interface should contain the methods Eat(), Move(),
and Speak(), which take no arguments and return no values.

The Eat() method should print the animal’s food, the Move() method
should print the animal’s locomotion, and the Speak() method
should print the animal’s spoken sound. Define three types Cow, Bird, and Snake.

For each of these three types, define methods Eat(), Move(), and Speak() so that
the types Cow, Bird, and Snake all satisfy the Animal interface.

When the user creates an animal, create an object of the appropriate type.

Your program should call the appropriate method when the user issues a query command.
*/
