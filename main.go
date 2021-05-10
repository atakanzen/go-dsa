package main

import (
	"algorithms/recursion"
	"algorithms/search"
	"fmt"
)

func main() {
	nums := []int{3, 5, 4, 1, 9, 105, 7, 6, 8, 10, 2}

	// Recursion Sum
	sumResult := recursion.RecursionSum(nums)
	fmt.Printf("Sum Result: %v\n", sumResult)

	// Recursion Count
	items := []string{"a", "t", "a", "k", "a", "n"}
	countResult := recursion.RecursionCount(items, 0)
	fmt.Printf("Count Result: %v\n", countResult)

	// Recursion Max
	maxResult := recursion.RecursionMax(nums)
	fmt.Printf("Max Result: %v\n", maxResult)

	// Recursion Quicksort
	quicksortResult := recursion.Quicksort(nums)
	fmt.Printf("Quicksort Result: %v\n", quicksortResult)

	// Binary Serch
	sortedList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	binarySearchResult, err := search.BinarySearch(5, sortedList)

	if err != nil {
		fmt.Printf("Binary Search Error: %v\n", err)
	} else {
		fmt.Printf("Binary Search Result: %v\n", binarySearchResult)
	}

}
