package main

import (
	"fmt"
	"net/url"
)

func main(){

	myurl := "http://localhost:3000"

	fmt.Println("Welcome to handling Urls in golang")
	fmt.Println(myurl)

	// parsing
	result , _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are : %T\n" , qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams{
		fmt.Println("Param is : ", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "http",
		Host: "localhost:3000",
		Path:"",
		RawPath:"",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}