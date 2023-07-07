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
