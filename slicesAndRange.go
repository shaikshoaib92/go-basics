package main

import "fmt"

func slicesExample() {

	//SLICES

	/**
	NOTE:
	1. Unlike arrays slices can be extented, you can add more values to them.
	2. Declaring a slice is almost similar to arry but here we dont specify the size of array.
	**/

	var scores = []int{10, 20, 30}

	fmt.Println(scores)

	scores = append(scores, 40) //append methods return a new slice, hence updating the previous slice

	fmt.Println(scores)

	scores = append(scores, 55)

	fmt.Println(scores)

}

func sliceRange() {

	// SLICE RANGE

	/**
	1. As name suggest is helps in taking a range of slice.
	2. Since it is a slice range we can append them.

	SYNTAX:
	slice[x:y]
	here x and y are positions
	NOTE:
	a. here x in included and y is excluded
	**/

	colors := []string{"red", "green", "black", "purple"}

	betweenRange := colors[1:3]

	fmt.Println(betweenRange)

	fromTillEnd := colors[3:] //To take the values till end just add ":"

	fmt.Println(fromTillEnd)

	staringToPosition := colors[:3] //If we add ":" in the start, it will conside all the values before position.

	fmt.Println(staringToPosition)

	//Appending Slice Range

	colors = append(colors, "sian")

	fmt.Println(colors)

}
