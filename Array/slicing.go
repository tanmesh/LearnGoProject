package Array

import "fmt"

//
// https://teivah.medium.com/slice-length-vs-capacity-in-go-af71a754b7d8
//

func Slicing() {
	a := []int{1, 2, 3, 4} // index starts from 0
	fmt.Println("a[:1] - ", a[:1])
	fmt.Println("a[2:] - ", a[2:])
	fmt.Println("a[:] - ", a[:])

	// a slice holds a pointer to the backing array plus a length and a capacity

	//b = b[1:4]  // Doest work because b size should be 4
	a = a[1:4] // a[low : high) -- including low index and excluding high index
	fmt.Println("Slicing :", a)
}

func LengthAndCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Println("Operation: `s = s[:0]`")
	s = s[:0]
	fmt.Printf("\tlen=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Println("Operation: `s[1]`")
	fmt.Println("\t=> Gives error")
	//fmt.Println(s[1]) // This gives error

	fmt.Println("Operation: `s = s[:4]`")
	s = s[:4]
	fmt.Printf("\tlen=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Println("\t=> Extend its length.")

	fmt.Println("Operation: `s = s[2:]`")
	s = s[2:]
	fmt.Printf("\tlen=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Println("\t=> Drop its first two values.")

	fmt.Println("Operation: `s = s[:4]`")
	s = s[:4]
	fmt.Printf("\tlen=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Println("\t=> Extend its length.")

	// Capacity doubles when it is full
	fmt.Println(s)
	s = append(s, 14, 15)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	s = append(s, 16)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}

func KeepingRef() {
	names := [4]string{"John", "Paul", "George", "Ringo"}
	println("Before: ")
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	fmt.Println("After: ")
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	fmt.Println("\t=> Changes are visible in a, b, names.")
	fmt.Println("\t   All these slices reference to the same backing array")

	fmt.Println("Operation: `s1 = append(s1, 2)")
	fmt.Println("            s2 := s1[1:3]`")
	s1 := make([]int, 3, 6)
	s2 := s1[1:3]
	printSlice("\ts1", s1)
	printSlice("\ts2", s2)
	println("\t=> s2 points to 1'st index of s1 but both have different capacity")

	fmt.Println("Operation: `s1[1] = 1`")
	s1[1] = 1
	printSlice("\ts1", s1)
	printSlice("\ts2", s2)
	fmt.Println("\t=> Changes are visible in both s1 and s2")

	fmt.Println("Operation: `s1 = append(s1, 2)`")
	s1 = append(s1, 2)
	printSlice("\ts1", s1)
	printSlice("\ts2", s2)
	fmt.Println("\t=> appending change is only visible in s1 and not in s2")

	fmt.Println("Operation: `Multiple append on s2`")
	s2 = append(s2, 2, 3, 4, 5)
	printSlice("\ts1", s1)
	printSlice("\ts2", s2)
	fmt.Println("\t=> Now, both s1 and s2 points to different array")
	fmt.Println("\t   s2 makes copy of the previous backing array and append new element to it")
}

func Make() {
	{
		fmt.Println("Operation: `a := make([]int, 5)`")
		a := make([]int, 5)
		printSlice("\ta", a)
		fmt.Println("\t=> length is 5")

		fmt.Println("Operation: `b := make([]int, 0, 5)`")
		b := make([]int, 0, 5)
		printSlice("\tb", b)
		fmt.Println("\t=> length is 0 and capacity is 5")

		fmt.Println("Operation: `c := b[:2]`")
		c := b[:2]
		printSlice("\tc", c)

		fmt.Println("Operation: `d := c[2:5]`")
		d := c[2:5]
		printSlice("\td", d)
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
