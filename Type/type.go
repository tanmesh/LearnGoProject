package Type

import "fmt"

func TypeAssertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Printf("(%v, %v) \n", s, ok)

	f, ok := i.(float64)
	fmt.Printf("(%v, %v) \n", f, ok)

	//f = i.(float64) // This will create panic since ok=false
}

func TypeSwitch() {
	do(12)
	do(12.34)
	do("Hello")
	do(true)
}

func do(i interface{}) {
	switch val := i.(type) {
	case int:
		fmt.Println("Value is Int ", val)
	case float64:
		fmt.Println("Value is Float64 ", val)
	case string:
		fmt.Println("Value is String ", val)
	case bool:
		fmt.Println("Value is Bool ", val)
	default:
		fmt.Println("Dont know the type")

	}
}
