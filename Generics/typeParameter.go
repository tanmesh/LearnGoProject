package Generics

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func TypeParameter() {
	s1 := []int{10, 20, 30}
	fmt.Println(Index(s1, 20))

	s2 := []string{"foo", "bar"}
	fmt.Println(Index(s2, "hello"))
}
