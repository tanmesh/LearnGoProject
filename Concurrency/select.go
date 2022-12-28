package Concurrency

import (
	"fmt"
	"time"
)

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func Select() {
	c := make(chan int, 10)
	quit := make(chan int)
	go func() {
		for {
			fmt.Println(<-c)
		}
		//quit <- 0
	}()
	fibonacci2(c, quit)
}

func DefaultSelect() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
