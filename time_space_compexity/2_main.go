package main

import (
	"fmt"
)

// Function to compute x raised to the power n using Exponentiation by Squaring
func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0 // Handle negative exponents if needed
	}
	if n%2 == 0 {
		halfPow := pow(x, n/2)
		return halfPow * halfPow
	}
	return x * pow(x, n-1)
}

func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
}
