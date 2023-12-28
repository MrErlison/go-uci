package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	fmt.Print("Enter the filename: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filename := scanner.Text()

	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	names := make([]Name, 0)
	person := Name{}
	for _, l := range strings.Split(string(content), "\n") {
		lineSplited := strings.Split(l, " ")
		if len(lineSplited) < 2 {
			continue
		}
		person.fname = lineSplited[0]
		person.lname = lineSplited[1]
		names = append(names, person)
	}

	for _, n := range names {
		fmt.Printf("\nfirst name: %s, last name: %s", n.fname, n.lname)
	}

	fmt.Printf("\n")
}
