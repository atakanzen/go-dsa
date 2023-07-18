package linkedlist_test

import (
	"testing"

	"github.com/atakanzen/go-dsa/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedListCreation(t *testing.T) {
	newDoublyLinkedList := linkedlist.NewDoublyLinkedList(34, 15)

	assert.Equal(t, 34, newDoublyLinkedList.Head.Value)
	assert.Equal(t, 15, newDoublyLinkedList.Tail.Value)
	assert.Equal(t, 34, newDoublyLinkedList.Tail.Prev.Value)
	assert.Nil(t, newDoublyLinkedList.Head.Prev)
	assert.Nil(t, newDoublyLinkedList.Tail.Next)
}

func TestDoublyLinkedListGet(t *testing.T) {
	newDoublyLinkedList := linkedlist.NewDoublyLinkedList(45, 15, 26, 12, 56, 34, 11, 65)

	actual := newDoublyLinkedList.Get(11)
	wantValue := 11
	wantPrevValue := 34
	wantNextValue := 65

	assert.Equal(t, wantValue, actual.Value)
	assert.Equal(t, wantPrevValue, actual.Prev.Value)
	assert.Equal(t, wantNextValue, actual.Next.Value)
}

func TestDoublyLinkedListPrepend(t *testing.T) {
	testCases := []struct {
		desc      string
		toPrepend []int
		want      int
	}{
		{
			desc:      "prepend 5 items with last being 10",
			toPrepend: []int{27, 3, 22, 43, 10},
			want:      10,
		},
		{
			desc:      "prepend 1 item",
			toPrepend: []int{23},
			want:      23,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			newDoublyLinkedList := linkedlist.NewDoublyLinkedList()
			for _, v := range tC.toPrepend {
				newDoublyLinkedList.Prepend(v)
			}

			assert.Equal(t, tC.want, newDoublyLinkedList.Head.Value)
		})
	}
}

func TestDoublyLinkedListRemove(t *testing.T) {
	testCases := []struct {
		desc                       string
		initialValues, valuesAfter []interface{}
		removeItem                 int
		want                       bool
	}{
		{
			desc:          "should remove 24 from 5 values",
			initialValues: []interface{}{45, 15, 24, 14, 55},
			removeItem:    24,
			valuesAfter:   []interface{}{45, 15, 14, 55},
			want:          true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			doubleLinkedList := linkedlist.NewDoublyLinkedList(tC.initialValues...)

			actual := doubleLinkedList.Remove(tC.removeItem)
			values := doubleLinkedList.Values()
			assert.Equal(t, tC.want, actual)
			assert.Equal(t, tC.valuesAfter, values)
		})
	}
}

func BenchmarkDoublyLinkedListGet(b *testing.B) {
	newDoublyLinkedList := linkedlist.NewDoublyLinkedList(45, 15, 26, 12, 56, 34, 11, 65)
	for i := 0; i < b.N; i++ {
		newDoublyLinkedList.Get(56)
	}
}
