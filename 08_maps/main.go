package main

import "fmt"

func main() {
	//Define a ma

	emails := make(map[string]string)

	// Assing key values
	emails["Bob"] = "bobo@email.com"
	emails["Sharon"] = "sharon@email.com"
	emails["Mike"] = "mike@email.com"
	fmt.Println(emails, len(emails))
	fmt.Println(emails["Bob"])

	//Delete
	delete(emails, "Bob")
	fmt.Println(emails, len(emails))

	//Decalre map and add key values
	email := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@email.com"}
	fmt.Println(email)

}
