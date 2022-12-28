package ConditionalStatement

import "fmt"

// A defer statement defers the execution of a function until the surrounding function returns.
func Defer2() {
	defer fmt.Println("world")
	fmt.Println("hello")
	fmt.Println("!")
	// hello
	// !
	fmt.Println("=====")
	// world
}

func Defer1() {
	defer fmt.Println("world")
	fmt.Println("hello")
	// hello
	fmt.Println("=====")

	defer fmt.Println("world")
	fmt.Println("!")
	fmt.Println("hello")
	// !
	// hello
	fmt.Println("=====")

	defer fmt.Println("world")
	fmt.Println("hello")
	fmt.Println("!")
	// hello
	// !
	fmt.Println("=====")
	// world
	// world
	// world
}

func StackDefer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
