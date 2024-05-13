package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Grape"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Length of array is: ", len(fruitList))


	var vegList = [3]string {"potato", "tomato", "pea"}
	fmt.Println("Veg list is: ", vegList)
}