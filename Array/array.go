package Array

import "fmt"

func findAverage(a []int) float64 {
	count := 4
	sum := 0
	for i := 0; i < count; i++ {
		sum += (a[i])
	}
	return float64(sum)
}

func Array() {
	{
		a := []int{1, 2, 3, 4} // index starts from 0 // SLICE
		fmt.Println(a)
		fmt.Println("AVERAGE", findAverage(a)) // ARRAY
		b := [4]int{1, 2, 3, 4}
		fmt.Println(b)
		//fmt.Println("AVERAGE", findAverage(b)) // Cannot use since findAverage() parameter didnt specified the array length
	}

	{
		// Array of struct
		s := []struct {
			i int
			b bool
		}{
			{2, true},
			{3, false},
			{5, true},
			{7, true},
			{11, false},
			{13, true},
		}
		fmt.Println("2D array: ", s)
	}
}

func Nill() {
	var a []int
	fmt.Println(a, len(a), cap(a))
	if a == nil {
		fmt.Println("its nil")
	}
}
