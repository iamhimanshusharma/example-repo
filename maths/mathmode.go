package maths

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
