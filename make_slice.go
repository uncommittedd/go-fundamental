package main

import "fmt"

func main() {
	persons := make([]string, 3)
	persons[0] = "mario"
	persons[1] = "john"
	persons[2] = "doe"

	fmt.Println(persons)
	fmt.Println(len(persons))
	fmt.Println(cap(persons))

	// make new slice for target copy
	copySlice := make([]string, len(persons))
	// execution copy
	copy(copySlice, persons)
	fmt.Println(copySlice)

	thisArray := [...]int{1, 2, 3, 4, 5}
	thisSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("thisArray: %v\n", thisArray)
	fmt.Printf("thisSlice: %v\n", thisSlice)
	fmt.Println(len(thisSlice), len(thisSlice))
}
