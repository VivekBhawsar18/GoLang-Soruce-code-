// In Go, there are different types of variables, for example:

int		- stores integers (whole numbers), such as 123 or -123
float32 - stores floating point numbers, with decimals, such as 19.99 or -19.99
string  - stores text, such as "Hello World". String values are surrounded by double quotes
bool	- stores values with two states: true or false




// Difference Between var and :=

var	:=

1.)Can be used inside and outside of functions		
2.)Variable declaration and value assignment can be done separately

:=

1.)Can only be used inside functions
2.)Variable declaration and value assignment cannot be done separately (must be done in the same line)





// This example shows declaring variables outside of a function, with the var keyword:

package main
import ("fmt")

var a int
var b int = 2
var c 	  = 3

func main() {
  a = 1
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}

// Since := is used outside of a function, running the program results in an error.

package main
import ("fmt")

a := 1

func main() {
  fmt.Println(a)
}

Result:

'''./prog.go:5:1: syntax error: non-declaration statement outside function body'''