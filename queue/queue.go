package queue

type Queue struct {
	values []interface{}
	start  int
	end    int
}

func NewQueue() *Queue {
	return &Queue{
		values: make([]interface{}, 0),
		start:  0,
		end:    0,
	}
}

func (q *Queue) Peek() interface{} {
	return q.values[q.start]
}

func (q *Queue) Push(item interface{}) {
	q.values = append(q.values, item)
	q.end++
}

func (q *Queue) Pop() {
	q.values = q.values[q.start:]
	q.start--
}

func (q *Queue) Values() []interface{} {
	return q.values
}
