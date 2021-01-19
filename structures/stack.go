package structures

import "errors"

// StackNode represents a node in a stack
type StackNode struct {
	Next  *StackNode
	Value string
}

// Stack is a data structure
type Stack struct {
	Top *StackNode
}

// NewStack creates a new Stack
func NewStack() Stack {
	return Stack{}
}

// IsEmpty checks if the stack is empty
func (s Stack) IsEmpty() bool {
	return s.Top == nil
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	node := &StackNode{Value: str}
	if s.Top == nil {
		s.Top = node
		return
	}

	node.Next = s.Top
	s.Top = node
}

// Pop removes and returns the top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("the stack is empty, nothing to pop")
	}

	data := s.Top.Value
	s.Top = s.Top.Next
	return data, nil
}

// Peek fetches the value at the top of the stack
func (s *Stack) Peek() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("the stack is empty, nothing to pop")
	}

	return s.Top.Value, nil
}
