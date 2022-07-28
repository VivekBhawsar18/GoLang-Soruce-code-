// Example 1
// This example declares two arrays (arr1 and arr2) with defined lengths:

package main
import ("fmt")

func main() {
  var arr1 = [3]int{1,2,3}
  arr2 := [5]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)
}

---------------------------------------------------------------------------------------------

// Example 2
// This example declares two arrays (arr1 and arr2) with inferred lengths:

package main
import ("fmt")

func main() {
  var arr1 = [...]int{1,2,3}
  arr2 := [...]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)
}

---------------------------------------------------------------------------------------------

// Example 3
// This example declares an array of strings:

package main
import ("fmt")

func main() {
  var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  fmt.Print(cars)
}


---------------------------------------------------------------------------------------------


// Access Elements of an Array

// Example
// This example shows how to access the first and third elements in the prices array:

package main
import ("fmt")

func main() {
  prices := [3]int{10,20,30}

  fmt.Println(prices[0])
  fmt.Println(prices[2])
}

---------------------------------------------------------------------------------------------

// Change Elements of an Array

// Example
// This example shows how to change the value of the third element in the prices array: 

package main
import ("fmt")

func main() {
  prices := [3]int{10,20,30}

  prices[2] = 50
  fmt.Println(prices)
}


---------------------------------------------------------------------------------------------

// Array Initialization


// Example 1
package main
import ("fmt")

func main() {
  arr1 := [5]int{} //not initialized
  arr2 := [5]int{1,2} //partially initialized
  arr3 := [5]int{1,2,3,4,5} //fully initialized

  fmt.Println(arr1)
  fmt.Println(arr2)
  fmt.Println(arr3)
}

---------------------------------------------------------------------------------------------

// Example 2
This example initializes only the second and third elements of the array: 

package main
import ("fmt")

func main() {
  arr1 := [5]int{1:10,2:40}

  fmt.Println(arr1)
}

---------------------------------------------------------------------------------------------

// Find the Length of an Array

// Example 1
package main
import ("fmt")

func main() {
  arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  arr2 := [...]int{1,2,3,4,5,6}

  fmt.Println(len(arr1))
  fmt.Println(len(arr2))
}

---------------------------------------------------------------------------------------------

// How to Copy an Array into Another Array in Golang?


Syntax:

// creating a copy of an array by value
arr := arr1

// Creating a copy of an array by reference
arr := &arr1

---------------------------------------------------------------------------------------------

Example 1:

// Go program to illustrate how
// to copy an array by value
package main
  
import "fmt"
  
func main() {
  
    // Creating and initializing an array
    // Using shorthand declaration
    my_arr1 := [5]string{"Scala", "Perl", "Java", "Python", "Go"}
  
    // Copying the array into new variable
    // Here, the elements are passed by value
    my_arr2 := my_arr1
  
    fmt.Println("Array_1: ", my_arr1)
    fmt.Println("Array_2:", my_arr2)
  
    my_arr1[0] = "C++"
  
    // Here, when we copy an array
    // into another array by value
    // then the changes made in original
    // array do not reflect in the
    // copy of that array
    fmt.Println("\nArray_1: ", my_arr1)
    fmt.Println("Array_2: ", my_arr2)
}

---------------------------------------------------------------------------------------------

Example 2:

// Go program to illustrate how to
// copy an array by reference
package main
  
import "fmt"
  
func main() {
  
    // Creating and initializing an array
    // Using shorthand declaration
    my_arr1 := [6]int{12, 456, 67, 65, 34, 34}
  
    // Copying the array into new variable
    // Here, the elements are passed by reference
    my_arr2 := &my_arr1
  
    fmt.Println("Array_1: ", my_arr1)
    fmt.Println("Array_2:", *my_arr2)
  
    my_arr1[5] = 1000000
  
    // Here, when we copy an array 
    // into another array by reference
    // then the changes made in original 
    // array will reflect in the
    // copy of that array
    fmt.Println("\nArray_1: ", my_arr1)
    fmt.Println("Array_2:", *my_arr2)
}