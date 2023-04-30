package main

import "fmt"

func sumAll(number ...uint32) uint32 {
	var total uint32 = 0
	for _, v := range number {
		total += v
	}
	return total
}

func main() {
	fmt.Println(sumAll(2, 3, 4, 56, 2))
	numbers := []uint32{34, 45, 2, 4545}
	fmt.Println(sumAll(numbers...))
}
