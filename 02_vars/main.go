package main

import "fmt"

func main() {

	// Using the var keyword
	//var name = "Jose"
	var age = 37.5
	const isCool = true

	//Shorthand
	// name := "Jose"
	// size := 1.5

	name, size := "Jose", 1.5

	fmt.Println(age, name, isCool, size)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)

}
