package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://web.whatsapp.com/"

func main() {
	fmt.Println("Handling web requests in golang")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type is: %T\n", response)

	defer response.Body.Close() //caller's responsibility to close the connection

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(databytes))
}
