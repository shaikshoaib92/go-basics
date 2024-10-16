package main

import "fmt"

func mapMethod(n string) {
	menu := map[string]float64{
		"biryani": 100.25,
		"mandi":   400.35,
		"dal":     40,
	}

	fmt.Println(menu)
	fmt.Println(menu[n])
	menu["mandi"] = 5000 //we can modify the map
	fmt.Println(menu)

}
