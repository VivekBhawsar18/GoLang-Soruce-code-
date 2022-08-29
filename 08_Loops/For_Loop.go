
// Syntax

for statement1; statement2; statement3 {
   // code to be executed for each iteration
}

---------------------------------------------------------------------------------------------

// for Loop Examples
Example 1
This example will print the numbers from 0 to 4:  

package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }
}

---------------------------------------------------------------------------------------------

// Example 2
This example counts to 100 by tens: 

package main
import ("fmt")

func main() {
  for i:=0; i <= 100; i+=10 {
    fmt.Println(i)
  }

---------------------------------------------------------------------------------------------

// The continue Statement

The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.

// Example
This example skips the value of 3:

package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      continue
    }
   fmt.Println(i)
  }
}

---------------------------------------------------------------------------------------------

// The break Statement
The break statement is used to break/terminate the loop execution.

// Example
This example breaks out of the loop when i is equal to 3:

package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      break
    }
   fmt.Println(i)
  }
}

---------------------------------------------------------------------------------------------

// Nested Loops
It is possible to place a loop inside another loop.

Here, the "inner loop" will be executed one time for each iteration of the "outer loop":

// Example
package main
import ("fmt")

func main() {
  adj := [2]string{"big", "tasty"}
  fruits := [3]string{"apple", "orange", "banana"}
  for i:=0; i < len(adj); i++ {
    for j:=0; j < len(fruits); j++ {
      fmt.Println(adj[i],fruits[j])
    }
  }
}

---------------------------------------------------------------------------------------------

// The Range Keyword

The range keyword is used to more easily iterate over an array, slice or map. It returns both the index and the value.

The range keyword is used like this:

---------------------------------------------------------------------------------------------

// Syntax

for index, value := array|slice|map {
   // code to be executed for each iteration
}

---------------------------------------------------------------------------------------------

// Example

This example uses range to iterate over an array and print both the indexes and the values at each (idx stores the index, val stores the value):

package main
import ("fmt")

func main() {
  fruits := [3]string{"apple", "orange", "banana"}
  for idx, val := range fruits {
     fmt.Printf("%v\t%v\n", idx, val)
  }
}

---------------------------------------------------------------------------------------------

// Example
Here, we want to omit the indexes (idx stores the index, val stores the value):

package main
import ("fmt")

func main() {
  fruits := [3]string{"apple", "orange", "banana"}
  for _, val := range fruits {
     fmt.Printf("%v\n", val)
  }
}

---------------------------------------------------------------------------------------------

// Example

Here, we want to omit the values (idx stores the index, val stores the value):

package main
import ("fmt")

func main() {
  fruits := [3]string{"apple", "orange", "banana"}

  for idx, _ := range fruits {
     fmt.Printf("%v\n", idx)
  }
}

---------------------------------------------------------------------------------------------



