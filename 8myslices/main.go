package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Mango", "Banana"}
	fmt.Printf("Type of fruitlist is  %T\n", fruitList)

	fruitList = append(fruitList, "Grapes", "Watermelon")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	// fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 210
	highScore[1] = 211
	highScore[2] = 212
	highScore[3] = 213

	highScore = append(highScore, 214, 215, 216)

	fmt.Println(highScore)

	// Sorting the slice using sort function
	sort.Sort(sort.Reverse(sort.IntSlice(highScore)))
	fmt.Println(highScore)

	//how to remove a value from slicong based on index
	var courses = []string{"reactjs", "js", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
