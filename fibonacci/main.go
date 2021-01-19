package fibonacci

// Run returns the nth number of a fibonacci sequence
func Run(n int) int {
	if n <= 1 {
		return n
	}

	return Run(n-1) + Run(n-2)
}
