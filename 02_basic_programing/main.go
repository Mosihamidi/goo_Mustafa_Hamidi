package main

import (
	"fmt"
)

var pi = 3.14

func main() {
	var radius, height float64

	fmt.Println("radius of the tube")
	fmt.Scanf("%f", &radius)

	fmt.Println("height of the tube")
	fmt.Scanf("%f", &height)

	volume := pi * radius * radius * height

	fmt.Printf("The volume of the tube is: %.2f\n", volume)
}
