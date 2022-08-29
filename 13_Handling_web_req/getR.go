package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// const url string = "https://lco.dev"
const url string = "http://localhost:3000" 

func main ()  {

		responce , err := http.Get(url)

		if err!= nil{
			panic(err)
		}

		fmt.Printf("Responce is of type %T\n" , responce)

		responce.Body.Close() //Callers responsibilty to close the  req.

// -------------------------------------------------------------------------------------
		databytes , err := ioutil.ReadAll(responce.Body)

		if err!= nil{
			panic(err)
		}

		content := string(databytes)

		fmt.Println(content)
}

