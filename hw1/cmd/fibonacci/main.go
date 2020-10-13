package main

import (
	"fmt"
	"go-hardcode/hw1/pkg/fibonacci"
)

func main() {

	for n := -1; n < 25; n ++ {
		f, err := fibonacci.Fib(n)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(f)
	}
}