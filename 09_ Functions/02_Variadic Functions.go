// Variadic Functions

// Syntax:  

function function_name(para1, para2...type)type {// code...}

---------------------------------------------------------------------------------------------

// Go program to illustrate the
// concept of variadic function

package main

import (
	"fmt"
	"strings"
)

// Variadic function to join strings
func joinstr(elements ...string) string {
	return strings.Join(elements, "-")
}

func main() {

	// zero argument
	fmt.Println(joinstr())

	// multiple arguments
	fmt.Println(joinstr("GEEK", "GFG"))
	fmt.Println(joinstr("Geeks", "for", "Geeks"))
	fmt.Println(joinstr("G", "E", "E", "k", "S"))

}

---------------------------------------------------------------------------------------------

// Go program to illustrate the
// concept of variadic function
package main

import (
	"fmt"
	"strings"
)

// Variadic function to join strings
func joinstr(elements ...string) string {
	return strings.Join(elements, "-")
}

func main() {
	// pass a slice in variadic function
	elements := []string{"geeks", "FOR", "geeks"}
	fmt.Println(joinstr(elements...))
}

---------------------------------------------------------------------------------------------