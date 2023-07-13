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
