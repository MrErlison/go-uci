package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Enter a string: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)

	if strings.HasPrefix(text, "i") && strings.Contains(text, "a") && strings.HasSuffix(text, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
