// Go Output Functions

// Go has three functions to output text:

Print()
Println()
Printf()





// The Print() Function

1.)Example
Print the values of i and j:

package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i)
  fmt.Print(j)
}
Result:

HelloWorld


2.)Example
If we want to print the arguments in new lines, we need to use \n.

package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i, "\n")
  fmt.Print(j, "\n")
}
Result:

Hello
World




// The Printf() Function
// The Printf() function first formats its argument based on the given formatting verb and then prints them.

Here we will use two formatting verbs:

%v is used to print the value of the arguments
%T is used to print the type of the arguments

package main
import ("fmt")

func main() {
  var i string = "Hello"
  var j int = 15

  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("j has value: %v and type: %T", j, j)
}
Result:

i has value: Hello and type: string
j has value: 15 and type: int