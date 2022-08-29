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
	EncodeJson()
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


