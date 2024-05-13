package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions in golang")
	greeter()
	greeterTwo()

	// func greeterTwo() {
	// 	fmt.Println("Another method")
	// }
	// greeterTwo()   will give error

	result := adder(3, 5)
	fmt.Println("Add is: ", result) // 8

	proRes := proAdder(2, 3, 5, 6, 7, 8)
	fmt.Println("Pro add is: ", proRes) // 31

	proRes2, myMessage := proAdder2(2, 3, 5, 6, 7, 8)
	fmt.Println("Pro add is: ", proRes2) // 31
	fmt.Println("Pro add is: ", myMessage) //hii

}

func greeterTwo() {
	fmt.Println("Another method")
}

func greeter() {
	fmt.Println("Namastey from golang.")
}

func adder(a int, b int) int {
	add := a + b
	return add
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func proAdder2(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "hii"
}