package Struct

import "fmt"

type Vertex struct {
	X, Y int
}

func Struct() {
	{
		v := Vertex{1, 2}
		fmt.Println("Before modification: ", v)
		v.X = 12
		fmt.Println("After modification: ", v)
	}

	{
		v := Vertex{1, 2}
		p := &v
		p.X = 1e9
		fmt.Println("Modification via pointer - p : ", p)
		fmt.Println("Modification via pointer - *p: ", *p)
		fmt.Println("Modification via pointer - v : ", v)
	}

	{
		v2 := Vertex{X: 1}           // {1, 0}
		v3 := Vertex{}               // {0, 0}
		p1 := &Vertex{1, 2}          // {1, 2}
		fmt.Println(v2, v3, p1, *p1) // {1 0} {0 0} &{1 2} {1 2}
	}
}
