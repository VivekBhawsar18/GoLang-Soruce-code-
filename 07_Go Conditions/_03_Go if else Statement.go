// Go if else Statement

// Syntax

if condition {
  // code to be executed if condition is true
} else {
  // code to be executed if condition is false
}

---------------------------------------------------------------------------------------------

Using The if else Statement

package main
import ("fmt")

func main() {
  time := 20
  if (time < 18) {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }
}

---------------------------------------------------------------------------------------------

// Example

package main
import ("fmt")

func main() {
  temperature := 14
  if (temperature > 15) {
    fmt.Println("It is warm out there")
  } else {
    fmt.Println("It is cold out there")
  }
}

---------------------------------------------------------------------------------------------

// Go program to illustrate the
// use of if...else statement
package main

import "fmt"

func main() {
	
// taking a local variable
var v int = 1200

// using if statement for
// checking the condition
if v < 1000 {
		
	// print the following if
	// condition evaluates to true
	fmt.Printf("v is less than 1000\n")
		
} else {
		
	// print the following if
	// condition evaluates to true
	fmt.Printf("v is greater than 1000\n")
}
	
}


---------------------------------------------------------------------------------------------

