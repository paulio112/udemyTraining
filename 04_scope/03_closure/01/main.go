package  main

import "fmt"

func main()  {
  x := 42
  fmt.Println(x)
  {
     //fmt.Println(x)
     y := "This is within the scope"
     fmt.Println(y)
  }
  // Y is not available as it is out of scope.
//  fmt.println(y)
}
