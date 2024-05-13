package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on Pointers")

	// var ptr *int 
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
//    referencing means & gets the address in memory where myNumber resides, so we can use it as a pointer
	var ptr = &myNumber
	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)

}
