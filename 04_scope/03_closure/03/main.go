package main

import "fmt"

func main() {
	var x = 1
	// when a function is embedded within a function it has to be anon, eg unamed but still expect a type
	// thi is called a func expression when a func is assigned to a variable
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

}
