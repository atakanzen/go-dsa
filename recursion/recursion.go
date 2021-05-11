package recursion

import (
	"math"
	"math/rand"
)

func RecursionSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return nums[0] + RecursionSum(nums[1:])
}

func RecursionCount(items []string, count int) int {
	if len(items) == 0 {
		return 0
	}
	count++
	return count + RecursionCount(items[count:], count)
}

func RecursionMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	return int(math.Max(float64(nums[0]), float64(RecursionMax(nums[1:]))))
}

func Quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	left, right := 0, len(nums)-1

	pivot := rand.Intn(len(nums))

	nums[pivot], nums[right] = nums[right], nums[pivot]

	for i := range nums {
		if nums[i] < nums[right] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}

	nums[left], nums[right] = nums[right], nums[left]

	Quicksort(nums[:left])
	Quicksort(nums[left+1:])

	return nums
}
