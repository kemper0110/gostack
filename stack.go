package main

type EmptyStackError struct{}

func (err *EmptyStackError) Error() string {
	return "Stack is empty"
}

type Node struct {
	Value interface{}
	Next  *Node
}
type Stack struct {
	head *Node
}

func NewStack() *Stack {
	return &Stack{nil}
}

func (s *Stack) top() (interface{}, error) {
	if s.head == nil {
		return nil, &EmptyStackError{}
	}
	return *s.head, nil
}

func (s *Stack) pop() (interface{}, error) {
	if s.head == nil {
		return nil, &EmptyStackError{}
	}
	tmp := s.head.Value
	s.head = s.head.Next
	return tmp, nil
}

func (s *Stack) push(value interface{}) {
	s.head = &Node{value, s.head}
}

func (s *Stack) empty() bool {
	return s.head == nil
}
