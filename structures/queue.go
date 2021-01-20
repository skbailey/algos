package structures

import "errors"

// QueueNode represents a node in a queue
type QueueNode struct {
	Next  *QueueNode
	Value string
}

// Queue is a data structure
type Queue struct {
	First *QueueNode
	Last  *QueueNode
}

// NewQueue creates a new Queue
func NewQueue() Queue {
	return Queue{}
}

// IsEmpty checks if the stack is empty
func (q Queue) IsEmpty() bool {
	return q.First == nil
}

// Peek fetches the value at the top of the stack
func (q *Queue) Peek() (string, error) {
	if q.IsEmpty() {
		return "", errors.New("the stack is empty, nothing to pop")
	}

	return q.First.Value, nil
}

// Add places an item at the back of the Queue
func (q *Queue) Add(value string) {
	node := &QueueNode{Value: value}

	if q.First == nil {
		q.First = node
	}

	if q.Last != nil {
		q.Last.Next = node
	}
	q.Last = node
}

// Remove removes an item from the front of the Queue
func (q *Queue) Remove() (string, error) {
	if q.IsEmpty() {
		return "", errors.New("the queue is empty, nothing to remove")
	}

	value := q.First.Value
	q.First = q.First.Next

	return value, nil
}
