package stack

type Stack struct {
	values  []int
	pointer int
}

func NewStack() *Stack {
	return &Stack{
		values:  make([]int, 1),
		pointer: 0,
	}
}

func (s *Stack) Push(pushedValue int) {
	s.values = append(s.values, pushedValue)
	s.pointer++
}

func (s *Stack) Pop() {
	s.values = s.values[:s.pointer]
	s.pointer--
}

func (s *Stack) Peek() int {
	return s.values[s.pointer]
}
