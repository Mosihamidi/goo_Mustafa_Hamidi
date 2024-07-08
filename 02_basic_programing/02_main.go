package main

import "fmt"

func main() {
	var score int

	fmt.Println("Enter your score: ")
	fmt.Scan(&score)

	if score < 0 || score > 100 {
		fmt.Println("Invalid score")
	} else if score >= 85 {
		fmt.Println("Grade: A")
	} else if score >= 70 {
		fmt.Println("Grade: B")
	} else if score >= 55 {
		fmt.Println("Grade: C")
	} else if score >= 40 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: E")
	}
}
