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

func TestQueueGetStart(t *testing.T) {
	newQueue := queue.NewQueue()

	newQueue.Push(1)
	newQueue.Push(2)
	newQueue.Push(3)

	newQueue.Pop()
	newQueue.Pop()

	actual := newQueue.GetStart()
	expected := 3

	assert.Equal(t, expected, actual)
}

func TestQueueGetEnd(t *testing.T) {
	newQueue := queue.NewQueue()

	newQueue.Push(3)
	newQueue.Push(4)

	actual := newQueue.GetEnd()
	expected := 4

	assert.Equal(t, expected, actual)
}

// This is way too slow
func BenchmarkQueueInsertAndRemove(b *testing.B) {
	newQueue := queue.NewQueue()
	for i := 0; i < b.N; i++ {
		newQueue.Push(i)
	}

	for i := 0; i < b.N; i++ {
		newQueue.Pop()
	}
}
