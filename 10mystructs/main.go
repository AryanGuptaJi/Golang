package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	//no inheritance in golang, no super ou parent as well

	aryan := User{"Aryan", "aryang23936@gmail.com", true, 20}
	fmt.Println(aryan)
	//	{Aryan aryang23936@gmail.com true 20}

	fmt.Printf("Aryan details are: %v\n",aryan)
	//  Aryan details are: {Aryan aryang23936@gmail.com true 20}  

	fmt.Printf("Aryan details are: %+v\n",aryan)
	//  Aryan details are: {Name:Aryan Email:aryang23936@gmail.com Status:true Age:20}

	fmt.Printf("Name is %v, Email is %v", aryan.Name, aryan.Email)
	// Name is Aryan, Email is aryang23936@gmail.com
}

type User struct {
	Name   string //public field
	Email  string
	Status bool
	Age    int
}
