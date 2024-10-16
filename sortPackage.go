package main

import (
	"fmt"
	"sort"
)

func sortPackage() {

	//sort package alters the orignal  array, it wort return the new one
	num := []int{10, 20, 60, 9, 3, 6, 88, 568, 754, 8555}

	idx := sort.SearchInts(num, 855855)
	fmt.Println(idx)

	// names := []string{"shaik", "black", "house", "car", "pulser"}

	// sort.Strings(names) //It sorts them in alphabetical order
	// res := sort.SearchStrings(names, "car")

	// fmt.Println(names)
	// fmt.Println(res)

}
