package main

import (
	"flag"
	"fmt"
	"go-hardcode/hw1/pkg/fibonacci"
)

func main() {
	n := flag.Int("n", 1, "fib number")
	flag.Parse()
	if *n > 20 {
		fmt.Println("value cannot be greater than 20")
		return
	}

	f := fibonacci.Fib(*n)

	fmt.Println(f)
}
