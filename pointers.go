package main

import "fmt"

//Pointers as basically address location.

/**
NOTE:
1. "&" is used to get the pointer value / address of a variable.
2. "*" is used to derefrence the pointer value / or we can say we can get the value of
		that address by derefrencing.
*/

func updateOrignalValue(n *string) { // *string represent that it accepts the pointer
	*n = "new value" //Here we are derefrencing the pointer, and assigning a new value to that pointer
}

func main() {

	x := "value" //since its a string we cannot update it direclty.

	//To update it we first find its pointer value and then derefrence it.

	m := &x //To get the pointer value of x we use "&", here m also has its pointer value.

	/**
	|------x-------|-------m--------|
	| 0xc00002c0a0 | 0xc000060050   |
	|--------------|----------------|
	|     value	   |  0xc00002c0a0  |
	|--------------|----------------|
	*/

	//Now if we want to update the value of x, we pass the pointer to function.
	//That funcation will defrefence the pointer and then assing a new value to that pointer

	updateOrignalValue(m) //we are passing in the pointer value

	fmt.Println(x) //res: new value

}
