package main

import (
	"fmt"
	"strings"
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
	animals := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	for {
		var name, info string

		fmt.Println("Enter the animal [cow, bird or snake]> ")
		fmt.Scanln(&name)
		animal, ok := animals[name]
		if !ok {
			fmt.Printf("The name %s is not in (cow, bird, snake)\n", name)
			continue
		}

		fmt.Println("Enter the information [eat, move or speak]> ")
		fmt.Scanln(&info)
		switch strings.ToLower(info) {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Printf("The info %s is not in (eat, move, speak)\n", info)
		}
	}
}
