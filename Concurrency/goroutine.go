package Concurrency

import "fmt"

func say(s string) {
	fmt.Println(s)
}

func Goroutine() {
	go say("world!")
	say("hello")
}
