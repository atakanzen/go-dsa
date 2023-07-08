package stack

import (
	"fmt"
	"strings"
)

type Stack struct {
	values  []interface{}
	pointer int
}

func NewStack() *Stack {
	return &Stack{
		values:  make([]interface{}, 0),
		pointer: -1,
	}
}

func (s *Stack) Push(pushedValue interface{}) {
	s.values = append(s.values, pushedValue)
	s.pointer++
}

func (s *Stack) Pop() bool {
	if len(s.values) == 0 {
		return false
	}

	s.values = s.values[:s.pointer]
	s.pointer--

	return true
}

func (s *Stack) Peek() interface{} {
	if len(s.values) == 0 {
		return nil
	}

	return s.values[s.pointer]
}

func (s *Stack) Values() []interface{} {
	return s.values
}

func (s *Stack) String() string {
	var sb strings.Builder

	for _, v := range s.values {
		sb.WriteString(fmt.Sprintf("%v, ", v))
	}

	return sb.String()
}
