package main

import "fmt"

func main() {
	fmt.Println("Race condition is when multiple threads")
	fmt.Println("are trying to manipulate the same variable")
	fmt.Println("and the final result is nondeterministic.")

	for x := 0; x < 10; x++ {
		rc := 0
		for y := 0; y < 100; y++ {
			go func(rc *int) {
				*rc++
			}(&rc)

			go func(rc *int) {
				*rc--
			}(&rc)
		}
		fmt.Printf("rc = %d\n", rc)
	}

	fmt.Println("Due to race condition the results is unpredictable.")
}
