package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, e := r.ReadString('\n') // here '\n means we capture the input after he press enter'

	return strings.TrimSpace(input), e
}

func prompOption(b bills) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Select an option ( a - add item, s - save bill, t - add tip): ", reader)

	//SWITCH CASE
	switch opt {
	case "a":
		name, _ := getInput("Itme name: ", reader)
		price, _ := getInput("Enter price: ", reader)
		p, e := strconv.Atoi(price) //By default the value of user input will be stiring, So we use the package called strconv.Atoi()
		if e != nil {
			fmt.Println("Please enter a valid number")
		}

		b.updateItems(name, p)
		fmt.Println("Updated Items with ", name, price)
		prompOption(b)
	case "t":
		tip, _ := getInput("Enter tip amount: ", reader)
		t, _ := strconv.ParseFloat(tip, 64) //By default the value of user input will be stiring, So we use the package called strconv.Atoi()

		b.updateTip(t)

		fmt.Println("Updated the tip: ", tip)
		prompOption(b)

	case "s":
		b.save()
		fmt.Println("Your file is saved", b.name)

	default:
		fmt.Println("You've pressed wrong key")
		prompOption(b)

	}

}

func createNewBill() bills {

	reader := bufio.NewReader(os.Stdin) //bufio is a package to read, and os.Stdin(standard input) is to caputure the use input.

	capturedName, _ := getInput("Create a new bill name: ", reader) // _ if for err which we return from that function

	newBill := myBill(capturedName)
	fmt.Printf("Create a bill - %v \n", newBill.name)

	return newBill

}
