package Errors

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v %v", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "it didn't work"}
}

func Error() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
