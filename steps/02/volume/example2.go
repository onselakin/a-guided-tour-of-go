package main

import "fmt"

const pi float64 = 3.14159265358979323846264338327950288419716939937510582097494459

func main() {
	fmt.Println("Pi is ", pi)
	fmt.Println("Two pies are better than one ", multiply(pi))
}

func multiply(input float64) float64 {
	return input * 2
}
