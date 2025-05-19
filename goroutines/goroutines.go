package main

import (
	"fmt"
	"time"
)

func sayWithDelay(s string) {
	time.Sleep(1 * time.Second)
	fmt.Println(s)
}

func main() {
	go sayWithDelay("Hello")
	sayWithDelay("World")
	// Wait for goroutine to finish
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}