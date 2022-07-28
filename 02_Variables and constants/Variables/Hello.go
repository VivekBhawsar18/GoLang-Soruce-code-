package main

import "fmt"

func main(){


	// Variable Declaration With Initial Value
	
	var num1 = 2
	var num2 = 3

	var result = num1 + num2 

	// OR

	var num1 int = 2
	var num2 int = 4

	var result = num1 + num2 

	// Value Assignment After Declaration

	var num1 int 
	var num2 int 

	num1 = 2
	num2 = 8

	var result = num1 + num2 

	// OR

	var num1 , num2 int 
	num1 , num2 = 2 ,7

	var result = num1 + num2 

	//OR

	result:= 9  // Similar as : var result = 9 

	// Go Variable Declaration in a Block

	   var (
     	a int
     	b int = 1
    	c string = "hello"
     	)
		
  		fmt.Println(a)
  		fmt.Println(b)
  		fmt.Println(c)

	//Constant
	 
	const result = 33
	result = 30 // This will not change the value of result 

	fmt.Print(result)
}