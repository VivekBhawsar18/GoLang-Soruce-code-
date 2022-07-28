

// 1. Ints: This function is used to only sorts a slice of ints and it sorts the elements of the slice in increasing order.

// Syntax:

func Ints(slc []int)

---------------------------------------------------------------------------------------------

// Go program to illustrate how
// to sort the slice of ints
package main

import (
	"fmt"
	"sort"
)

// Main function
func main() {
	
	// Creating and initializing slices
	// Using shorthand declaration
	scl1 := []int{400, 600, 100, 300, 500, 200, 900}
	scl2 := []int{-23, 567, -34, 67, 0, 12, -5}
	
	// Displaying slices
	fmt.Println("Slices(Before):")
	fmt.Println("Slice 1: ", scl1)
	fmt.Println("Slice 2: ", scl2)
	
	// Sorting the slice of ints
	// Using Ints function
	sort.Ints (scl1)
	sort.Ints (scl2)
	
	// Displaying the result
	fmt.Println("\nSlices(After):")
	fmt.Println("Slice 1 : ", scl1)
	fmt.Println("Slice 2 : ",scl2)
}

---------------------------------------------------------------------------------------------

// 2. IntsAreSorted: This function is used to check whether the given slice of ints is in the sorted form(in increasing order) or not. 

// Syntax:

func IntsAreSorted(scl []int) bool

// Go program to illustrate how to check
// whether the given slice of ints is in
// sorted form or not
package main

import (
	"fmt"
	"sort"
)

// Main function
func main() {

	// Creating and initializing slices
	// Using shorthand declaration
	scl1 := []int{100, 200, 300, 400, 500, 600, 700}
	scl2 := []int{-23, 567, -34, 67, 0, 12, -5}

	// Displaying slices
	fmt.Println("Slices:")
	fmt.Println("Slice 1: ", scl1)
	fmt.Println("Slice 2: ", scl2)

	// Checking the slice is in sorted form or not
	// Using IntsAreSorted function
	res1 := sort.IntsAreSorted(scl1)
	res2 := sort.IntsAreSorted(scl2)

	// Displaying the result
	fmt.Println("\nResult:")
	fmt.Println("Is Slice 1 is sorted?: ", res1)
	fmt.Println("Is Slice 2 is sorted?: ", res2)
}

---------------------------------------------------------------------------------------------

