package main

import "fmt"

func main() {
	var (
		firstName = "Mario"
		lastName  = "Haryzal"
		city      string
		age       int8
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	city = "Bengkulu"
	fmt.Println(city)
	city = "Jakarta"
	fmt.Println("Now =", city)
	age = 25
	fmt.Println(age)
}
