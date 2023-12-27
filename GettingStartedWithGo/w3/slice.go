package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input string
	var s []int = make([]int, 0, 3)

	for {
		sort.Ints(s)
		fmt.Printf("slice: %v, len:%d, cap:%d\n", s, len(s), cap(s))

		fmt.Print("Enter an integer or X to quit: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()

		if strings.ToUpper(input)[0] == 'X' {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid")
			continue
		}

		s = append(s, num)
	}

	fmt.Printf("Bye")
}
