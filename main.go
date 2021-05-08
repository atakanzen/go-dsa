package main

import (
	"algorithms/recursion"
	"fmt"
)

func main() {
	nums := []int{3, 5, 4, 1, 9, 105, 7, 6, 8, 10, 2}

	// Recursion Sum
	sumResult := recursion.RecursionSum(nums)

	// Recursion Count
	items := []string{"a", "t", "a", "k", "a", "n"}
	countResult := recursion.RecursionCount(items, 0)

	// Recursion Max
	maxResult := recursion.RecursionMax(nums)

	// Recursion Quicksort
	quicksortResult := recursion.Quicksort(nums)

	fmt.Printf("Sum result: %v\n", sumResult)
	fmt.Printf("Count result: %v\n", countResult)
	fmt.Printf("Max result: %v\n", maxResult)
	fmt.Printf("Quicksort result: %v\n", quicksortResult)
}
