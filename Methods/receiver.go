package Methods

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(math.Pow(float64(v.X), 2) + math.Pow(float64(v.Y), 2))
}

type MyFloat float64

// <receiver> <MethodName> <outputType>
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
func Receiver() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-2)
	fmt.Println(f.Abs())
}
