package main

import "fmt"

func main() {
	var e float32

	fmt.Print("Enter a float number: ")

	if _, err := fmt.Scanf("%f", &e); err != nil {
		fmt.Println("Error: ", err)
		return 
	} 
	
	fmt.Println("A float number converted to int is", int(e))
}
