package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter an address: ")
	scanner.Scan()
	address := scanner.Text()

	person := make(map[string]string)
	person["name"] = name
	person["address"] = address

	json, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	os.Stdout.Write(json)
}
