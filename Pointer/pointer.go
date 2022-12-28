package Pointer

func Pointer() {
	a := 10
	println("Before modification: ", a)
	p := &a
	println("Via pointer: ", *p)
	*p = 5
	println("After modification: ", *p)
}
