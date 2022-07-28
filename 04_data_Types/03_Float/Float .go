// Go Float Data Types

// Reffer image no 1 




// The float32 Keyword
// Example
// This example shows how to declare some variables of type float32:

package main
import ("fmt")

func main() {
  var x float32 = 123.78
  var y float32 = 3.4e+38
  fmt.Printf("Type: %T, value: %v\n", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}






// The float64 Keyword
// The float64 data type can store a larger set of numbers than float32.

// Example
// This example shows how to declare a variable of type float64:

package main
import ("fmt")

func main() {
  var x float64 = 1.7e+308
  fmt.Printf("Type: %T, value: %v", x, x)
}