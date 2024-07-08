package main

import "fmt"

func main() {
	num := 26
	result := ""
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			result += "I"
		} else {
			result += "O"
		}
	}
	fmt.Println(result)
}
