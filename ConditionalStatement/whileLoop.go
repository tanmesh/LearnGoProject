package ConditionalStatement

import "fmt"

func whileLoop() {
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n) // 8 (1*2*2*2)
}
