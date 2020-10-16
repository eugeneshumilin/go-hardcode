package main

import (
	"errors"
	"flag"
	"fmt"
	"go-hardcode/hw1/pkg/fibonacci"
)

func main() {
	n := flag.Int("n", 0, "fib number")
	flag.Parse()
	if *n > 20 {
		fmt.Println(errors.New("value cannot be greater than 20"))
		return
	}

	f, err := fibonacci.Fib(*n)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)
}
