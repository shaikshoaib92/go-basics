package main

import "fmt"

func conditions() {
	//If else is similar to any other lang.
	x := 1

	if x != 10 {
		fmt.Println("lulu")
	} else if x == 10 {
		fmt.Println("world")
	} else {
		fmt.Println("cat")
	}

	// CONTINIUE
	// it will go back to loop and continue

	// BREAK
	// it will stop the execution of loop

	names := []string{"hello", "challo", "khelo", "khudo", "khelo", "khudo"}

	for idx, name := range names {
		if idx == 1 {
			fmt.Println(name)
			continue
		} else if idx == 3 {
			fmt.Print(name)
			break
		}

		fmt.Println(name, idx)
	}

}
