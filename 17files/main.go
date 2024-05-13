package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in golang")
	content := "Written by Aryan"

	file, err := os.Create("./newFile.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile("./newFile.txt")
}

func readFile(filename string) {

	databyte, err := os.ReadFile(filename)
	checkNilError((err))  // We can use this
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Text  inside file is: ", string(databyte))
}

func checkNilError(err error){
	if err != nil {
		panic(err)
	}
}