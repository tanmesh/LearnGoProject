package Concurrency

import "fmt"

// Channel and BufferChannel
// Provide the buffer length as the second argument to make to initialize a buffered channe
// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
func BufferedChannel() {
	c := make(chan int, 5)

	for i := 0; i < 5; i++ {
		c <- i
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}
