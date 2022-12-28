package Function

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func FunctionValue() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("hypot: ", hypot(3, 4))

	fmt.Println("compute(hypot): ", compute(hypot))
	fmt.Println("compute(math.Pow): ", compute(math.Pow)) // return x**y => 3**4 = 81
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// A closure is a function value that references variables from outside its body.
// https://gobyexample.com/closures
func FunctionClosures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(i, "\t", pos(i), "\t", neg(-i))
	}
}
