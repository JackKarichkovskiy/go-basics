package main

import (
	"fmt"
)

func fibonacci() func() int {
	n1, n2 := 0, 1

	return func() int {
		newResult := n1 + n2
		n1 = n2
		n2 = newResult
		return newResult
	}
}

func FibonacciMain() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}