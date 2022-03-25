package stack

import "fmt"

type Stack[T any] struct {
	S []T
}

func NewStack[T any]() *Stack[T] {
	stack := make([]T, 0)
	return &Stack[T]{S: stack}
}
func (s *Stack[T]) Push(val T) {
	stack := s.S
	stack = append([]T(stack), val)
	s.S = stack
}
func (s *Stack[T]) Pop() T {
	var res T
	stack := s.S
	if len((stack)) > 0 {
		res = stack[len([]T(stack))-1] // element to be popped
		stack = stack[:len([]T(stack))-1]
		s.S = stack
	}
	return res
}
func (s *Stack[T]) String() string {
	stack := s.S
	res := fmt.Sprint(stack)
	return res
}

func (s *Stack[T]) IsEmpty() bool {
	if len([]T(s.S)) > 0 {
		return false
	}
	return true
}
