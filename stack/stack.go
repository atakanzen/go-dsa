package stack

type Stack struct {
	values  []int
	pointer int
}

func NewStack() *Stack {
	stack := Stack{
		values:  make([]int, 1),
		pointer: 0,
	}

	return &stack
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
