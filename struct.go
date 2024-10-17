package main

//STRUCT are basically type, its similar to DTO's in nest js. it defines shape of an object.
type bill struct {
	name  string
	items map[string]int
	tip   float64
}

func createBill(n string) bill {
	x := bill{
		name: n,
		items: map[string]int{
			"biryani": 100,
			"water":   10,
		},
		tip: 19,
	}

	return x
}
