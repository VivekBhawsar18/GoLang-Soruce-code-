

// In Go, there are several ways to create a slice:

// Syntax:  

[]T

or 
[]T{}

or 
[]T{value1, value2, value3, ...value n}

or
slice_name := []datatype{values}

// Here, T is the type of the elements. For example: 

var my_slice[]int 

---------------------------------------------------------------------------------------------


// Create a slice from an array
Using the make() function

---------------------------------------------------------------------------------------------

// A common way of declaring a slice is like this:

myslice := []int	// Empty slice of 0 length and 0 capacity.

myslice := []int{1,2,3}		//slice of integers of length 3 and also the capacity of 3. 

---------------------------------------------------------------------------------------------

// In Go, there are two functions that can be used to return the length and capacity of a slice:

1.) len() function - returns the length of the slice (the number of elements in the slice)

2.) cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)


// Example
This example shows how to create slices using the []datatype{values} format:

package main
import ("fmt")

func main() {
  myslice1 := []int{}
  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1))
  fmt.Println(myslice1)

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
  fmt.Println(myslice2)
}

---------------------------------------------------------------------------------------------

// Create a Slice From an Array
You can create a slice by slicing an array:


// Syntax
var myarray = [length]datatype{values} 	// An array
myslice 	:= myarray[start:end] 		// A slice made from the array


// Example
This example shows how to create a slice from an array:

package main
import ("fmt")

func main() {
  arr1 := [6]int{10, 11, 12, 13, 14,15}
  myslice := arr1[2:4]

  fmt.Printf("myslice = %v\n", myslice)
  fmt.Printf("length = %d\n", len(myslice))
  fmt.Printf("capacity = %d\n", cap(myslice))
}

---------------------------------------------------------------------------------------------

Example: 

// Golang program to illustrate
// the working of the slice components
package main
 
import "fmt"
 
func main() {
 
    // Creating an array
    arr := [7]string{"This", "is", "the", "tutorial", "of", "Go", "language"}
 
    // Display array
    fmt.Println("Array:", arr)
 
    // Creating a slice
    myslice := arr[1:6]
 
    // Display slice
    fmt.Println("Slice:", myslice)
 
    // Display length of the slice
    fmt.Printf("Length of the slice: %d", len(myslice))
 
    // Display the capacity of the slice
    fmt.Printf("\nCapacity of the slice: %d", cap(myslice))
}

---------------------------------------------------------------------------------------------

slice is the reference of the array!!

Example:

// Golang program to illustrate how to
// create slices from the array
package main
 
import "fmt"
 
func main() {
 
    // Creating an array
    arr := [4]string{"Geeks", "for", "Geeks", "GFG"}
 
    // Creating slices from the given array
    var my_slice_1 = arr[1:2]
    my_slice_2 := arr[0:]
    my_slice_3 := arr[:2]
    my_slice_4 := arr[:]
 
    // Display the result
    fmt.Println("My Array: ", arr)
    fmt.Println("My Slice 1: ", my_slice_1)
    fmt.Println("My Slice 2: ", my_slice_2)
    fmt.Println("My Slice 3: ", my_slice_3)
    fmt.Println("My Slice 4: ", my_slice_4)
}

---------------------------------------------------------------------------------------------

// Using make() function:

Syntax: 
func make([]T, len, cap) []T

Example:

// Go program to illustrate how to create slices
// Using make function
package main
 
import "fmt"
 
func main() {
 
    // Creating an array of size 7
    // and slice this array  till 4
    // and return the reference of the slice
    // Using make function
    var my_slice_1 = make([]int, 4, 7)
    fmt.Printf("Slice 1 = %v, \nlength = %d, \ncapacity = %d\n", my_slice_1, len(my_slice_1), cap(my_slice_1))
 
    // Creating another array of size 7
    // and return the reference of the slice
    // Using make function
    var my_slice_2 = make([]int, 7)
    fmt.Printf("Slice 2 = %v, \nlength = %d, \ncapacity = %d\n", my_slice_2, len(my_slice_2), cap(my_slice_2))
     
}

---------------------------------------------------------------------------------------------

// Golang program to illustrate
// how to modify the slice
package main

import "fmt"

func main() {

	// Creating a zero value slice
	arr := [6]int{55, 66, 77, 88, 99, 22}
	slc := arr[0:4]

	// Before modifying

	fmt.Println("Original_Array: ", arr)
	fmt.Println("Original_Slice: ", slc)

	// After modification
	slc[0] = 100
	slc[1] = 1000
	slc[2] = 1000

	fmt.Println("\nNew_Array: ", arr)
	fmt.Println("New_Slice: ", slc)
}

---------------------------------------------------------------------------------------------

// Go program to illustrate the multi-dimensional slice
package main

import "fmt"

func main() {

	// Creating multi-dimensional slice
	s1 := [][]int{{12, 34},
		{56, 47},
		{29, 40},
		{46, 78},
	}

	// Accessing multi-dimensional slice
	fmt.Println("Slice 1 : ", s1)

	// Creating multi-dimensional slice
	s2 := [][]string{
			{"Geeks", "for"},
			{"Geeks", "GFG"},
			{"gfg", "geek"},
	}

	// Accessing multi-dimensional slice
	fmt.Println("Slice 2 : ", s2)

}

---------------------------------------------------------------------------------------------

// Go program to illustrate how to sort
// the elements present in the slice
package main

import (
	"fmt"
	"sort"
)

func main() {

	// Creating Slice
	slc1 := []string{"Python", "Java", "C#", "Go", "Ruby"}
	slc2 := []int{45, 67, 23, 90, 33, 21, 56, 78, 89}

	fmt.Println("Before sorting:")
	fmt.Println("Slice 1: ", slc1)
	fmt.Println("Slice 2: ", slc2)

	// Performing sort operation on the
	// slice using sort function
	sort.Strings(slc1)
	sort.Ints(slc2)

	fmt.Println("\nAfter sorting:")
	fmt.Println("Slice 1: ", slc1)
	fmt.Println("Slice 2: ", slc2)

}
