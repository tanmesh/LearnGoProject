package Methods

import (
	"fmt"
	"math"
)

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsWithPointer(v *Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func ScaleWithPointer(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Function() {
	v := Vertex{3, 4}
	fmt.Println("Before:", v)
	ScaleWithPointer(&v, 10)
	fmt.Println("After:", v)
	fmt.Println("=> Pass by Ref: ", Abs(v))

	fmt.Println("====")

	fmt.Println("Before:", v)
	Scale(v, 10)
	fmt.Println("After:", v)
	fmt.Println("=> Pass by Value: ", Abs(v))

	fmt.Println("====")

	w := &Vertex{6, 8}
	fmt.Println("Before:", w)
	ScaleWithPointer(w, 10)
	fmt.Println("After:", w)
	//fmt.Println("=> Pass by Value: ", Abs(w)) // This gives error
	fmt.Println("=> Pass by Ref: ", AbsWithPointer(w))
}
