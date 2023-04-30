package main

import "fmt"

func main() {
	firstName := "Mario"
	lastName := "Haryzal"

	var isSameString bool = firstName == lastName
	fmt.Println(isSameString)

	value1 := 100
	value2 := 200
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
