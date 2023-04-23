package main

import (
	"fmt"
)

func main() {
	var number32 int32 = 100000
	var number64 int64 = int64(number32)
	var number8 int8 = int8(number32)

	fmt.Println(number32)
	fmt.Println(number64)
	fmt.Println(number8)

	e := "Mario"[0]
	fmt.Println(e)
	fmt.Println(string(e))
}
