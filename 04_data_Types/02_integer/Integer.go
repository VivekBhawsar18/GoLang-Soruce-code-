// Go Integer Data Types


// Integer data types are used to store a whole number without decimals, like 35, -50, or 1345000.

// The integer data type has two categories:

// 1.)Signed integers - can store both positive and negative values
// 2.)Unsigned integers - can only store non-negative values




// 1.)Signed Integers
// Signed integers, declared with one of the int keywords, can store both positive and negative values:

// Example
package main
import ("fmt")

func main() {
  var x int = 500
  var y int = -4500
  fmt.Printf("Type: %T, value: %v", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}

// Reffer image no 1 and 2




// 2.)Unsigned Integers
// Unsigned integers, declared with one of the uint keywords, can only store non-negative values:

// Example
package main
import ("fmt")

func main() {
  var x uint = 500
  var y uint = 4500
  fmt.Printf("Type: %T, value: %v", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}

// Reffer image no 2 and 3 and 4