package Methods

import "fmt"

func (v *Vertex) ScaleUsingPR(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func PointerReceiver() {
	fmt.Println("Using Pointer Receiver:")
	v := Vertex{3, 4}
	fmt.Println("\tBefore:", v)
	v.ScaleUsingPR(10) // Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.
	fmt.Println("\tAfter: ", v)
	fmt.Println("\t=> ", v.Abs())

	fmt.Println("======")

	p := &v
	fmt.Println("\tBefore:", v)
	p.ScaleUsingPR(10)
	fmt.Println("\tAfter: ", v)
	fmt.Println("\t=> ", p.Abs())

	fmt.Println("======")

	fmt.Println("Without using Pointer Receiver:")
	fmt.Println("\tBefore:", v)
	v.Scale(10)
	fmt.Println("\tAfter: ", v)
	fmt.Println("\t=> ", v.Abs())
}
