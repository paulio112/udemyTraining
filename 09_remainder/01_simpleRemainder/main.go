package main

import "fmt"

func main() {

	// The percent sign acts as a what is the remainder operator
	x := 12 % 3
	fmt.Println(x)

	for i := 1; i < 100; i++ {

		if i%2 == 1 {
			fmt.Println("Fizz")
		} else {
			fmt.Println("Buzz")
		}
	}

}
