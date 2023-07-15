package recursion

import (
	"math"
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
