package main

import "fmt"

func factorialRecursive(number int) int {
	if number == 1 {
		return number
	} else {
		return number * factorialRecursive(number-1)
	}
}

func main() {
	resultFactorial := factorialRecursive(5)
	fmt.Println(resultFactorial)
	fmt.Println(5 * 4 * 3 * 2 * 1)
}
