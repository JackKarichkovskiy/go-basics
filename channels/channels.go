package main

import "fmt"

func sum(slice []int, ch chan int) {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	fmt.Println("Calculater sum:", sum)
	ch <- sum
}

func worker(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	input := []int{1, 20, 3, 40, 5, 89, 23}
	ch := make(chan int)

	go sum(input[:len(input)/2], ch)
	go sum(input[len(input)/2:], ch)

	result1 := <-ch
	result2 := <-ch
	fmt.Println("Overall sum:", result1 + result2)

	// Buffered channel example
	fmt.Println("\nBuffered channel example")
	ch2 := make(chan int, 2) // modify size to 1 to have a deadlock
	ch2 <- 1
	ch2 <- 2
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	// Channel close
	fmt.Println("\nChannel close example")
	ch3 := make(chan int, 10)
	go worker(ch3)
	for val, closed := <-ch3; closed; val, closed = <-ch3 {
		fmt.Println(val)
	}
	fmt.Println("Channel closed")
}