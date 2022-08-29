package main

import (
	// "context"
	"fmt"
	"strings"

	// "image/color/palette"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main(){
	fmt.Println("Welcome to Get req : ") 
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFromRequest()

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

func PerformPostJsonRequest(){

	const myurl = "http://localhost:3000/get"

	requestBody := strings.NewReader(`
		{
			"coursename" :"Learn Go"
			"price" : "0"
			"platform":"Assimilate Technologies"
		}
	`)

	responce , err := http.Post(myurl , "application/json" , requestBody)

	if err != nil{
		panic(err)
	}

	defer responce.Body.Close()

	content , _ := ioutil.ReadAll(responce.Body)

	fmt.Println(string(content))

}

func PerformPostFromRequest(){
	const myurl = "http://localhost:3000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname" , "Vivek")
	data.Add("lastname" , "Bhawsar")
	data.Add("email" , "vivek@go.dev")

	responce , err := http.PostForm(myurl , data)

	if err!= nil{
		panic(err)
	}

	defer responce.Body.Close()

	content , _ := ioutil.ReadAll(responce.Body)
	fmt.Println(string(content))

}
