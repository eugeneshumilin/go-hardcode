package fibonacci

import "errors"

// Fib function
func Fib(n int) (int, error) {
	if n > 20 || n == 0 || n < 0 {
		return -1, errors.New("value cannot be greater than 20 and less than 0")
	}
	var a, b int = 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return b - a, nil
}