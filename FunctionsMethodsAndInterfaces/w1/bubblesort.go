package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Swap(slice []int, index int) {
	slice[index], slice[index+1] = slice[index+1], slice[index]
}

func BubbleSort(slice []int) {
	s := len(slice)
	for i := 0; i < s; i++ {
		for j := 0; j < s-1-i; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

func main() {
	fmt.Println("Enter an integer sequence of up to 10 integers separated by whitespace: ")
	userInput := bufio.NewScanner(os.Stdin)
	userInput.Scan()
	inputString := userInput.Text()

	var integers []int
	for _, v := range strings.Fields(inputString) {
		integer, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("Error:", err)
		}
		integers = append(integers, integer)
	}

	BubbleSort(integers)

	fmt.Printf("Sorted: %v\n", integers)
}
