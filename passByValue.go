package main

import "fmt"

func updateNonPointer(n string) {
	n = "new value"
}

func updatedPointer(menu map[string]int) {
	menu["Pie"] = 600
}

func passByValueMain() {

	// here we declare a variable and the call the function to update it.
	// but when we print it, it sill displays the old value.

	//EXPLAINATION:

	/**
	There are two groups in which types lie.

	1. NON-POINTER VALUES
	 	a. String.
		b. Int.
		c. Floats.
		d. Booleans.
		e. Arrays.
		f. Structs

	2. POINTER WRAPPER VALUES
		a. slices.
		b. map.
		c. functions.

	=> When we store NON-POINTER variable in a mamory it gets stored in an address, and when we pass it
	to any other function a copy of the orignal variable is passed not the orignal var.

	That copy is also gets stored at any memory location. now when we change the var in that passed
	function, the copy is updated but the orignal variable will be as it is.

	=> When we store a POINTER variable, the value is stored at different address and a copy of that wil
	be soted at the variable address and a pointer will be stored which will point it to
	the orignal value.

	Now if we try to update it, it will update the orignal value.
	**/

	// NON-POINTER VALUES
	x := "value"
	updateNonPointer(x)
	fmt.Println(x) //res: "values"

	// POINTER WRAPPER VALUES
	menu := map[string]int{
		"Biryani": 100,
		"Mandi":   500,
	}

	fmt.Println(menu)
	updatedPointer(menu)
	fmt.Println(menu) //res: map[Biryani:100 Mandi:500 Pie:600]

}
