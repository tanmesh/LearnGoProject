package Errors

import (
	"fmt"
	"math"
)

type MyError2 struct {
	What string
}

func (e *MyError2) Error() string {
	return fmt.Sprint(e.What)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, &MyError2{"Less than zero"}
	}
	return math.Sqrt(x), nil
}

func NegativeRootExample() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
