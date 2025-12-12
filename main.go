package main

import (
	"fmt"
)

func main() {
	// Create 3 variables with string, int, and boolean types
	var name string = "John Doe"
	var age int = 10
	var isActive bool = true

	// Print the variables using fmt.Printf with appropriate Go verbs
	fmt.Printf("%s\n", name)
	fmt.Printf("%d\n", age)
	fmt.Printf("%t\n", isActive)

	// Print the data types using %T verb
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isActive)
}
