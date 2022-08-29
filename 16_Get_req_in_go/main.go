package main

import (
	// "context"
	"fmt"
	"strings"

	// "image/color/palette"
	"io/ioutil"
	"net/http"
)

func main(){
	fmt.Println("Welcome to Get req : ") 
	PerformGetRequest()
}

func PerformGetRequest(){
	const myurl = "http://localhost:3000/get"

	responce , err := http.Get(myurl)
	if  err!= nil {
		panic(err)
	}

	defer responce.Body.Close()

	fmt.Println("Status code :" , responce.StatusCode)
	fmt.Println("Content Length is : :" , responce.ContentLength)

	var responceString strings.Builder
	content , _ := ioutil.ReadAll(responce.Body)
	byteCount , _ := responceString.Write(content)

	fmt.Println("ByteCount is :" , byteCount)
	fmt.Println("responceString is  :" , responceString.String())

	
	// content , _ := ioutil.ReadAll(responce.Body)
	// fmt.Println(string(content))
}