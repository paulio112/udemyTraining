package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Println("Enter how far you swam")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println("This is meters", meters, "This is yards", yards)

}
