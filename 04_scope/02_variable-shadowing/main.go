package main

import "fmt"

func max(x int) int {
	return 20 * x
}

func main() {
	// example of assigning a variable to the func
	max := max(7)
	fmt.Println(max) // Result not the function anymore.

}
