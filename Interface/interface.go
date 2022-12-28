package Interface

import (
	"fmt"
	"learn-go-project/Methods"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

type F float64

// Implementation of method M() of Interface I
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func (f F) M() {
	fmt.Println(f)
}

type Abser interface {
	Abs() float64
}

func Interface() {
	{
		var a Abser
		f := Methods.MyFloat(-2)
		v := &Methods.Vertex{3, 4}

		a = v
		fmt.Println(a.Abs())

		a = f
		fmt.Println(a.Abs())
	}

	{
		var i I = &T{"hello"}
		i.M()
		fmt.Printf("%v %T \n", i, i) // Value Type

		var i2 = F(math.Pi)
		i2.M()
		fmt.Printf("%v %T \n", i2, i2) // Value Type
	}

	{
		var t *T
		var i I = t
		fmt.Printf("%v %T \n", i, i)

		i = &T{"hello again"}
		fmt.Printf("%v %T \n", i, i)
	}

	{
		var i I
		fmt.Printf("%v %T \n", i, i)
		//i.M() // This gives run-time error as no type to indicate which method to call
	}

	{
		var i interface{}
		fmt.Printf("%v %T \n", i, i)

		i = 42
		fmt.Printf("%v %T \n", i, i)

		i = "hello"
		fmt.Printf("%v %T \n", i, i)
	}
}
