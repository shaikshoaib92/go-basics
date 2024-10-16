package main

import "fmt"

func loops() {
	//similart to while loop

	x := 0

	for x < 5 {
		fmt.Println(x)
		x++
	}

	// similar to for loop

	nums := []int{1, 2, 3, 5, 6}

	for i := 0; i < len(nums); i++ {
		fmt.Printf("The value at position %v is %v \n", i, nums[i])
	}

	//similar to for in loop

	names := []string{"hello", "challo", "khelo", "khudo"}

	for idx, name := range names {
		fmt.Printf("The value at position %v is %v \n", idx, name)
	}

	//If we not want any of idx or name just replace it with "_"

	for _, name := range names {
		fmt.Printf("The values are  %v \n", name)
	}

}
