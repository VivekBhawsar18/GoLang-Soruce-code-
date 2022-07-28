// Declaring a Constant

package main

import ("fmt")

const PI = 3.14

func main() {
  fmt.Println(PI)
}





// Constant Types
// There are two types of constants:-

// 1.) Typed constants
// 2.) Untyped constants







// Typed Constants
// Typed constants are declared with a defined type:

package main
import ("fmt")

const A int = 1

func main() {
  fmt.Println(A)
}


// Untyped Constants
// Untyped constants are declared without a type:

package main
import ("fmt")

const A = 1

func main() {
  fmt.Println(A)









// Constants: Unchangeable and Read-only
// When a constant is declared, it is not possible to change the value later:

package main
import ("fmt")

func main() {
  const A = 1
  A = 2
  fmt.Println(A)
}
Result:

'''./prog.go:8:7: cannot assign to A'''







// Multiple Constants Declaration
// Multiple constants can be grouped together into a block for readability:

package main
import ("fmt")

const (
  A int = 1
  B = 3.14
  C = "Hi!"
)

func main() {
  fmt.Println(A)
  fmt.Println(B)
  fmt.Println(C)
}