package main

import "fmt"

func main() {

	a := 25

	fmt.Println(a)
	// Memory address of a
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 10 //assign value of 10 to what b points to.
	fmt.Println(a)

}
