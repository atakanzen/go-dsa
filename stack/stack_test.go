package stack_test

import (
	"algorithms/stack"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPopWhenEmpty(t *testing.T) {
	newStack := stack.NewStack()

	actual := newStack.Pop()

	assert.False(t, actual)
}

func TestStackPeekWhenEmpty(t *testing.T) {
	newStack := stack.NewStack()

	actual := newStack.Peek()

	assert.Nil(t, actual)
}

func TestStackPeekWithItems(t *testing.T) {
	newStack := stack.NewStack()

	newStack.Push("baz")
	newStack.Push("bar")
	newStack.Push("foo")

	newStack.Pop()
	newStack.Pop()

	actual := newStack.Peek()
	expected := "baz"

	assert.Equal(t, expected, actual)
}

func BenchmarkStack(b *testing.B) {
	newStack := stack.NewStack()

	for i := 0; i < b.N; i++ {
		newStack.Push(i)
	}

}
