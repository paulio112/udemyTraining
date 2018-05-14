package main

import (
	"fmt"
	// hit this path to see what variables are available
	"github.com/GoesToEleven/GolangTraining/04_scope/01_package-scope/02_visibility/vis"
)

func main() {
	// The import on the below works because it has a capital letter, this makes it available externally
	fmt.Println(vis.MyName)
	// Looking at the scope of the printer file
	vis.PrintVar()

}
