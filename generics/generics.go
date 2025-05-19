package main

import (
	"fmt"
)

func Index[T comparable](s []T, elem T) int {
	for i, v := range s {
		if v == elem {
			return i
		}
	}
	return -1
}

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(Index(primes, 7))
	fmt.Println(Index(primes, 4))

	strs := []string{"Hello", "World", "Go"}
	fmt.Println(Index(strs, "World"))
	fmt.Println(Index(strs, "Python"))
}