package main

import "fmt"

func arrays() {
	//ARRAYS

	/**

	NOTE:
	1. We cannot add anything to array when the size is declared, arrays are fixed.
	**/

	var num [3]int = [3]int{1, 5, 9}
	var extraNum = [5]int{1, 2, 3, 6, 9}
	names := [2]string{"shaik", "shoaib"}

	fmt.Println(num, len(num))
	fmt.Println(extraNum, len(extraNum))
	fmt.Println(names, len(names))

	names[1] = "shaik"
	fmt.Println(names) //Here we are replacing the 1st idx with new value
}
