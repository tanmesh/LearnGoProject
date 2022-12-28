package Interface

import "fmt"

//type Stringer interface {
//	String() string
//}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}

func Stringer() {
	a := Person{"Arthur Dent", 42}
	b := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, b)
}
