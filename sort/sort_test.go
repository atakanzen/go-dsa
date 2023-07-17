package sort_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/atakanzen/go-dsa/sort"
	"github.com/stretchr/testify/assert"
)

// Used for Benchmarks
func generateRandomSlice(size int) []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1000) // Modify the range as needed
	}
	return slice
}

func TestInsertionSort(t *testing.T) {
	testCases := []struct {
		desc        string
		input, want []int
	}{
		{
			desc:  "sort 3 items",
			input: []int{5, 2, 1},
			want:  []int{1, 2, 5},
		},
		{
			desc:  "sort 6 items",
			input: []int{5, 2, 1, 6, 3, 4},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			desc:  "sort 12 items",
			input: []int{5, 2, 1, 6, 3, 4, 12, 9, 8, 11, 7, 10},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			desc:  "sort 12 items with 3 duplicate",
			input: []int{5, 2, 3, 6, 3, 4, 12, 3, 8, 11, 7, 10},
			want:  []int{2, 3, 3, 3, 4, 5, 6, 7, 8, 10, 11, 12},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := sort.InsertionSort(tC.input)
			assert.Equal(t, tC.want, actual)
		})
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	testCases := []struct {
		desc  string
		input []int
	}{
		{
			desc:  "sort 3 items",
			input: []int{5, 2, 1},
		},
		{
			desc:  "sort 6 items",
			input: []int{5, 2, 1, 6, 3, 4},
		},
		{
			desc:  "sort 12 items",
			input: []int{5, 2, 1, 6, 3, 4, 12, 9, 8, 11, 7, 10},
		},
		{
			desc:  "sort 12 items with 3 duplicate",
			input: []int{5, 2, 3, 6, 3, 4, 12, 3, 8, 11, 7, 10},
		},
	}

	for _, tC := range testCases {
		b.Run(tC.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.InsertionSort(tC.input)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	testCases := []struct {
		desc        string
		input, want []int
	}{
		{
			desc:  "sort 3 items",
			input: []int{5, 2, 1},
			want:  []int{1, 2, 5},
		},
		{
			desc:  "sort 6 items",
			input: []int{5, 2, 1, 6, 3, 4},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			desc:  "sort 12 items",
			input: []int{5, 2, 1, 6, 3, 4, 12, 9, 8, 11, 7, 10},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			desc:  "sort 12 items with 3 duplicate",
			input: []int{5, 2, 3, 6, 3, 4, 12, 3, 8, 11, 7, 10},
			want:  []int{2, 3, 3, 3, 4, 5, 6, 7, 8, 10, 11, 12},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := sort.SelectionSort(tC.input)
			assert.Equal(t, tC.want, actual)
		})
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	testCases := []struct {
		desc  string
		input []int
	}{
		{
			desc:  "sort 3 items",
			input: []int{5, 2, 1},
		},
		{
			desc:  "sort 6 items",
			input: []int{5, 2, 1, 6, 3, 4},
		},
		{
			desc:  "sort 12 items",
			input: []int{5, 2, 1, 6, 3, 4, 12, 9, 8, 11, 7, 10},
		},
		{
			desc:  "sort 12 items with 3 duplicate",
			input: []int{5, 2, 3, 6, 3, 4, 12, 3, 8, 11, 7, 10},
		},
	}

	for _, tC := range testCases {
		b.Run(tC.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.SelectionSort(tC.input)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
	testCases := []struct {
		desc        string
		input, want []int
	}{
		{
			desc:  "sort 3 items",
			input: []int{5, 2, 1},
			want:  []int{1, 2, 5},
		},
		{
			desc:  "sort 6 items",
			input: []int{5, 2, 1, 6, 3, 4},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			desc:  "sort 12 items",
			input: []int{5, 2, 1, 6, 3, 4, 12, 9, 8, 11, 7, 10},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			desc:  "sort 12 items with 3 duplicate",
			input: []int{5, 2, 3, 6, 3, 4, 12, 3, 8, 11, 7, 10},
			want:  []int{2, 3, 3, 3, 4, 5, 6, 7, 8, 10, 11, 12},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := sort.BubbleSort(tC.input)
			assert.Equal(t, tC.want, actual)
		})
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	testCases := []struct {
		desc  string
		input []int
	}{
		{
			desc:  "sort 3 items",
			input: []int{5, 2, 1},
		},
		{
			desc:  "sort 6 items",
			input: []int{5, 2, 1, 6, 3, 4},
		},
		{
			desc:  "sort 12 items",
			input: []int{5, 2, 1, 6, 3, 4, 12, 9, 8, 11, 7, 10},
		},
		{
			desc:  "sort 12 items with 3 duplicate",
			input: []int{5, 2, 3, 6, 3, 4, 12, 3, 8, 11, 7, 10},
		},
	}

	for _, tC := range testCases {
		b.Run(tC.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.BubbleSort(tC.input)
			}
		})
	}
}

func TestQuickSortSingleItemArray(t *testing.T) {
	singleItemArray := []int{10}

	quicksortResult := sort.Quicksort(singleItemArray)

	assert.Equal(t, singleItemArray, quicksortResult)
}

func TestQuickSortingMultipleItemArray(t *testing.T) {
	multipleItemArray := []int{9, -3, 2, 100, 4}

	quicksortResult := sort.Quicksort(multipleItemArray)
	expectedResult := []int{-3, 2, 4, 9, 100}

	assert.Equal(t, expectedResult, quicksortResult)
}

func BenchmarkQuicksort(b *testing.B) {
	testCases := []struct {
		desc  string
		input []int
	}{
		{
			desc:  "sort 150 items",
			input: generateRandomSlice(150),
		},
		{
			desc:  "sort 1500 items",
			input: generateRandomSlice(1500),
		},
		{
			desc:  "sort 15000 items",
			input: generateRandomSlice(15000),
		},
	}

	for _, tC := range testCases {
		b.Run(tC.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.Quicksort(tC.input)
			}
		})
	}
}

func TestMergesort(t *testing.T) {
	testCases := []struct {
		desc        string
		input, want []int
	}{
		{
			desc:  "sort single item array",
			input: []int{5},
			want:  []int{5},
		},
		{
			desc:  "sort 3 item array",
			input: []int{45, 34, 5},
			want:  []int{5, 34, 45},
		},
		{
			desc:  "sort 10 item array",
			input: []int{5, 34, 45, 11, 2, 14, 56, 87, 10, 3},
			want:  []int{2, 3, 5, 10, 11, 14, 34, 45, 56, 87},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := sort.Mergesort(tC.input)
			assert.Equal(t, tC.want, actual)
		})
	}
}
