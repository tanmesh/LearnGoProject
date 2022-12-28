package Map

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func Map() {
	{
		m = make(map[string]Vertex)
		m["Bell Lab"] = Vertex{40.68433, -74.39967}
		fmt.Println("Operation: `print(m)`")
		fmt.Println("\t=>", m)

		fmt.Println("Operation: `print(m[\"Bell Lab\"])`")
		fmt.Println("\t=>", m["Bell Lab"])

		m["Google"] = Vertex{37.42202, -122.08408}

		fmt.Println("Operation: `print(m)` after adding new key-val")
		fmt.Println("\t=>", m)
	}

	{
		var m = map[string]Vertex{
			"Bell Labs": Vertex{40.68433, -74.39967},
			"Google":    Vertex{37.42202, -122.08408},
		}
		fmt.Println("Multiple K-V pairs: ", m)
	}

	{
		var m = map[string]Vertex{
			"Bell Labs": {40.68433, -74.39967},
			"Google":    {37.42202, -122.08408},
		}
		fmt.Println("Without using vertex: ", m)
	}
}

func MutatingMap() {
	m := map[string]int{
		"Person1_Score": 42,
		"Person2_Score": 42,
	}
	fmt.Println(m)

	m["Person1_Score"] = 48
	fmt.Println("Operation: `m[\"Person1_Score\"] = 48`")
	fmt.Println("\t=> ", m)

	delete(m, "Person1_Score")
	fmt.Println("Operation: `delete(m, \"Person1_Score\")`")
	fmt.Println("\t=> ", m)

	v, ok := m["Person1_Score"]
	fmt.Println("\t=> The value: ", v, "Present?", ok)
}
