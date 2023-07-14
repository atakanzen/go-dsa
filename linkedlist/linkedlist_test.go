package linkedlist_test

import (
	"algorithms/linkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedListCreation(t *testing.T) {
	newSinglyLinkedList := linkedlist.NewSinglyLinkedList(34, 15)

	assert.Equal(t, 34, newSinglyLinkedList.Head.Value)
	assert.Equal(t, 15, newSinglyLinkedList.Head.Next.Value)
	assert.Nil(t, newSinglyLinkedList.Head.Next.Next)
}

func TestSinglyLinkedListGet(t *testing.T) {
	newSinglyLinkedList := linkedlist.NewSinglyLinkedList(45, 15, 26, 12, 56, 34, 11, 65)

	actual := newSinglyLinkedList.Get(56)
	want := 56

	assert.Equal(t, want, actual)
}

func BenchmarkSinglyLinkedListGet(b *testing.B) {
	newSinglyLinkedList := linkedlist.NewSinglyLinkedList(45, 15, 26, 12, 56, 34, 11, 65)
	for i := 0; i < b.N; i++ {
		newSinglyLinkedList.Get(56)
	}
}
