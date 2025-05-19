package main

import (
	"fmt"
	"time"
)

func generator() func () int {
	i := 0
	return func() (res int) {
		res = i
		i++
		return
	}
}

func worker(mainCh, quit chan int) {
	nextInt := generator()
	for {
		select {
			case mainCh <- nextInt():
				time.Sleep(1 * time.Second)
			case <-quit:
				fmt.Println("Worker quit")
				close(mainCh)
				return
		}
	}	
}

func main() {
	mainCh := make(chan int)
	quit := make(chan int)

	go worker(mainCh, quit)
	go func() {
		for {
			nextInt, ok := <-mainCh
			if ok {
				fmt.Println("Received:", nextInt)
			} else {
				fmt.Println("Main channel closed")
				return
			}
		}
	}()
	go func() {
		time.Sleep(5 * time.Second)
		quit <- 0
	}()


	time.Sleep(10 * time.Second)
	fmt.Println("Main quit")
}