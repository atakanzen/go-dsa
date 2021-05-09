package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecursionSumEmptyArray(t *testing.T) {
	emptyArray := []int{}

	sumResult := RecursionSum(emptyArray)

	assert.Equal(t, 0, sumResult)
}

func TestRecursionSumValidArray(t *testing.T) {
	validArray := []int{1, 2, 3}

	sumResult := RecursionSum(validArray)

	assert.Equal(t, 6, sumResult)
}

func TestRecursionCountEmptyArray(t *testing.T) {
	emptyArray := []string{}

	countResult := RecursionCount(emptyArray, 0)

	assert.Equal(t, 0, countResult)
}

func TestRecursionCountValidArray(t *testing.T) {
	validArray := []string{"f", "o", "o"}

	countResult := RecursionCount(validArray, 0)

	assert.Equal(t, 3, countResult)
}

func TestRecursionMaxEmptyArray(t *testing.T) {
	emptyArray := []int{}

	maxResult := RecursionMax(emptyArray)

	assert.Equal(t, 0, maxResult)
}

func TestRecursionMaxSingleItemArray(t *testing.T) {
	singleItemArray := []int{10}

	maxResult := RecursionMax(singleItemArray)

	assert.Equal(t, 10, maxResult)
}

func TestRecursionMaxMultipleItemArray(t *testing.T) {
	multipleItemArray := []int{1, 6, 2, 99}

	maxResult := RecursionMax(multipleItemArray)

	assert.Equal(t, 99, maxResult)
}

func TestQuickSortSingleItemArray(t *testing.T) {
	singleItemArray := []int{10}

	quicksortResult := Quicksort(singleItemArray)

	assert.Equal(t, singleItemArray, quicksortResult)
}

func TestQuickSortingMultipleItemArray(t *testing.T) {
	multipleItemArray := []int{9, -3, 2, 100, 4}

	quicksortResult := Quicksort(multipleItemArray)
	expectedResult := []int{-3, 2, 4, 9, 100}

	assert.Equal(t, expectedResult, quicksortResult)
}
