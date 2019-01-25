package main

import "fmt"

func main() {
	// Array
	//var fruitArr [2]string

	//Assign values
	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"

	//Declare and assign
	fruitArr := [2]string{"Apple", "Orange"}

	//Slices
	fruitSlice := []string{"Banana", "Pinneaple", "Mango"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
	fmt.Println(fruitSlice[1], fruitSlice, len(fruitSlice), fruitSlice[1:3])
}
