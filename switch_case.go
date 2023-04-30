package main

import "fmt"

func main() {
	name := "Mario Haryzal"

	// regular switch case
	switch name {
	case "Mario Haryzal":
		fmt.Println("Good Looking")
		fmt.Println("Handsome")
		fmt.Println("Cool")
	case "Mario":
		fmt.Println("Another level")
	default:
		fmt.Println("idk")
	}

	// short switch case
	hobby := "Gaming"
	switch length := len(hobby); length == 6 {
	case true:
		fmt.Println("true boys")
	case false:
		fmt.Println("sorry")
	}

	// condition on case, without switch condition
	x := 5
	switch {
	case x > 2:
		fmt.Println("tru")
	case x < 2:
		fmt.Println("sad, it's wrong")
	}
}
