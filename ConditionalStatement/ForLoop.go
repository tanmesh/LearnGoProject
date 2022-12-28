package ConditionalStatement

import "fmt"

func ForLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func ForEachLoop() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	fmt.Println("Printing both index and value")
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	fmt.Println("Printing only indexes")
	for i := range pow {
		fmt.Println(i)
	}

	fmt.Println("Printing only values")
	for _, val := range pow {
		fmt.Println(val)
	}
}

func InfinitLoop() {
	sum := 0
	for {
		sum++ // repeated forever
		fmt.Println("Inside Loop...")
	}
	fmt.Println(sum) // never reached
}
