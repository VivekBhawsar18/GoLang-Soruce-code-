package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name 		string 	 `json:"coursename"`
	Price		int		
	Platform	string	 `json:"website"`
	Password	string	 `json:"-"`
	Tags		[]string `json:"tags,omitempty"`	
}

func main(){
	fmt.Println("Welcome to Json video")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	VivCourses := []course{

		{"ReactJS Bootcamp" , 299 , "Youtube.com" , "abc123" , []string{"web-dev " , "js"}},
		{"MERN Bootcamp" , 199 , "Youtube.com" , "def456" , []string{"Full-Stack" , "js"}},
		{"Angular Bootcamp" , 499 , "Youtube.com" , "ghi789" ,nil},

	}

	//package this data as JSON data

	finalJson , err := json.MarshalIndent(VivCourses , "" , "\t")

	if err!= nil{
		panic(err)
	}

	fmt.Printf("%s\n" , finalJson)

}

func DecodeJson(){

	jsonDataFromWeb := []byte(`
	{
		"coursename":"Go-Lang Bootcamp",
		"Price" : 299,
		"website":"Youtube.com",
		"tags":["web-dev" , "js"]

	}
	`)

	var  vivcourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid{
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb , &vivcourse)
		fmt.Printf("%#v\n" , vivcourse)
	}else{
		fmt.Println("JSON WAS NOT VALID")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal( jsonDataFromWeb , &myOnlineData )
		fmt.Printf("%#v\n" , myOnlineData)

	for k , v := range myOnlineData{
		fmt.Printf("Key is %v and value is %v and Type is: %T\n " , k , v , v)
	}

}

