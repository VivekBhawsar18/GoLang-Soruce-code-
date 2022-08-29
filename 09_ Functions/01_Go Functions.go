// Go Functions

// Syntax

func FunctionName() {
  // code to be executed
}

---------------------------------------------------------------------------------------------

// Call a Function

Functions are not executed immediately. They are "saved for later use", and will be executed when they are called.

// Example

package main
import ("fmt")

func myMessage() {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage() // call the function
}

---------------------------------------------------------------------------------------------

// A function can be called multiple times.

// Example

package main
import ("fmt")

func myMessage() {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage()
  myMessage()
  myMessage()
}

---------------------------------------------------------------------------------------------

// Function With Parameter Example


// Example

package main
import ("fmt")

func familyName(fname string) {
  fmt.Println("Hello", fname, "Refsnes")
}

func main() {
  familyName("Liam")
  familyName("Jenny")
  familyName("Anja")
}

---------------------------------------------------------------------------------------------

// Multiple Parameters

// Example

package main
import ("fmt")

func familyName(fname string, age int) {
  fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func main() {
  familyName("Liam", 3)
  familyName("Jenny", 14)
  familyName("Anja", 30)
}

---------------------------------------------------------------------------------------------

// Go Function Returns

// Syntax

func FunctionName(param1 type, param2 type) type {
  // code to be executed
  return output

---------------------------------------------------------------------------------------------

// Function Return Example

// Example

Here, myFunction() receives two integers (x and y) and returns their addition (x + y) as integer (int):

package main
import ("fmt")

func myFunction(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println(myFunction(1, 2))
}

---------------------------------------------------------------------------------------------

// Named Return Values

In Go, you can name the return values of a function.

// Example

package main
import ("fmt")

func myFunction(x int, y int) (result int) {
  result = x + y
  return
}

func main() {
  fmt.Println(myFunction(1, 2))
}

---------------------------------------------------------------------------------------------

// Example

package main
import ("fmt")

func myFunction(x int, y int) (result int) {
  result = x + y
  return result
}

func main() {
  fmt.Println(myFunction(1, 2))
}

---------------------------------------------------------------------------------------------

// Store the Return Value in a Variable

You can also store the return value in a variable, like this:

// Example
Here, we store the return value in a variable called total:

package main
import ("fmt")

func myFunction(x int, y int) (result int) {
  result = x + y
  return
}

func main() {
  total := myFunction(1, 2)
  fmt.Println(total)
}

---------------------------------------------------------------------------------------------

// Multiple Return Values

Go functions can also return multiple values.

// Example

Here, myFunction() returns one integer (result) and one string (txt1):

package main
import ("fmt")

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
  fmt.Println(myFunction(5, "Hello"))
}

---------------------------------------------------------------------------------------------

// Example

Here, we store the two return values into two variables (a and b):

package main
import ("fmt")

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
  a, b := myFunction(5, "Hello")
  fmt.Println(a, b)
}

---------------------------------------------------------------------------------------------

// Example

Here, we want to omit the first returned value (result - which is stored in variable a):

package main
import ("fmt")

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
   _, b := myFunction(5, "Hello")
  fmt.Println(b)
}

---------------------------------------------------------------------------------------------

// Example

Here, we want to omit the second returned value (txt1 - which is stored in variable b):

package main
import ("fmt")

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
   a, _ := myFunction(5, "Hello")
  fmt.Println(a)
}

---------------------------------------------------------------------------------------------

// Go Recursion Functions

// Example

package main
import ("fmt")

func testcount(x int) int {
  if x == 11 {
    return 0
  }
  fmt.Println(x)
  return testcount(x + 1)
}

func main(){
  testcount(1)
}

---------------------------------------------------------------------------------------------

// Example

package main
import ("fmt")

func factorial_recursion(x float64) (y float64) {
  if x > 0 {
     y = x * factorial_recursion(x-1)
  } else {
     y = 1
  }
  return
}

func main() {
  fmt.Println(factorial_recursion(4))
}

---------------------------------------------------------------------------------------------

