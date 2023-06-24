package main

import "fmt"

type Person struct {
	name, address string
	age           int
}

func main() {
	var mario Person
	mario.name = "Mario"
	mario.address = "Jakarta"
	mario.age = 24

	fmt.Println(mario)
	fmt.Println(mario.name)
	fmt.Println(mario.address)
	fmt.Println(mario.age)

	john := Person{
		name:    "john",
		address: "Land Of Down",
		age:     24,
	}
	fmt.Println(john)

	glo := Person{"Glo", "Land of Down", 24}
	fmt.Println(glo)
}
