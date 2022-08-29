package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main ()  {
	fmt.Println("Welcome to files in goLang ..")
	content := "This needs to go in a file"

	file , err := os.Create("./myGofile.txt")

	// if err != nil {
	// 	panic(err)
	// }

	checkNilErr(err)

	length , err := io.writeString(file , content)
	fmt.Println("length is : " , length)
	// file.Close()
	defer file.Close()
	readFile("./myGofile.txt")
}

func readFile(filename string){
	databyte , err := ioutil.ReadFile(filename)

	// if err != nil {
	// 	panic(err)
	// }  

	checkNilErr(err)

	// fmt.Println("Text data inside the file is : ", databyte)

	fmt.Println("Text data inside the file is : ", string (databyte))
}


func checkNilErr(err error){
	
	if err != nil {
		panic(err)
	}
}