package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type animal struct {
	food, locomotion, sound string
}

func (a animal) Eat() {
	fmt.Println(a.food)
}

func (a animal) Move() {
	fmt.Println(a.locomotion)
}

func (a animal) Speak() {
	fmt.Println(a.sound)
}

func main() {
	animals := make(map[string]animal)

	animals["cow"] = animal{"grass", "walk", "moo"}
	animals["bird"] = animal{"worms", "fly", "peep"}
	animals["snake"] = animal{"mice", "slither", "hsss"}

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
