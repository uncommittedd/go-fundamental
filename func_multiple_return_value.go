package main

import "fmt"

func getFullName() (string, string) {
	return "Mario", "Haryzal"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
