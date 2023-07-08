package queue_test

import (
	"algorithms/queue"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueWithNoItems(t *testing.T) {
	newQueue := queue.NewQueue()

	assert.Empty(t, newQueue.Values())
}

func TestQueuePopWhenEmpty(t *testing.T) {
	newQueue := queue.NewQueue()

	actual := newQueue.Pop()
	actualTwo := newQueue.Pop()

	assert.False(t, actual)
	assert.False(t, actualTwo)
}

func TestQueuePeekWhenEmpty(t *testing.T) {
	newQueue := queue.NewQueue()

	actual := newQueue.Peek()

	assert.Empty(t, actual)
}

func TestQueueReturnsFirstItem(t *testing.T) {
	newQueue := queue.NewQueue()

	newQueue.Push("foo")
	newQueue.Push("bar")
	newQueue.Push("baz")

	newQueue.Pop()

	actual := newQueue.Peek()
	expected := "bar"

	assert.Equal(t, expected, actual)
}
