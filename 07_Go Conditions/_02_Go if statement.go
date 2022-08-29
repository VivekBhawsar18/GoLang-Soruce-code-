// Go if statement

Syntax

if condition {
  // code to be executed if condition is true
}

---------------------------------------------------------------------------------------------

// Example

package main
import ("fmt")

func main() {
  if 20 > 18 {
    fmt.Println("20 is greater than 18")
  }
}

---------------------------------------------------------------------------------------------

// Go program to illustrate the
// use of if statement
package main

import "fmt"

func main() {
	
// taking a local variable
var v int = 700

// using if statement for
// checking the condition
if v < 1000 {
		
	// print the following if
	// condition evaluates to true
	fmt.Printf("v is less than 1000\n")
}
	
fmt.Printf("Value of v is : %d\n", v)
	
}

---------------------------------------------------------------------------------------------

// Example

package main
import ("fmt")

func main() {
  x:= 20
  y:= 18
  if x > y {
    fmt.Println("x is greater than y")
  }
}
