package main

import "fmt"

func main() {
	fmt.Println("If Else in golang")

	//1.
	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10{
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)  //Watch out


	//2.
	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("umber is odd")
	}


	//3.
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	//4.
	// if err != nil {

	// }
}
