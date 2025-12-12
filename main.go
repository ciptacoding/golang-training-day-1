package main

import (
	"fmt"
)

func main() {
	// Create 3 variables with string, int, and boolean types
	var name string = "Raihana Rahiela"
	var age int = 22
	var myGirlfriend bool = true

	// Print the variables using fmt.Printf with appropriate Go verbs
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Is Active: %t\n", myGirlfriend)

	// Print the data types using %T verb
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", myGirlfriend)
}
