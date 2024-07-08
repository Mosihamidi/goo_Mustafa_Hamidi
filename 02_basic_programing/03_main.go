package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%4 == 0 && i%7 == 0 {
			fmt.Println("buzz")
		} else if i%4 == 0 {
			fmt.Println(i * i)
		} else if i%7 == 0 {
			fmt.Println(int(math.Pow(float64(i), 3)))
		} else {
			fmt.Println(i)
		}
	}
}
