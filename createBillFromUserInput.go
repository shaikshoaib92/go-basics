package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, e := r.ReadString('\n') // here '\n means we capture the input after he press enter'

	return strings.TrimSpace(input), e
}

func createNewBill() bills {

	reader := bufio.NewReader(os.Stdin) //bufio is a package to read, and os.Stdin(standard input) is to caputure the use input.

	capturedName, _ := getInput("Create a new bill name: ", reader) // _ if for err which we return from that function

	newBill := myBill(capturedName)
	fmt.Printf("Create a bill - %v \n", newBill.name)

	return newBill

}
