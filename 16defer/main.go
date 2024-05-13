package main

import "fmt"

func main() {

	// defer statement allows you to delay the execution of a function until the surrounding function exits.
	// When there are multiple defer functions then the output will come in reverse order.
	defer fmt.Println("Three")
	defer fmt.Println("Two")
	defer fmt.Println("One")
	fmt.Println("Defer in golang")
	myDefer()
	
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}