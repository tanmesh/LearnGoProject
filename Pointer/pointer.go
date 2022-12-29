package Pointer

func Pointer() {
	a := 10
	println("a := 10, => a = ", a)
	p := &a
	println("p = &a, => p = ", *p)
	*p = 5
	println("*p = 5, => *p = ", *p)
}
