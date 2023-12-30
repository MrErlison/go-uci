package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return ((0.5)*a*math.Pow(t, 2) + v0*t + s0)
	}
}

func main() {
	var a, v0, s0, t float64

	fmt.Print("Enter the acceleration: ")
	fmt.Scanln(&a)

	fmt.Print("Enter the velocity: ")
	fmt.Scanln(&v0)

	fmt.Print("Enter the displacement: ")
	fmt.Scanln(&s0)

	fmt.Print("Enter the elapsed time: ")
	fmt.Scanln(&t)

	f := GenDisplaceFn(a, v0, s0)
	fmt.Printf("Displacement at t=%f : %f\n", t, f(t))
	fmt.Printf("Displacement at t=%f : %f\n", 3.0, f(3.0))
	fmt.Printf("Displacement at t=%f : %f\n", 5.0, f(5.0))
}
