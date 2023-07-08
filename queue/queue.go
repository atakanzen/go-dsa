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
		end:    -1,
	}
}

func (q *Queue) Peek() interface{} {
	if len(q.values) == 0 {
		return nil
	}

	return q.values[q.start]
}

func (q *Queue) Push(item interface{}) {
	q.values = append(q.values, item)
	q.end += 1
}

func (q *Queue) Pop() bool {
	if len(q.values) != 0 {
		q.values[q.start] = nil
		copy(q.values[q.start:], q.values[q.start+1:len(q.values)])
		return true
	}

	return false
}

func (q *Queue) GetStart() interface{} {
	return q.values[q.start]
}

func (q *Queue) GetEnd() interface{} {
	return q.values[q.end]
}

func (q *Queue) Values() []interface{} {
	return q.values
}
