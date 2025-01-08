// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
)

const (
	INCREMENT = iota
	DECREMENT
)

type numberServer struct {
	messages        chan int
	finish          chan int
	data            int
	finishedWorkers int
	totalWorkers    int
}

var server = numberServer{make(chan int), make(chan int), 0, 0, 2}

func incrementing() {
	for n := 0; n < 1000000; n++ {
		server.messages <- INCREMENT
	}
	server.finish <- 1
}

func decrementing() {
	for n := 0; n < 1000000; n++ {
		server.messages <- DECREMENT
	}
	server.finish <- 1
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(2)

	go incrementing()
	go decrementing()

	// Update the server
eventLoop:
	for {
		select {
		case op := <-server.messages:
			switch op {
			case INCREMENT:
				server.data++
			case DECREMENT:
				server.data--
			}
		case <-server.finish:
			server.finishedWorkers++
			if server.finishedWorkers >= server.totalWorkers {
				Println("The magic number is:", server.data)
				break eventLoop
			}
		}
	}
}
