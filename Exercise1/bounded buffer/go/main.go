package main

import (
	"fmt"
	"time"
)

func producer(buf chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[producer]: pushing %d\n", i)
		buf <- i
		// TODO: push real value to buffer
	}
}

func consumer(buf chan int) {
	time.Sleep(1 * time.Second)
	for {
		i := <-buf
		fmt.Printf("[consumer]: %d\n", i)
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	// TODO: make a bounded buffer
	buf := make(chan int, 5)

	go consumer(buf)
	go producer(buf)

	select {}
}
