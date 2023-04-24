package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Mario"
	names[1] = "Haryzal"
	names[2] = "Blm Ada"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names)
	fmt.Println(len(names))

	values := [3]uint8{10, 23, 34}
	fmt.Println(values)
	fmt.Println(len(values))
}
