package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeList(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func TestBinarySearchFound(t *testing.T) {
	list := makeList(0, 500000)
	item := 221017

	position, err := BinarySearch(item, list)

	assert.Equal(t, item, position)
	assert.ErrorIs(t, err, nil)
}

func TestBinarySearchNotFound(t *testing.T) {
	list := makeList(0, 1)
	item := 311098

	position, err := BinarySearch(item, list)

	assert.NotEqualValues(t, item, position)
	assert.Error(t, err)
}
