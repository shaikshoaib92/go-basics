package main

import "fmt"

func numbers() {

	//INT

	var one int = 1
	var two = 2
	three := 3

	fmt.Println(one, two, three)

	//FLOAT

	var flotone float32 = 555.355411288788884
	var floattwo float64 = 5555555555555555555555555555555.333333333333333333333333333

	fmt.Println(flotone)
	fmt.Println(floattwo)
}

func bitsAndMemory() {

	//BITS AND MEMORY

	var eightbit int8 = 25
	var sixteenbit int16 = 5416
	var unsignedint uint = 52 //it wont accept neg values
	var unsignedintwithmemory uint32 = 652138421

	fmt.Println(eightbit)
	fmt.Println(sixteenbit)
	fmt.Println(unsignedint)
	fmt.Println(unsignedintwithmemory)

}
