// Package stack provides a generic implementation of a stack data structure.
// The stack supports common operations such as pushing, popping, peeking at
// the top element, and checking if the stack is empty.
package stack

import "errors"

// Stack is a generic stack data structure that holds elements of type T.
type Stack[T any] struct {
	items []T
}

// Push adds an item to the top of the stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push - добавляет элемент в стек
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack.
func (s *Stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		var zero T
		return zero, errors.New("already empty") // Стек пуст
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, nil
}

// Peek returns the top item from the stack without removing it.
func (s *Stack[T]) Peek() (T, error) {
	if len(s.items) == 0 {
		var zero T
		return zero, errors.New("already empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty checks whether the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
