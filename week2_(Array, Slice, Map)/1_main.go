package main

import "fmt"

func merge(data [][]int) []int {
	seen := make(map[int]bool) // Map to track unique elements
	result := []int{}

	// Flatten the 2D array and add unique elements to result
	for _, subArray := range data {
		for _, value := range subArray {
			if !seen[value] {
				seen[value] = true
				result = append(result, value)
			}
		}
	}

	return result
}

func main() {
	res1 := merge([][]int{
		{1, 2, 5, 4, 3},
		{1, 2, 7, 8},
	})
	fmt.Println(res1) // [1, 2, 5, 4, 3, 7, 8]

	res2 := merge([][]int{
		{1, 2, 5, 4, 5},
		{7, 9, 10, 5},
	})
	fmt.Println(res2) // [1, 2, 5, 4, 7, 9, 10]

	res3 := merge([][]int{
		{1, 4, 5},
		{7},
	})
	fmt.Println(res3) // [1, 4, 5, 7]
}
