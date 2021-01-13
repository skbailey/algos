package structures

// Stack is a data structure
type Stack []string

// NewStack creates a new Stack
func NewStack() Stack {
	return Stack{}
}

// IsEmpty checks if the stack is empty
func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Pop removes and returns the top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	(*s) = (*s)[:index]    // Remove it from the stack by slicing it off.
	return element, true
}

func main() {
}
