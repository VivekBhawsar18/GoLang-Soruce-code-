// Anonymous function

// Syntax:

func(parameter_list)(return_type){
// code..

// Use return statement if return_type are given
// if return_type is not given, then do not 
// use return statement
return
}()

---------------------------------------------------------------------------------------------

// Go program to illustrate how
// to create an anonymous function

package main

import "fmt"

func main() {
	
	// Anonymous function
func(){

	fmt.Println("Welcome! to GeeksforGeeks")
}()
	
}

---------------------------------------------------------------------------------------------

// Go program to illustrate
// use of an anonymous function
package main

import "fmt"

func main() {
	
	// Assigning an anonymous
// function to a variable
value := func(){
	fmt.Println("Welcome! to GeeksforGeeks")
}
value()
	
}


---------------------------------------------------------------------------------------------

// Go program to pass arguments
// in the anonymous function
package main

import "fmt"

func main() {
	
	// Passing arguments in anonymous function
func(ele string){
	fmt.Println(ele)
}("GeeksforGeeks")
	
}


---------------------------------------------------------------------------------------------

// Go program to pass an anonymous
// function as an argument into
// other function
package main

import "fmt"


// Passing anonymous function
// as an argument
func GFG(i func(p, q string)string){
	fmt.Println(i ("Geeks", "for"))
	
}
	
func main() {
	value:= func(p, q string) string{
		return p + q + "Geeks"
	}
	GFG(value)
}

---------------------------------------------------------------------------------------------

// Go program to illustrate
// use of anonymous function
package main

import "fmt"

// Returning anonymous function
func GFG() func(i, j string) string{
	myf := func(i, j string)string{
		return i + j + "GeeksforGeeks"
	}
	return myf
}
	
func main() {
	value := GFG()
	fmt.Println(value("Welcome ", "to "))
}

---------------------------------------------------------------------------------------------

