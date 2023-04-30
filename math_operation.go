package main

import "fmt"

func main() {
	var result int8 = 10 + 10
	fmt.Println(result)

	var (
		a uint32
		b uint32
	)
	a = 20
	b = 30
	var resultOperation uint32 = a * b
	fmt.Println(resultOperation)

	// augmented assignments
	var subTotal uint32 = 10000
	subTotal += 2000
	fmt.Println(subTotal)

	// unary operator
	subTotal++
	fmt.Println(subTotal)
	subTotal++
	fmt.Println(subTotal)
}
