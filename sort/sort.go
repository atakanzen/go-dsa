package sort

import (
	"math/rand"
)

// O(N^2)
func InsertionSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		current := i
		for current > 0 && list[current] < list[current-1] {
			temp := list[current]
			list[current] = list[current-1]
			list[current-1] = temp
			current--
		}
	}

	return list
}

func SelectionSort(list []int) []int {
	for i := range list {
		minIndex := i
		for j := i; j < len(list); j++ {
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}

		temp := list[i]
		list[i] = list[minIndex]
		list[minIndex] = temp
	}

	return list
}

func BubbleSort(list []int) []int {
	if len(list) == 1 {
		return list
	}

	for i := 0; i < len(list); i++ {
		swapped := false
		for j := range list {
			if j < len(list)-1 && list[j] > list[j+1] {
				temp := list[j]
				list[j] = list[j+1]
				list[j+1] = temp
				swapped = true
			}
		}

		if !swapped {
			return list
		}
	}

	return list
}

func Mergesort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2

	leftList := Mergesort(nums[:mid])
	rightList := Mergesort(nums[mid:])
	sortedList := make([]int, 0)

	sortedList = mergeLists(leftList, rightList, sortedList)

	return sortedList
}

func mergeLists(leftList []int, rightList []int, sortedList []int) []int {
	for leftPointer, rightPointer := 0, 0; leftPointer < len(leftList) || rightPointer < len(rightList); {
		if leftPointer == len(leftList) {
			sortedList = append(sortedList, rightList[rightPointer])
			rightPointer++
		} else if rightPointer == len(rightList) {
			sortedList = append(sortedList, leftList[leftPointer])
			leftPointer++
		} else if leftList[leftPointer] <= rightList[rightPointer] {
			sortedList = append(sortedList, leftList[leftPointer])
			leftPointer++
		} else {
			sortedList = append(sortedList, rightList[rightPointer])
			rightPointer++
		}
	}
	return sortedList
}

func Quicksort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	left, right := 0, len(nums)-1

	pivot := rand.Intn(len(nums))
	nums[pivot], nums[right] = nums[right], nums[pivot]
	pivot = right

	for i := 0; i < pivot; i++ {
		if nums[i] < nums[pivot] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}

	nums[left], nums[pivot] = nums[pivot], nums[left]

	Quicksort(nums[:left])
	Quicksort(nums[left+1:])

	return nums
}
