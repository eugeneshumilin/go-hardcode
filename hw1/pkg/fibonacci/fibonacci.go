package fibonacci

// Fib функция для вывода числа Фибоначчи
func Fib(n int) int {
	var a, b int = 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return b - a
}
