package main

import (
	"fmt"
	"strings"
)

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mult(a int, b int) int {
	return a * b
}

func calculator(a int, b int, f func(int, int) int) {
	rs := f(a, b)
	fmt.Println("Result is :", rs)
} // We can pass a fun as arg, but if that func has a return type specify that as well

func kindMain() {
	calculator(55, 99, mult)
}

func getInits(n string) (string, string) {
	up := strings.ToUpper(n)

	segrigated := strings.Split(up, " ")

	var initials []string
	for _, v := range segrigated {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	} else {
		return initials[0], "_"
	}

}
