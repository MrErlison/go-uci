package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func convert(slice []string) ([]int, error) {
	var s []int = []int{}
	for _, v := range slice {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("only integers are accepted.")
		}
		s = append(s, i)
	}
	return s, nil
}

func main() {
	const PART = 4

	fmt.Println("Enter a sequence of integers multiple by 4. Whitespace as delimiter.")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	inputSlice := strings.Fields(input)
	if len(inputSlice)%4 != 0 {
		log.Fatal("the sequence of integers must be multiples by 4.")
	}

	slice, err := convert(inputSlice)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(PART)

	sortedSlice := []int{}
	totalElements := len(slice) / PART
	for i := 0; i < len(slice); i += totalElements {
		s := slice[i : i+totalElements]
		go func(s []int) {
			fmt.Println("unsorted part", s)
			sort.Ints(s)
			wg.Done()
		}(s)
		sortedSlice = append(sortedSlice, s...)
	}
	wg.Wait()

	sort.Ints(sortedSlice)
	fmt.Println("sorted result", sortedSlice)
}
