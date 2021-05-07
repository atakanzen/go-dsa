package recursion

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
