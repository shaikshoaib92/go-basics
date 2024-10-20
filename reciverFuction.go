package main

import (
	"fmt"
	"os"
)

type bills struct {
	name  string
	items map[string]int
	tip   float64
}

/*
Receiver funcations are the function which are associated with STRUCT.
Genrally wheh we declare a function it will be availabe whin the folder from everywhere.
But recevier funcation when associated with struct are accessable from the struct object.
*/

func myBill(n string) bills {
	x := bills{
		name:  n,
		items: map[string]int{},
		tip:   0,
	}

	return x
}

/*
To associate a receiver funcation with the bill object, we must recive the bill object in the function.
*/

//Now the below function is only accessable with the bill object.
/*
Eg: res := myBill("shoaib")
	res.format() //Since the return type of myBill is bills (struct)
*/
func (b bills) format() string {
	//This function creates a copy of bills.
	//Not only this every argument we pass to any function creates a copy of that.

	fs := fmt.Sprintf("The Total Bill BreakDown: \n")

	var total int = 0

	for k, v := range b.items {

		fs += fmt.Sprintf("%-25v ...Rs.%v \n", k+":", v) //Here we use -25 padding to add 25characher space between them
		total += v
	}

	fs += fmt.Sprintf("%-25v ...%v \n", "TIP:", b.tip)
	fs += fmt.Sprintf("%-25v ...%v", "Total:", total+int(b.tip)) //Here we are doing typecasting

	return fs

}

//RECIVER FUNCTION WITH POINTER.

/*
	In the above example we have passing the copy of bills, which when updated will
	not reflect.
	To solve this we can pass the pointer of bills.

	NOTE: We dont need to derefrence the pointers when using with STRUCT
*/

func (b *bills) updateTip(t float64) {
	b.tip = t
}

func (b *bills) updateItems(item string, price int) {
	b.items[item] = price
}

func (b *bills) save() {

	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644) //This will save the data,
	/*
		It takes 3 args.
		1. location.
		2. data.
		3. permission.
	*/

	if err != nil {
		panic(err)
	}

}

func dummyMain() {
	newBill := myBill("shaik")

	newBill.updateItems("Lassi", 55)
	newBill.updateTip(100)
	fmt.Println(newBill.format())
}
