package sort

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
