package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=nuebfefbiioenfiw23398"

func main() {
	fmt.Println("Handling urls in golang")

	//paring the urls
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qParams := result.Query()

	fmt.Printf("Type of query params are: %T\n", qParams)

	fmt.Println(qParams["coursename"])

	for _, val := range qParams {
		fmt.Println("Param is: ", val)
	}



	// When construct the url from chunks
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/css",
		RawPath: "user=Aryan",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
