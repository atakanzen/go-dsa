package search

import (
	"errors"
	"math"
)

func BinarySearch(item int, list []int) (int, error) {
	low := 0
	high := len(list) - 1
	for i := 0; i <= high; i++ {
		mid := int(math.Round(float64((low + high) / 2)))
		guess := list[mid]

		if guess == item {
			return mid, nil
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0, errors.New("item does not exist")
}
