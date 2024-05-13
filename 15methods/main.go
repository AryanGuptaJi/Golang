package main

import "fmt"

func main() {
	fmt.Println("Methods in golang")

	aryan := User{"Aryan", "aryang23936@gmail.com", true, 20}
	fmt.Println(aryan)
	fmt.Printf("Name is %v\n", aryan.Name)

	aryan.GetStatus()
	aryan.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "aryangupta140404@gmail.com"
	fmt.Println("New email is: ", u.Email)
}