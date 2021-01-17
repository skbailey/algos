package structures

import "fmt"

// Node is an element of a LinkedList
type Node struct {
	Value int
	Next  *Node
}

// LinkedList a data structure
type LinkedList struct {
	Head *Node
}

// Append adds a node to the linked list
func (list *LinkedList) Append(value int) {
	last := &Node{Value: value}

	if list.Head == nil {
		list.Head = last
		return
	}

	current := list.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = last
}

// Delete removes nodes from the linked list
func (list *LinkedList) Delete(value int) {
	if list.Head == nil {
		return
	}

	if list.Head.Value == value {
		list.Head = list.Head.Next
		return
	}

	current := list.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			return
		}

		current = current.Next
	}
}

// Traverse prints out all values in the linked list
func (list *LinkedList) Traverse() {
	if list.Head == nil {
		return
	}

	current := list.Head
	fmt.Println(current.Value)

	for current.Next != nil {
		fmt.Println(current.Next.Value)
		current = current.Next
	}
}

// NewLinkedList creates a new LinkedList
func NewLinkedList() LinkedList {
	return LinkedList{}
}
