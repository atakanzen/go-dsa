package sort_test

import (
	"algorithms/sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
