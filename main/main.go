package main

import (
	"algorithms/recursion"
	"fmt"
)

func main() {
	// Recursion Sum
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sumResult := recursion.RecursionSum(nums)

	// Recursion count
	items := []string{"a", "t", "a", "k", "a", "n"}
	countResult := recursion.RecursionCount(items, 0)

	fmt.Printf("Sum result: %v\n", sumResult)
	fmt.Printf("Count result: %v\n", countResult)
}
