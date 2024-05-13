package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["Go"] = "Golang"
	languages["RB"] = "Ruby"

	fmt.Println(languages)

	//Access a single key value pair
	fmt.Println("JS shorts for: ", languages["JS"])

	//Delete a key value pair from map
	delete(languages, "RB")
	fmt.Println(languages)

	//loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
