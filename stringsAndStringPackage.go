package main

import (
	"fmt"
	"strings"
)

func stringFunc() {

	var name string = "This is string"

	fmt.Println(name)
}

func formatedString() {

	text := "this is the example of formatted string and different format specifier"

	fmt.Println(text)

	var stringWithFormatSpecifiers string = "Hello my name is %q, my age is %v and my height is %f"

	name := "shoaib"
	age := 23
	height := 5.5

	fmt.Println(stringWithFormatSpecifiers, name, age, height)

}

func stringPackage() {

	//It gives the new string, it wont updat the orignal string

	greetings := "Hello there friends! Hello"

	fmt.Println(strings.Contains(greetings, "Hello!"))            //This returns boolean
	fmt.Println(strings.ReplaceAll(greetings, "Hello", "Challo")) //It replace all the similar words with the new word
}
