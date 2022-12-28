package Concurrency

import "fmt"

func Channels() {
	s := []int{7, 2, 8, -9, 4, 0}
	a := s[:len(s)/2]
	b := s[len(s)/2:]
	fmt.Println(s)
	fmt.Println(a, b)

	c := make(chan int) // make channel before using it
	go sum(a, c)
	go sum(b, c)
	x, y := <-c, <-c // receive from channel c

	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to channel c
}

func Ping() {
	messages := make(chan string)
	go func() { messages <- "ping" }()
	fmt.Println(<-messages)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// A sender can close a channel to indicate that no more values will be sent.
// Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after
func RangeAndClose() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	val, ok := <-c
	if ok == true {
		fmt.Println("ok = true ->  Channel is close")
	}

	fmt.Println(val)
	for i := range c {
		fmt.Println(i)
	}
}
